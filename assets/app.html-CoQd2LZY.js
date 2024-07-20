import{_ as n,c as s,o as a,b as e}from"./app-BSoEOH6w.js";const p={},l=e(`<h1 id="gob-app" tabindex="-1"><a class="header-anchor" href="#gob-app"><span>gob:app</span></a></h1><h2 id="服务介绍" tabindex="-1"><a class="header-anchor" href="#服务介绍"><span>服务介绍：</span></a></h2><p>提供基础的 app 框架目录结构获取功能</p><h2 id="支持命令" tabindex="-1"><a class="header-anchor" href="#支持命令"><span>支持命令：</span></a></h2><p><a href="../command/app">app</a></p><h2 id="支持配置-无" tabindex="-1"><a class="header-anchor" href="#支持配置-无"><span>支持配置：无</span></a></h2><h2 id="提供方法" tabindex="-1"><a class="header-anchor" href="#提供方法"><span>提供方法：</span></a></h2><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre><code><span class="line"><span class="token keyword">type</span> App <span class="token keyword">interface</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token comment">// AppID 表示当前这个app的唯一id, 可以用于分布式锁等</span></span>
<span class="line">	<span class="token function">AppID</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// Version 定义当前版本</span></span>
<span class="line">	<span class="token function">Version</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// IsToolMode 是否纯工具运行模式</span></span>
<span class="line">	<span class="token function">IsToolMode</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">bool</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// BaseFolder 定义项目基础地址</span></span>
<span class="line">	<span class="token function">BaseFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ---------------- 根目录下</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// AppFolder 定义业务代码所在的目录，用于监控文件变更使用</span></span>
<span class="line">	<span class="token function">AppFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// ConfigFolder 定义了配置文件的路径</span></span>
<span class="line">	<span class="token function">ConfigFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// TestFolder 存放测试所需要的信息</span></span>
<span class="line">	<span class="token function">TestFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// StorageFolder 存储文件地址</span></span>
<span class="line">	<span class="token function">StorageFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// DeployFolder 存放部署的时候创建的文件夹</span></span>
<span class="line">	<span class="token function">DeployFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ---------------- app 目录下</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ConsoleFolderr 定义业务自己的命令行服务提供者地址</span></span>
<span class="line">	<span class="token function">ConsoleFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// HttpFolderr 定义业务自己的web服务提供者地址</span></span>
<span class="line">	<span class="token function">HttpFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// ProviderFolder 定义业务自己的通用服务提供者地址</span></span>
<span class="line">	<span class="token function">ProviderFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ---------------- config 目录下</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ---------------- storage 目录下</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// LogFolder 定义了日志所在路径</span></span>
<span class="line">	<span class="token function">LogFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// MiddlewareFolder 定义业务自己定义的中间件</span></span>
<span class="line">	<span class="token function">MiddlewareFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// CommandFolder 定义业务定义的命令</span></span>
<span class="line">	<span class="token function">CommandFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line">	<span class="token comment">// RuntimeFolder 定义业务的运行中间态信息</span></span>
<span class="line">	<span class="token function">RuntimeFolder</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">string</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// LoadAppConfig 加载新的AppConfig，key为对应的函数转为小写下划线，比如ConfigFolder =&gt; config_folder</span></span>
<span class="line">	<span class="token function">LoadAppConfig</span><span class="token punctuation">(</span>kv <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,8),i=[l];function t(c,o){return a(),s("div",null,i)}const u=n(p,[["render",t],["__file","app.html.vue"]]),r=JSON.parse('{"path":"/provider/app.html","title":"gob:app","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"gob:app","description":null},"headers":[{"level":2,"title":"服务介绍：","slug":"服务介绍","link":"#服务介绍","children":[]},{"level":2,"title":"支持命令：","slug":"支持命令","link":"#支持命令","children":[]},{"level":2,"title":"支持配置：无","slug":"支持配置-无","link":"#支持配置-无","children":[]},{"level":2,"title":"提供方法：","slug":"提供方法","link":"#提供方法","children":[]}],"git":{"updatedTime":1721480673000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":5}]},"filePathRelative":"provider/app.md"}');export{u as comp,r as data};
