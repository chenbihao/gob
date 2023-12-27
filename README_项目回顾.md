# 自我回顾

梳理复习在使用 Golang 搭建一个 web 框架过程中的迭代步骤。

项目地址：[GitHub - chenbihao/gob: go语言编写的web框架](https://github.com/chenbihao/gob)

## 01、使用 net-http 标准库搭建 Web Server

### net-http 标准库

Web Server 的本质就是通过接收并解析 HTTP 请求传输的文本字符，处理后包装成 HTTP 响应文本返回给客户端。

Go 官方提供了 `net/http` 库，方便我们直接创建 web 服务：

```go
// 创建一个Foo路由和处理函数
http.Handle("/foo", fooHandler)

// 创建一个bar路由和处理函数
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

// 监听8080端口
log.Fatal(http.ListenAndServe(":8080", nil))
```

`net/http` 库主要提供了三类对外库函数（功能、func）：

- 为服务端提供创建 HTTP 服务的函数（名字中一般包含 Serve 字样）
- 为客户端提供调用 HTTP 服务的类库（以 HTTP 的 method 同名）
- 提供中转代理的一些函数（Proxy*）

核心结构（模块、struct）：

- Client 负责构建 HTTP 客户端
- Server 负责构建 HTTP 服务端
- ServerMux 负责 HTTP 服务端路由
- Transport、Request、Response、Cookie 负责客户端和服务端传输对应的不同模块

核心函数（能力、method）：

- 第一层：标准库创建 HTTP 服务是通过创建一个 Server 数据结构完成的
- 第二层：Server 数据结构在 for 循环中不断监听每一个连接
- 第三层：每个连接默认开启一个 Goroutine 为其服务
- 第四五层：serverHandler 结构代表请求对应的处理逻辑，并且通过这个结构进行具体业务逻辑处理
- 第六层：Server 数据结构如果没有设置处理函数 Handler，默认使用 DefaultServerMux 处理请求
- 第七层：DefaultServerMux 是使用 map 结构来存储和查找路由规则

## 02、使用 Context 标准库控制请求上下文

### Context 标准库

> 包上下文定义 Context 类型，该类型跨 API 边界和进程之间传输截止时间、取消信号和其他请求范围的值。
>
> 对服务器的传入请求应创建 Context，对服务器的传出调用应接受 Context。它们之间的函数调用链必须传播 Context，也可以将其替换为使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 创建的派生 Context。当一个 Context 被取消时，从它派生的所有 Context 也会被取消。
>
> 不要将 Contexts 存储在结构类型中; 相反，将 Context 显式传递给每个需要它的函数。Context 应为第一个参数，通常命名为 ctx。

库函数（功能、func）：

- `WithCancel`：直接创建可以操作退出的子节点，
- `WithTimeout`：为子节点设置了超时时间（还有多少时间结束）
- `WithDeadline`：为子节点设置了结束时间线（在什么时间结束）

核心结构（模块、struct）：

```go
type Context interface {
    // 当 Context 被取消或者到了 deadline，返回一个被关闭的 channel
    Done() <-chan struct{}
    ...
}

//函数句柄
type CancelFunc func()
```

在树形逻辑链条上， **一个节点其实有两个角色：一是下游树的管理者；二是上游树的被管理者**，那么就对应需要有两个能力：

- 一个是能让整个下游树结束的能力，也就是函数句柄 `CancelFunc`；
- 另外一个是在上游树结束的时候被通知的能力，也就是 `Done()` 方法。同时因为通知是需要不断监听的，所以 `Done()` 方法需要通过 `channel` 作为返回值让使用方进行监听。

官方示例：

```go
package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
    // 创建截止时间
	d := time.Now().Add(shortDuration)
    // 创建有截止时间的 Context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

    // 使用 select 监听 1s 和有截止时间的 Context 哪个先结束
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

### net-http 标准库中应用的 Context 逻辑

![Context生成层次.jpg](https://cos.noobbb.cn/pictures/202312/19_Context%E7%94%9F%E6%88%90%E5%B1%82%E6%AC%A1.jpg)

每个连接的 Context 都是基于 baseContext 复制来的。

对应到代码中就是，在为某个连接开启 Goroutine 的时候，为当前连接创建了一个 `connContext`，这个 `connContext` 是基于 server 中的 Context （ `baseContext`）而来。

这两处都可以注入修改：

- `BaseContext` 是整个 Context 生成的源头，如果我们不希望使用默认的 `context.Backgroud()`，可以替换这个源头。
- 而在每个连接生成自己要使用的 Context 时，会调用 `ConnContext` ，它的第二个参数是 `net.Conn`，能让我们对某些特定连接进行设置，比如要针对性设置某个调用 IP。

源代码：

```go
type Server struct {
	...
    // BaseContext 用来为整个链条创建初始化 Context
    // 如果没有设置的话，默认使用 context.Background()
	BaseContext func(net.Listener) context.Context{}
	
    // ConnContext 用来为每个连接封装 Context
    // 参数中的 context.Context 是从 BaseContext 继承来的
	ConnContext func(ctx context.Context, c net.Conn) context.Context{}
    ...
}
```

### 目标：封装一个自己的 Context

前置：

```go
// 创建 framework 包存放框架文件
// framework 包外为业务文件
```

目标：封装一个自己的 Context：

```go
// 未封装的原生控制器的使用
func Foo1(request *http.Request, response http.ResponseWriter) {}

// 期待封装 Context 后的控制器使用
func Foo2(ctx *framework.Context) error {
	obj := map[string]interface{}{
		"data":   nil,
	}
    // 从请求体中获取参数
 	fooInt := ctx.FormInt("foo", 10)
    // 构建返回结构
	obj["data"] = fooInt
    // 输出返回结构
	return ctx.Json(http.StatusOK, obj)
}
```

将 `request` 和 `response` 封装到自定义的 Context 中，

并且兼容标准库的 Context 接口，`context.go` ：

```go
// 自定义 Context
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	...
}

// 直接返回原生 Context
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// implement context.Context （实现标准 Context 接口）

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}
```

并且自己封装 Context 最终需要提供四类功能函数：

- `base` 封装基本的函数功能（比如获取 http.Request 结构）
- `context` 实现标准 Context 接口
- `request` 封装了 http.Request 的对外接口（query url、form post、json post 等）
- `response` 封装了 http.ResponseWriter 对外接口（Json、HTML、Text 等）

ControllerHandler 定义，框架目录 `controller.go` ：

```go
type ControllerHandler func(c *Context) error  
```

控制器使用，业务目录 `controller.go` ：

```go
func FooControllerHandler(ctx *framework.Context) error {  
    return ctx.Json(200, map[string]interface{}{  
        "code": 0,  
    })  
}  
```

### 目标：为单个请求设置超时

自定义 Context 设置超时：

1. 继承 request 的 Context，创建出一个设置超时时间的 Context；
2. 创建一个新的 Goroutine 来处理具体的业务逻辑；
3. 设计事件处理顺序，当前 Goroutine 监听超时时间 Contex 的 Done() 事件，和具体的业务处理结束事件，哪个先到就先处理哪个。

业务 `controller.go`：

```go
func FooControllerHandler(c *framework.Context) error {
	// 生成一个超时的 Context
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), 1*time.Second)
	// 当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()

	finish := make(chan struct{}, 1)       // 这个 channel 负责通知结束
	panicChan := make(chan interface{}, 1) // 这个 channel 负责通知 panic 异常

	// 创建一个新的 Goroutine 来处理业务逻辑
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		
		// 这里做具体的业务
		time.Sleep(1 * time.Second)
		
		c.Json(200, "ok")
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:      
		// 监听 panic
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()   // 考虑边界情况，加锁
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:                
		// 监听结束事件
		fmt.Println("finish")
	case <-durationCtx.Done():    
		// 监听超时事件
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()   // 考虑边界情况，加锁
		c.Json(500, "time out")
		c.SetHasTimeout()        // 考虑边界情况，当触发超时后避免其他协程重复写入 （todo:没作写保护）
	}
	return nil
}
```

框架上下文添加超时与写保护 `context.go` ：

```go
// 自定义 Context
type Context struct {
	...
	hasTimeout bool        // 是否超时标记位
	writerMux  *sync.Mutex // 写保护机制
}
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

...

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

```

## 03、自定义路由功能

### 目标

路由一般使用的是请求头里的 `Method` 和 `Request-URI` 这两个部分。

希望使用者高效、易用地使用路由模块，基本需求可以有哪些呢？

- 需求 1：HTTP 方法匹配
- 需求 2：静态路由匹配
- 需求 3：批量通用前缀
- 需求 4：动态路由匹配
- 扩展需求：分组嵌套

### 如何实现

简单来讲，核心结构 `Core` 去实现 Handler 接口（`ServeHTTP`），来接管请求处理。

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

并且在 `ServeHTTP` 函数里面，实现框架上下文 Context 的封装以及路由功能 router。

### 代码实现

最终代码 `main.go`：

```go
func main() {

	// 核心框架初始化
	core := framework.NewCore()

	// 设置路由
	registerRouter(core)

	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
```

定义 ControllerHandler ，框架文件夹 `controller.go`：

```go
type ControllerHandler func(c *Context) error
```

并且简单实现几个业务 Controller ，例如业务文件夹的 `user_controller.go`：

```go
func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}
```

业务文件夹的注册路由 ，`router.go`：

```go
// 注册路由规则
func registerRouter(core *framework.Core) {
	// 需求1+2:HTTP方法+静态路由匹配
	core.Get("/user/login", UserLoginController)

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
		
		// 扩展需求：分组嵌套
		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
```

利用接口设计通用的分组定义，并且实现 `group.go`：

```go
// IGroup 代表前缀分组
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// 实现嵌套group
	Group(string) IGroup
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core  // 指向core结构
	parent *Group // 指向上一个Group，如果有的话
	prefix string // 这个group的通用前缀
}

// 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// 实现Get方法
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

...  //  POST、PUT、DELETE

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// 实现 Group 方法
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
```

核心框架代码，包含 Handler 接口实现、初始化路由、分组接口实现 。`core.go`：

```go
// 框架核心结构
type Core struct {
	router      map[string]*Tree    // all routers              // 一级匹配HTTP方法，二级字典树匹配
}

// 初始化Core结构
func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

// 匹配GET 方法, 增加路由规则
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

...  //  POST、PUT、DELETE

// 前缀分组
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// 框架核心结构实现 Handler 接口
// 所有请求都进入这个函数, 这个函数负责路由分发
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
```

字典树的实现 ，`trie.go`：

```go
// 代表树结构
type Tree struct {
	root *node // 根节点
}

// 代表节点
type node struct {
	isLast   bool                // 是否可以成为最终的路由规则。该节点是否能成为一个独立的uri, 是否终极节点
	segment  string              // uri中的字符串，代表这个节点表示的路由中某个段的字符串
	handler ControllerHandler    // 代表这个节点中包含的控制器，用于最终加载调用
	childs   []*node             // 代表这个节点下的子节点
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
}

// 判断一个segment是否是通用segment，即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}
	// 如果segment是通配符，则所有下一层子节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}
	nodes := make([]*node, 0, len(n.childs))
	// 过滤所有的下一层子节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			// 如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// 如果下一层子节点没有通配符，但是文本完全匹配，则满足需求
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	// 使用分隔符将uri切割为两个部分
	uri = strings.TrimPrefix(uri, "/") //  【/】开头的 url 去 Split 后第一层是空的，节省一层
	segments := strings.SplitN(uri, "/", 2)
	// 第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)

	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在, 直接返回nil
	if len(cnodes) == 0 {
		return nil
	}

	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		// 如果segment已经是最后一个节点，判断这些cnode是否有isLast标志
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		// 都不是最后一个节点
		return nil
	}

	// 如果有2个segment, 递归每个子节点继续进行查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// 增加路由节点
func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root  
	
	// 确认路由是否冲突
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}
	uri = strings.TrimPrefix(uri, "/") //  【/】开头的 url 去 Split 后第一层是空的，节省一层
	segments := strings.Split(uri, "/")
	// 对每个segment
	for index, segment := range segments {

		// 最终进入Node segment的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前node的节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true  
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}

// 匹配uri
func (tree *Tree) FindHandler(uri string) ControllerHandler {
	// 直接复用matchNode函数，uri是不带通配符的地址
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
```

## 04、利用中间件提升扩展性

### 目标

03 说过，核心结构 `Core` 去实现 Handler 接口（`ServeHTTP`），来接管请求处理。

并且在 `ServeHTTP` 函数里面，实现框架上下文 Context 的封装以及路由功能 router。

在此基础上，把非业务逻辑的通用型需求，抽取成中间件来使用。

- 扩展需求 1：全局注册中间件
- 扩展需求 2：为单个路由注册中间件
- 扩展需求 3：为组嵌套中的单个路由注册中间件

以上都可以注册单个或者多个中间件

### 如何实现

改造成链路调用。

引入 pipeline 思想，将所有中间件做成一个链条，通过这个链条的调用，来实现中间件机制。

在架构层面，中间件机制就相当于，在每个请求的横切面统一注入了一个逻辑。

### 代码实现

最终使用效果，业务文件夹 `router.go`：

```go
// 注册路由规则
func registerRouter(core *framework.Core) {

	// 扩展需求1：core中使用use注册全局中间件 （需放在前面）
	core.Use(middleware.Recovery(), middleware.Cost())
	
	// 扩展需求2：在core中使用middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	subjectApi := core.Group("/subject")
	{
		...
		// 扩展需求3：在 group 中使用 middleware.Test3() 为单个路由增加中间件
		subjectApi.Get("/middleware/test3", middleware.Test3(), SubjectAddController)
	}
	core.Get("/timeout", middleware.Timeout(time.Second), TimeoutController)
}
```

字典树中的 Handler 改造成控制器链路 Handlers，找到路由 node 时也就能找到对应的控制器链路，`trie.go`：

```go
// 代表节点
type node struct {
	... 
	handlers []ControllerHandler // 中间件+控制器
}
...
// 增加路由节点
func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	...
				cnode.handlers = handlers
	...
}

// 匹配uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}

```

改造框架上下文，由上下文存储 handler 链条，并且维护一个链路下标，`context.go` ：

```go
// 自定义 Context
type Context struct {
	...
	handlers []ControllerHandler // 当前请求的handler链条
	index    int                 // 当前请求调用到调用链的哪个节点
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		...
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 核心函数，调用context的下一个函数 
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}
```

`Next()` 函数会在框架的两个地方被调用：

- 第一个是在此次请求处理的入口处，即 Core 的 ServeHttp；
- 第二个是在**每个中间件**的逻辑代码中，用于调用下个中间件。

上面是链路的改造，使框架中间件链路能顺利连起来，下面开始中间件注册。

- 首先为 Group 和 Core 两个结构增加注册中间件入口 `Use()`
- 并且在路由注册时，需要支持可变参数（`handlers ...ControllerHandler`）、聚合控制器（`allHandlers`）。

改造 Group，使中间件思想融入嵌套分组中 `group.go`：

```go
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	// 实现嵌套group
	Group(string) IGroup
	// 嵌套中间件
	Use(middlewares ...ControllerHandler)
}
type Group struct {
	...
	middlewares []ControllerHandler // 存放中间件
}

// 实现Get方法
func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)  // 聚合
	g.core.Get(uri, allHandlers...)
}

...  // NewGroup  //  POST、PUT、DELETE

// 获取某个group的middleware
// 这里就是获取除了Get/Post/Put/Delete之外设置的middleware
func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}
```

核心同步修改实现，增加全局中间件应用 ，`core.go`：

```go
// 框架核心结构
type Core struct {
	...
	middlewares []ControllerHandler // 从 core 这边设置的中间件   
}

... // NewCore

// 匹配 GET 方法, 增加路由规则
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	// 将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

...  //  POST、PUT、DELETE   、Group

// 注册全局中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	...
}

// 框架核心结构实现 Handler 接口
// 所有请求都进入这个函数, 这个函数负责路由分发
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	handlers := c.FindRouteByRequest(request)
	if handlers == nil {
		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	// 设置context中的handlers字段
	ctx.SetHandlers(handlers)

	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
```

### 中间件的编写

在中间件中调用下一个链路（`ctx.Next()`），形成闭环。例如超时 `timeout.go`：

```go
func Timeout(d time.Duration) framework.ControllerHandler {
	// 使用函数回调
	return func(ctx *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 使用next执行具体的业务逻辑
			ctx.Next()

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			ctx.Json(500, "time out")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			ctx.Json(500, "time out")
			ctx.SetHasTimeout()
		}
		return nil
	}
}
```

## 05、封装让框架更好用

### 目标

尽量在 context 这个数据结构中，封装“读取请求数据”和“封装返回数据”中的方法。

- 读取请求数据
    - Header 信息
        - 基础信息，比如请求地址、请求方法、请求 IP、请求域名、Cookie 信息等。
        - 更细节的内容编码格式、缓存时长等，由于涉及的 HTTP 协议细节内容比较多，我们很难将每个细节都封装出来，但是它们都是以 key=value 的形式传递到服务端的，所以这里也考虑封装一个通用的方法。
    - Body 信息（HTTP 是已经以某种形式封装好的）
        - 可能是 JSON 格式、XML 格式、其他格式
        - 也可能是 Form 表单格式
            - 它可能包含 File 文件，请求参数和返回值肯定和其他的 Form 表单字段是不一样的，需要我们对其单独封装一个函数

- 封装返回数据
    - Header 头部
        - 我们经常要设置的是返回状态码和 Cookie，所以单独为其封装。
        - 其他的 Header 同样是 key=value 形式设置的，设置一个通用的方法即可。
    - 返回 Body 体
        - 比如 JSON、JSONP、XML、HTML 或者其他文本格式，要针对不同的 Body 体形式，进行不同的封装

### 如何实现

首先，定义一个清晰的、包含若干个方法的接口，可以让使用者更加清晰明了地使用框架，同时做到“实现解耦”。

实现接口，可以利用 [第三方库cast](https://github.com/spf13/cast) 方便编码，利用官方库 [html/template](https://golang.org/pkg/html/template/) 方便模板数据替换。

### 代码实现

#### IRequest 接口定义与实现

读取请求数据 IRequest，`request.go`：

```go
package framework
import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
)

// 代表请求包含的方法
type IRequest interface {
	// 请求地址 url 中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// 路由匹配中带的参数
	// 形如 /book/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	// form 表单中带的参数
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

	// json body
	BindJson(obj interface{}) error
	// xml body
	BindXml(obj interface{}) error
	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string][]string
	Header(key string) (string, bool)
	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

const defaultMultipartMemory = 32 << 20 // 32 MB

var _ IRequest = new(Context) // 确保类型实现接口

// 获取请求地址中所有参数
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// 请求地址url中带的参数
// 形如: foo.com?a=1&b=bar&c[]=bar

// 获取Int类型的请求参数
func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			// 使用cast库将string转换为Int
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) Query(key string) interface{} {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}

// 路由匹配中带的参数
// 形如 /book/:id
func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if val := ctx.Param(key); val != nil {
		// 通过cast进行类型转换
		return cast.ToInt(val), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if val := ctx.Param(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// 获取路由参数
func (ctx *Context) Param(key string) interface{} {
	if ctx.params != nil {
		if val, ok := ctx.params[key]; ok {
			return val
		}
	}
	return nil
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		ctx.request.ParseForm()
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0], true
	}
	return def, false
}

func (ctx *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {
	if ctx.request.MultipartForm == nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

func (ctx *Context) Form(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

// 将body文本解析到obj结构体中
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		// 读取文本
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		// 重新填充request.Body，为后续的逻辑二次读取做准备
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 解析到obj结构体中
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// xml body
func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// 其他格式
func (ctx *Context) GetRawData() ([]byte, error) {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return nil, err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil, errors.New("ctx.request empty")
}

// 基础信息
func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) Host() string {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIp() string {
	r := ctx.request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// header
func (ctx *Context) Headers() map[string][]string {
	return ctx.request.Header
}

func (ctx *Context) Header(key string) (string, bool) {
	vals := ctx.request.Header.Values(key)
	if vals == nil || len(vals) <= 0 {
		return "", false
	}
	return vals[0], true
}

// cookie
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

func (ctx *Context) Cookie(key string) (string, bool) {
	cookies := ctx.Cookies()
	if val, ok := cookies[key]; ok {
		return val, true
	}
	return "", false
}

```

#### IResponse 接口定义与实现

封装返回数据 IResponse，`response.go`：

```go
package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// IResponse 代表返回方法
type IResponse interface {
	// Json 输出
	Json(obj interface{}) IResponse
	// Jsonp 输出
	Jsonp(obj interface{}) IResponse
	//xml 输出
	Xml(obj interface{}) IResponse
	// html 输出
	Html(template string, obj interface{}) IResponse
	// string
	Text(format string, values ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse
	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse
	// 设置状态码
	SetStatus(code int) IResponse
	// 设置 200 状态
	SetOkStatus() IResponse
}

var _ IResponse = new(Context) // 确保类型实现接口

// Jsonp输出
func (ctx *Context) Jsonp(obj interface{}) IResponse {
	// 获取请求参数callback
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}
	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}
	// 输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

// xml输出
func (ctx *Context) Xml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/html")
	ctx.responseWriter.Write(byt)
	return ctx
}

// html输出
func (ctx *Context) Html(file string, obj interface{}) IResponse {
	// 读取模版文件，创建template实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	// 执行Execute方法将obj和模版进行结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

// string
func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

// 重定向
func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

// header
func (ctx *Context) SetHeader(key string, val string) IResponse {
	ctx.responseWriter.Header().Add(key, val)
	return ctx
}

// Cookie
func (ctx *Context) SetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

// 设置状态码
func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

// 设置200状态
func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return ctx
}

```

## 06、

## 07、

## 08、

## 09、

## 10、
