import{_ as n,c as e,o as s,b as a}from"./app-BBGivji7.js";const i={},l=a(`<h1 id="环境变量" tabindex="-1"><a class="header-anchor" href="#环境变量"><span>环境变量</span></a></h1><h2 id="设置" tabindex="-1"><a class="header-anchor" href="#设置"><span>设置</span></a></h2><p>gob 支持使用应用默认下的隐藏文件 <code>.env</code> 来配置各个机器不同的环境变量。</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">APP_ENV=dev</span>
<span class="line"></span>
<span class="line">DB_PASSWORD=mypassword</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>环境变量的设置可以在配置文件中通过 <code>env([环境变量])</code> 来获取到。</p><p>比如：</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">mysql:</span>
<span class="line">    hostname: 127.0.0.1</span>
<span class="line">    username: root</span>
<span class="line">    password: env(DB_PASSWORD)</span>
<span class="line">    timeout: 1</span>
<span class="line">    readtime: 1</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="应用环境" tabindex="-1"><a class="header-anchor" href="#应用环境"><span>应用环境</span></a></h2><p>gob 启动应用的默认应用环境为 dev。</p><p>你可以通过设置 <code>.env</code> 文件中的 <code>APP_ENV</code> 设置应用环境。</p><p>应用环境建议选择：</p><ul><li>dev // 开发使用</li><li>prod // 线上使用</li><li>test // 测试环境</li></ul><p>应用环境对应配置的文件夹，配置服务会去对应应用环境的文件夹中寻找配置。</p><p>比如应用环境为 dev，在代码中使用</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line">configService <span class="token operator">:=</span> container<span class="token punctuation">.</span><span class="token function">MustMake</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>ConfigKey<span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token punctuation">(</span>contract<span class="token punctuation">.</span>Config<span class="token punctuation">)</span></span>
<span class="line">url <span class="token operator">:=</span> configService<span class="token punctuation">.</span><span class="token function">GetString</span><span class="token punctuation">(</span><span class="token string">&quot;app.url&quot;</span><span class="token punctuation">)</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div></div></div><p>查找文件为：<code>config/dev/app.yaml</code></p><p>通过命令<code>./gob env</code>也可以获取当前应用环境：</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line"><span class="token operator">&gt;</span> ./gob <span class="token function">env</span></span>
<span class="line">environment: dev</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="命令" tabindex="-1"><a class="header-anchor" href="#命令"><span>命令</span></a></h2><p>相关的命令详见：<a href="../command/env">env</a></p>`,20),t=[l];function c(p,d){return s(),e("div",null,t)}const r=n(i,[["render",c],["__file","env.html.vue"]]),u=JSON.parse('{"path":"/guide/env.html","title":"环境变量","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"环境变量","description":null},"headers":[{"level":2,"title":"设置","slug":"设置","link":"#设置","children":[]},{"level":2,"title":"应用环境","slug":"应用环境","link":"#应用环境","children":[]},{"level":2,"title":"命令","slug":"命令","link":"#命令","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/env.md"}');export{r as comp,u as data};
