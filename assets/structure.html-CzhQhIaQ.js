import{_ as s,c as n,o as e,b as a}from"./app-CvecxDTg.js";const i={},l=a(`<h1 id="目录结构" tabindex="-1"><a class="header-anchor" href="#目录结构"><span>目录结构</span></a></h1><p>gob 框架不仅仅是一个类库，也是一个定义了开发模式和目录结构的框架。</p><p>gob 希望所有使用这个框架的开发人员遵照统一的项目结构进行开发。</p><h2 id="默认目录结构" tabindex="-1"><a class="header-anchor" href="#默认目录结构"><span>默认目录结构</span></a></h2><p>默认创建的项目结构为：</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line"><span class="token operator">&gt;</span> tree</span>
<span class="line"><span class="token builtin class-name">.</span></span>
<span class="line">├─app               // 服务端应用地址</span>
<span class="line">│  ├─console        // 存放自定义命令</span>
<span class="line">│  │  └─command</span>
<span class="line">│  │      └─demo</span>
<span class="line">│  ├─http           // 存放http服务</span>
<span class="line">│  │  ├─module      // 业务模块</span>
<span class="line">│  │  │  └─demo</span>
<span class="line">│  │  │     ├── api.go        // 业务模块接口</span>
<span class="line">│  │  │     ├── dto.go        // 业务模块输出结构</span>
<span class="line">│  │  │     ├── mapper.go     // 将服务结构转换为业务模块输出结构</span>
<span class="line">│  │  │     ├── model.go      // 数据库结构定义</span>
<span class="line">│  │  │     ├── repository.go // 数据库逻辑封装层</span>
<span class="line">│  │  │     └── service.go    // 服务层</span>
<span class="line">│  │  └─swagger     // swagger文件自动生成 </span>
<span class="line">│  └─provider       // 服务提供方</span>
<span class="line">│      └─demo</span>
<span class="line">│          ├── contract.go  // 服务接口层</span>
<span class="line">│          ├── provider.go  // 服务提供方</span>
<span class="line">│          └── service.go   // 服务实现层</span>
<span class="line">├─config            // 配置文件</span>
<span class="line">│  ├─dev</span>
<span class="line">│      ├── app.yaml         // app主应用的配置</span>
<span class="line">│      ├── database.yaml    // 数据库相关配置</span>
<span class="line">│      ├── deploy.yaml      // 部署相关配置</span>
<span class="line">│      ├── log.yaml         // 日志相关配置</span>
<span class="line">│      └── swagger.yaml     // swagger相关配置</span>
<span class="line">│  ├─prod</span>
<span class="line">│  └─test</span>
<span class="line">├─gob_frontend      // 前端应用地址</span>
<span class="line">│  └─src</span>
<span class="line">│      ├─App.vue    // vue入口文件</span>
<span class="line">│      ├─main.js    // 前端入口文件</span>
<span class="line">│      ├─assets</span>
<span class="line">│      ├─components // vue组件</span>
<span class="line">│      ├─router</span>
<span class="line">│      ├─stores</span>
<span class="line">│      └─views</span>
<span class="line">├─storage</span>
<span class="line">│  ├─log            // 存放业务日志</span>
<span class="line">│  ├─cache          // 存放本地缓存</span>
<span class="line">│  ├─coverage       // 存放覆盖率报告</span>
<span class="line">│  └─runtime        // 存放运行时文件</span>
<span class="line">└─test              </span>
<span class="line">    └── env.go      // 设置测试环境相关参数</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>这里主要介绍下业务模块的分层结构</p><h1 id="业务模块分层" tabindex="-1"><a class="header-anchor" href="#业务模块分层"><span>业务模块分层</span></a></h1><p>业务模块的分层设计两种分层模型：简化模型和标准模型。基本稍微复杂一些的业务，都需要使用标准模型开发。</p><h2 id="简化模型" tabindex="-1"><a class="header-anchor" href="#简化模型"><span>简化模型</span></a></h2><p>对于比较简单的业务，每个模块各自定义自己的 model 和 service，在一个 module 文件的文件夹中进行各自模块的业务开发</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line">├── api.go      // 业务模块接口</span>
<span class="line">├── dto.go      // 业务模块输出结构</span>
<span class="line">├── mapper.go   // 将服务结构转换为业务模块输出结构</span>
<span class="line">├── model.go    // 数据库结构定义</span>
<span class="line">├── repository.go   // 数据库逻辑封装层</span>
<span class="line">└── service.go  // 服务</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>具体实现可以参考初始化代码的 Demo 接口实现</p><h2 id="标准模型" tabindex="-1"><a class="header-anchor" href="#标准模型"><span>标准模型</span></a></h2><p>对于比较复杂的业务，模块与模块间的交互比较复杂，有很多公用性，所以提取 service provider 服务作为服务间的相互调用。</p><p>强烈建议使用这种开发模型</p><p>第一步：创建当前业务的 provider。可以使用命令行 <code>./gob provider new</code> 来创建。</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line"><span class="token operator">&gt;</span> ./gob provider new</span>
<span class="line">create a provider</span>
<span class="line">? please input provider name car</span>
<span class="line">? please input provider folder<span class="token punctuation">(</span>default: provider name<span class="token punctuation">)</span>:</span>
<span class="line">create provider success, folder path: /path/app/provider/car</span>
<span class="line">please remember <span class="token function">add</span> provider to kernel</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>定义好 provider 的协议</p><div class="language-golang line-numbers-mode" data-highlighter="prismjs" data-ext="golang" data-title="golang"><pre class="language-golang"><code><span class="line">package demo</span>
<span class="line"></span>
<span class="line">const DemoKey = &quot;demo&quot;</span>
<span class="line"></span>
<span class="line">type IService interface {</span>
<span class="line">	GetAllStudent() []Student</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line">type Student struct {</span>
<span class="line">	ID   int</span>
<span class="line">	Name string</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>实现对应协议：</p><div class="language-golang line-numbers-mode" data-highlighter="prismjs" data-ext="golang" data-title="golang"><pre class="language-golang"><code><span class="line">package demo</span>
<span class="line"></span>
<span class="line">import &quot;github.com/chenbihao/gob/framework&quot;</span>
<span class="line"></span>
<span class="line">type Service struct {</span>
<span class="line">	container framework.Container</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line">func NewService(params ...interface{}) (interface{}, error) {</span>
<span class="line">	container := params[0].(framework.Container)</span>
<span class="line">	return &amp;Service{container: container}, nil</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line">func (s *Service) GetAllStudent() []Student {</span>
<span class="line">	return []Student{</span>
<span class="line">		{</span>
<span class="line">			ID:   1,</span>
<span class="line">			Name: &quot;foo&quot;,</span>
<span class="line">		},</span>
<span class="line">		{</span>
<span class="line">			ID:   2,</span>
<span class="line">			Name: &quot;bar&quot;,</span>
<span class="line">		},</span>
<span class="line">	}</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>第二步：创建当前业务的模块。</p><p>可以按照demo文件夹中文件编写。</p><p>第三步：在当前业务中挂载业务模块。</p><p>第四步：使用 provider 来开发当前业务。</p><div class="language-golang line-numbers-mode" data-highlighter="prismjs" data-ext="golang" data-title="golang"><pre class="language-golang"><code><span class="line">// Demo godoc</span>
<span class="line">// @Summary 获取所有学生</span>
<span class="line">// @Description 获取所有学生</span>
<span class="line">// @Produce  json</span>
<span class="line">// @Tags demo</span>
<span class="line">// @Success 200 array []UserDTO</span>
<span class="line">// @Router /demo/demo2 [get]</span>
<span class="line">func (api *DemoApi) Demo2(c *gin.Context) {</span>
<span class="line">	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)</span>
<span class="line">	students := demoProvider.GetAllStudent()</span>
<span class="line">	usersDTO := StudentsToUserDTOs(students)</span>
<span class="line">	c.JSON(200, usersDTO)</span>
<span class="line">}</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>具体实现可以参考初始化代码的 Demo2 接口实现</p>`,28),d=[l];function p(c,r){return e(),n("div",null,d)}const t=s(i,[["render",p],["__file","structure.html.vue"]]),u=JSON.parse('{"path":"/guide/structure.html","title":"目录结构","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"目录结构","description":null},"headers":[{"level":2,"title":"默认目录结构","slug":"默认目录结构","link":"#默认目录结构","children":[]},{"level":2,"title":"简化模型","slug":"简化模型","link":"#简化模型","children":[]},{"level":2,"title":"标准模型","slug":"标准模型","link":"#标准模型","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/structure.md"}');export{t as comp,u as data};
