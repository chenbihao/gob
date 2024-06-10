import{_ as a,r as e,c as p,a as t,d as n,e as i,w as c,b as o,o as l}from"./app-BBGivji7.js";const u={},r=o(`<h1 id="服务提供者" tabindex="-1"><a class="header-anchor" href="#服务提供者"><span>服务提供者</span></a></h1><h2 id="指南" tabindex="-1"><a class="header-anchor" href="#指南"><span>指南</span></a></h2><p>gob 框架使用 ServiceProvider 机制来满足协议，通过 ServiceProvider 提供某个协议服务的具体实现。</p><p>这样如果开发者对具体的实现协议的服务类的具体实现不满意，则可以很方便的通过切换具体协议的 ServiceProvider 来进行具体服务的切换。</p><p>一个 ServiceProvider 是一个单独的文件夹，它包含服务提供和服务实现。具体可以参考 <code>framework/provider/demo</code></p><p>一个 SerivceProvider 就是一个独立的包，这个包可以作为插件独立地发布和分享。</p><p>你也可以定义一个无 contract 的 ServiceProvider ，其中的 <code>Name()</code> 需要保证唯一。</p><h2 id="创建" tabindex="-1"><a class="header-anchor" href="#创建"><span>创建</span></a></h2><p>我们可以使用命令 <code>./gob provider new</code> 来创建一个新的service provider</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line"><span class="token operator">&gt;</span> ./gob provider new</span>
<span class="line">创建一个服务</span>
<span class="line">? 请输入服务名称<span class="token punctuation">(</span>服务凭证<span class="token punctuation">)</span>： demop</span>
<span class="line">? 请输入服务所在目录名称<span class="token punctuation">(</span>默认: 同服务名称<span class="token punctuation">)</span>:</span>
<span class="line">创建服务成功, 文件夹地址: D:<span class="token punctuation">\\</span>DevProjects<span class="token punctuation">\\</span>自己库<span class="token punctuation">\\</span>gob<span class="token punctuation">\\</span>app<span class="token punctuation">\\</span>provider<span class="token punctuation">\\</span>demop</span>
<span class="line">请不要忘记挂载新创建的服务</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>该命令会在<code>app/provider/</code> 目录下创建一个对应的服务提供者文件夹。</p><p>并且初始化好三个文件： <code>contract.go</code>, <code>provider.go</code>, <code>service.go</code></p><h2 id="自定义" tabindex="-1"><a class="header-anchor" href="#自定义"><span>自定义</span></a></h2><p>我们需要编写这三个文件：</p><h3 id="contract-go" tabindex="-1"><a class="header-anchor" href="#contract-go"><span>contract.go</span></a></h3><p><code>contract.go</code> 定义了这个服务提供方提供的协议接口。</p><p>gob 框架任务，作为一个业务的服务提供者，定义一个好的协议是最重要的事情。</p><p>所以 <code>contract.go</code> 中定义了一个 Service 接口，在其中定义各种方法，包含输入参数和返回参数。</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">package demo</span>
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
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>其中还定义了一个Key， 这个 Key 是全应用唯一的，服务提供者将服务以 Key 关键字注入到容器中。服务使用者使用 Key 关键字获取服务。</p><h3 id="provider" tabindex="-1"><a class="header-anchor" href="#provider"><span>provider</span></a></h3><p>provider.go 提供服务适配的实现，实现一个 Provider 必须实现对应的五个方法</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">package</span> demo</span>
<span class="line"></span>
<span class="line"><span class="token keyword">import</span> <span class="token string">&quot;github.com/chenbihao/gob/framework&quot;</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">type</span> DemoProvider <span class="token keyword">struct</span> <span class="token punctuation">{</span></span>
<span class="line">	framework<span class="token punctuation">.</span>ServiceProvider</span>
<span class="line"></span>
<span class="line">	c framework<span class="token punctuation">.</span>Container</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>sp <span class="token operator">*</span>DemoProvider<span class="token punctuation">)</span> <span class="token function">Name</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token keyword">return</span> DemoKey</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>sp <span class="token operator">*</span>DemoProvider<span class="token punctuation">)</span> <span class="token function">Register</span><span class="token punctuation">(</span>c framework<span class="token punctuation">.</span>Container<span class="token punctuation">)</span> framework<span class="token punctuation">.</span>NewInstance <span class="token punctuation">{</span></span>
<span class="line">	<span class="token keyword">return</span> NewService</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>sp <span class="token operator">*</span>DemoProvider<span class="token punctuation">)</span> <span class="token function">IsDefer</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">bool</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token keyword">return</span> <span class="token boolean">false</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>sp <span class="token operator">*</span>DemoProvider<span class="token punctuation">)</span> <span class="token function">Params</span><span class="token punctuation">(</span>c framework<span class="token punctuation">.</span>Container<span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token keyword">return</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">{</span>sp<span class="token punctuation">.</span>c<span class="token punctuation">}</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>sp <span class="token operator">*</span>DemoProvider<span class="token punctuation">)</span> <span class="token function">Boot</span><span class="token punctuation">(</span>c framework<span class="token punctuation">.</span>Container<span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span></span>
<span class="line">	sp<span class="token punctuation">.</span>c <span class="token operator">=</span> c</span>
<span class="line">	<span class="token keyword">return</span> <span class="token boolean">nil</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li><code>Name()</code> // 指定这个服务提供者提供的服务对应的接口的关键字</li><li><code>Register()</code> // 这个服务提供者注册的时候调用的方法，一般是指定初始化服务的函数名</li><li><code>IsDefer()</code> // 这个服务是否是使用时候再进行初始化，false为注册的时候直接进行初始化服务</li><li><code>Params()</code> // 初始化服务的时候对服务注入什么参数，一般把 container 注入到服务中</li><li><code>Boot()</code> // 初始化之前调用的函数，一般设置一些全局的Provider</li></ul><h3 id="service-go" tabindex="-1"><a class="header-anchor" href="#service-go"><span>service.go</span></a></h3><p>service.go提供具体的实现，它至少需要提供一个实例化的方法 <code>NewService(params ...interface{}) (interface{}, error)</code>。</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">package</span> demo</span>
<span class="line"></span>
<span class="line"><span class="token keyword">import</span> <span class="token string">&quot;github.com/chenbihao/gob/framework&quot;</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">type</span> Service <span class="token keyword">struct</span> <span class="token punctuation">{</span></span>
<span class="line">	container framework<span class="token punctuation">.</span>Container</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token function">NewService</span><span class="token punctuation">(</span>params <span class="token operator">...</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span></span>
<span class="line">	container <span class="token operator">:=</span> params<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span>framework<span class="token punctuation">.</span>Container<span class="token punctuation">)</span></span>
<span class="line">	<span class="token keyword">return</span> <span class="token operator">&amp;</span>Service<span class="token punctuation">{</span>container<span class="token punctuation">:</span> container<span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token boolean">nil</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>s <span class="token operator">*</span>Service<span class="token punctuation">)</span> <span class="token function">GetAllStudent</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span>Student <span class="token punctuation">{</span></span>
<span class="line">	<span class="token keyword">return</span> <span class="token punctuation">[</span><span class="token punctuation">]</span>Student<span class="token punctuation">{</span></span>
<span class="line">		<span class="token punctuation">{</span></span>
<span class="line">			ID<span class="token punctuation">:</span>   <span class="token number">1</span><span class="token punctuation">,</span></span>
<span class="line">			Name<span class="token punctuation">:</span> <span class="token string">&quot;foo&quot;</span><span class="token punctuation">,</span></span>
<span class="line">		<span class="token punctuation">}</span><span class="token punctuation">,</span></span>
<span class="line">		<span class="token punctuation">{</span></span>
<span class="line">			ID<span class="token punctuation">:</span>   <span class="token number">2</span><span class="token punctuation">,</span></span>
<span class="line">			Name<span class="token punctuation">:</span> <span class="token string">&quot;bar&quot;</span><span class="token punctuation">,</span></span>
<span class="line">		<span class="token punctuation">}</span><span class="token punctuation">,</span></span>
<span class="line">	<span class="token punctuation">}</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="注入" tabindex="-1"><a class="header-anchor" href="#注入"><span>注入</span></a></h2><p>gob 的路由，controller 的定义是选择基于gin框架进行扩展的。</p><p>所有的 gin 框架的路由、参数获取、验证、context都和gin框架是相同的。</p><p>唯一不同的是 gin 的全局路由<code>gin.Engine</code>实现了gob的容器结构，可以对<code>gin.Engine</code>进行服务提供的注入，且可以从context中获取具体的服务。</p><p>gob 提供两种服务注入的方法：</p><ul><li>Bind: 将一个 ServiceProvider 绑定到容器中，可以控制其是否是单例</li><li>Singleton: 将一个单例 ServiceProvider 绑定到容器中</li></ul><p>建议在文件夹 <code>app/provider/kernel.go</code> 中进行服务注入</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">func</span> <span class="token function">RegisterCustomProvider</span><span class="token punctuation">(</span>c framework<span class="token punctuation">.</span>Container<span class="token punctuation">)</span> <span class="token punctuation">{</span></span>
<span class="line">	c<span class="token punctuation">.</span><span class="token function">Bind</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>demo<span class="token punctuation">.</span>DemoProvider<span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token boolean">true</span><span class="token punctuation">)</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>当然你也可以在某个业务模块路由注册的时候进行服务注入</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">func</span> <span class="token function">Register</span><span class="token punctuation">(</span>r <span class="token operator">*</span>gin<span class="token punctuation">.</span>Engine<span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span></span>
<span class="line">	api <span class="token operator">:=</span> <span class="token function">NewDemoApi</span><span class="token punctuation">(</span><span class="token punctuation">)</span></span>
<span class="line">	r<span class="token punctuation">.</span><span class="token function">Container</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Singleton</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>demoService<span class="token punctuation">.</span>DemoProvider<span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	r<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/demo/demo&quot;</span><span class="token punctuation">,</span> api<span class="token punctuation">.</span>Demo<span class="token punctuation">)</span></span>
<span class="line">	r<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/demo/demo2&quot;</span><span class="token punctuation">,</span> api<span class="token punctuation">.</span>Demo2<span class="token punctuation">)</span></span>
<span class="line">	<span class="token keyword">return</span> <span class="token boolean">nil</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="获取" tabindex="-1"><a class="header-anchor" href="#获取"><span>获取</span></a></h2><p>gob提供了三种服务获取的方法：</p><ul><li>Make: 根据一个Key获取服务，获取不到获取报错</li><li>MustMake: 根据一个Key获取服务，获取不到返回空</li><li>MakeNew: 根据一个Key获取服务，每次获取都实例化，对应的ServiceProvider必须是以非单例形式注入</li></ul><p>你可以在任意一个可以获取到 container 的地方进行服务的获取。</p><p>业务模块中:</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">func</span> <span class="token punctuation">(</span>api <span class="token operator">*</span>DemoApi<span class="token punctuation">)</span> <span class="token function">Demo2</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span></span>
<span class="line">	demoProvider <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">MustMake</span><span class="token punctuation">(</span>demoService<span class="token punctuation">.</span>DemoKey<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token punctuation">(</span>demoService<span class="token punctuation">.</span>IService<span class="token punctuation">)</span></span>
<span class="line">	students <span class="token operator">:=</span> demoProvider<span class="token punctuation">.</span><span class="token function">GetAllStudent</span><span class="token punctuation">(</span><span class="token punctuation">)</span></span>
<span class="line">	usersDTO <span class="token operator">:=</span> <span class="token function">StudentsToUserDTOs</span><span class="token punctuation">(</span>students<span class="token punctuation">)</span></span>
<span class="line">	c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span><span class="token number">200</span><span class="token punctuation">,</span> usersDTO<span class="token punctuation">)</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>命令行中：</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">var</span> CenterCommand <span class="token operator">=</span> <span class="token operator">&amp;</span>cobra<span class="token punctuation">.</span>Command<span class="token punctuation">{</span></span>
<span class="line">	Use<span class="token punctuation">:</span>   <span class="token string">&quot;direct_center&quot;</span><span class="token punctuation">,</span></span>
<span class="line">	Short<span class="token punctuation">:</span> <span class="token string">&quot;计算区域中心点&quot;</span><span class="token punctuation">,</span></span>
<span class="line">	RunE<span class="token punctuation">:</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>cobra<span class="token punctuation">.</span>Command<span class="token punctuation">,</span> args <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span></span>
<span class="line">		container <span class="token operator">:=</span> util<span class="token punctuation">.</span><span class="token function">GetContainer</span><span class="token punctuation">(</span>c<span class="token punctuation">.</span><span class="token function">Root</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span></span>
<span class="line">		app <span class="token operator">:=</span> container<span class="token punctuation">.</span><span class="token function">MustMake</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>AppKey<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>App<span class="token punctuation">)</span></span>
<span class="line">        <span class="token keyword">return</span> <span class="token boolean">nil</span></span>
<span class="line">    <span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>甚至于另外一个服务提供者中：</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">type</span> Service <span class="token keyword">struct</span> <span class="token punctuation">{</span></span>
<span class="line">	c framework<span class="token punctuation">.</span>Container</span>
<span class="line"></span>
<span class="line">	baseURL <span class="token builtin">string</span></span>
<span class="line">	userID  <span class="token builtin">string</span></span>
<span class="line">	token   <span class="token builtin">string</span></span>
<span class="line">	logger  contract<span class="token punctuation">.</span>Log</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">func</span> <span class="token function">NewService</span><span class="token punctuation">(</span>params <span class="token operator">...</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span></span>
<span class="line">	c <span class="token operator">:=</span> params<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span>framework<span class="token punctuation">.</span>Container<span class="token punctuation">)</span></span>
<span class="line">	config <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">MustMake</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>ConfigKey<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>Config<span class="token punctuation">)</span></span>
<span class="line">	baseURL <span class="token operator">:=</span> config<span class="token punctuation">.</span><span class="token function">GetString</span><span class="token punctuation">(</span><span class="token string">&quot;app.stsmap.url&quot;</span><span class="token punctuation">)</span></span>
<span class="line">	userID <span class="token operator">:=</span> config<span class="token punctuation">.</span><span class="token function">GetString</span><span class="token punctuation">(</span><span class="token string">&quot;app.stsmap.user_id&quot;</span><span class="token punctuation">)</span></span>
<span class="line">	token <span class="token operator">:=</span> config<span class="token punctuation">.</span><span class="token function">GetString</span><span class="token punctuation">(</span><span class="token string">&quot;app.stsmap.token&quot;</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	logger <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">MustMake</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>LogKey<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>Log<span class="token punctuation">)</span></span>
<span class="line">	<span class="token keyword">return</span> <span class="token operator">&amp;</span>Service<span class="token punctuation">{</span>baseURL<span class="token punctuation">:</span> baseURL<span class="token punctuation">,</span> logger<span class="token punctuation">:</span> logger<span class="token punctuation">,</span> userID<span class="token punctuation">:</span> userID<span class="token punctuation">,</span> token<span class="token punctuation">:</span> token<span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token boolean">nil</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="gob-provider" tabindex="-1"><a class="header-anchor" href="#gob-provider"><span>gob provider</span></a></h2><p>gob 框架默认自带了一些服务提供者，提供基础的服务接口协议，可以通过 <code>./gob provider list</code> 来获取已经安装的服务提供者。</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">&gt; ./gob provider list</span>
<span class="line">gob:cache</span>
<span class="line">gob:env</span>
<span class="line">gob:distributed</span>
<span class="line">gob:config</span>
<span class="line">gob:log</span>
<span class="line">gob:trace</span>
<span class="line">gob:orm</span>
<span class="line">gob:redis</span>
<span class="line">gob:kernel</span>
<span class="line">gob:app</span>
<span class="line">gob:id</span>
<span class="line">gob:ssh</span>
<span class="line">demo</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>gob 框架自带的服务提供者的 key 是以 <code>gob:</code> 开头。目的为的是与自定义服务提供者的 key 区别开。</p>`,51);function d(k,v){const s=e("RouteLink");return l(),p("div",null,[r,t("p",null,[n("gob 框架自带的服务提供者具体定义的协议可以参考："),i(s,{to:"/provider/"},{default:c(()=>[n("provider")]),_:1})])])}const b=a(u,[["render",d],["__file","provider.html.vue"]]),g=JSON.parse('{"path":"/guide/provider.html","title":"服务提供者","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"服务提供者","description":null},"headers":[{"level":2,"title":"指南","slug":"指南","link":"#指南","children":[]},{"level":2,"title":"创建","slug":"创建","link":"#创建","children":[]},{"level":2,"title":"自定义","slug":"自定义","link":"#自定义","children":[{"level":3,"title":"contract.go","slug":"contract-go","link":"#contract-go","children":[]},{"level":3,"title":"provider","slug":"provider","link":"#provider","children":[]},{"level":3,"title":"service.go","slug":"service-go","link":"#service-go","children":[]}]},{"level":2,"title":"注入","slug":"注入","link":"#注入","children":[]},{"level":2,"title":"获取","slug":"获取","link":"#获取","children":[]},{"level":2,"title":"gob provider","slug":"gob-provider","link":"#gob-provider","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/provider.md"}');export{b as comp,g as data};
