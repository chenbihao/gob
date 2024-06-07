import{_ as s,c as n,o as a,b as e}from"./app-Dprcp3s0.js";const l={},p=e(`<h1 id="部署命令" tabindex="-1"><a class="header-anchor" href="#部署命令"><span>部署命令</span></a></h1><h2 id="命令介绍" tabindex="-1"><a class="header-anchor" href="#命令介绍"><span>命令介绍：</span></a></h2><p>部署命令</p><h2 id="前置需求" tabindex="-1"><a class="header-anchor" href="#前置需求"><span>前置需求：</span></a></h2><p>app</p><h2 id="支持命令" tabindex="-1"><a class="header-anchor" href="#支持命令"><span>支持命令：</span></a></h2><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line">./gob deploy frontend<span class="token variable"><span class="token variable">\`</span>	// 部署前端</span>
<span class="line">	<span class="token parameter variable">-s</span> --skip-build     	// 跳过前端构建</span>
<span class="line">./gob deploy backend<span class="token variable">\`</span></span>	// 部署后端</span>
<span class="line">./gob deploy all<span class="token variable"><span class="token variable">\`</span>		// 同时部署前后端</span>
<span class="line">	<span class="token parameter variable">-s</span> --skip-build     	// 跳过前端构建</span>
<span class="line">./gob deploy rollback<span class="token variable">\`</span></span>	// 部署回滚</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="支持配置" tabindex="-1"><a class="header-anchor" href="#支持配置"><span>支持配置：</span></a></h2><p><code>deploy.yaml</code> 支持配置：</p><div class="language-yaml line-numbers-mode" data-highlighter="prismjs" data-ext="yml" data-title="yml"><pre class="language-yaml"><code><span class="line"><span class="token key atrule">connections</span><span class="token punctuation">:</span> <span class="token comment"># 要自动化部署的连接</span></span>
<span class="line">  <span class="token punctuation">-</span> ssh.web<span class="token punctuation">-</span>key</span>
<span class="line"></span>
<span class="line"><span class="token key atrule">remote_folder</span><span class="token punctuation">:</span> <span class="token string">&quot;/home/demo/deploy/&quot;</span>  <span class="token comment"># 远端的部署文件夹</span></span>
<span class="line"></span>
<span class="line"><span class="token key atrule">frontend</span><span class="token punctuation">:</span> <span class="token comment"># 前端部署配置</span></span>
<span class="line">  <span class="token key atrule">pre_action</span><span class="token punctuation">:</span> <span class="token comment"># 部署前置命令</span></span>
<span class="line">	<span class="token punctuation">-</span> <span class="token string">&quot;pwd&quot;</span></span>
<span class="line">  <span class="token key atrule">post_action</span><span class="token punctuation">:</span> <span class="token comment"># 部署后置命令</span></span>
<span class="line">	<span class="token punctuation">-</span> <span class="token string">&quot;pwd&quot;</span></span>
<span class="line"></span>
<span class="line"><span class="token key atrule">backend</span><span class="token punctuation">:</span> <span class="token comment"># 后端部署配置</span></span>
<span class="line">  <span class="token key atrule">goos</span><span class="token punctuation">:</span> linux <span class="token comment"># 部署目标操作系统</span></span>
<span class="line">  <span class="token key atrule">goarch</span><span class="token punctuation">:</span> amd64 <span class="token comment"># 部署目标cpu架构</span></span>
<span class="line">  <span class="token key atrule">pre_action</span><span class="token punctuation">:</span> <span class="token comment"># 部署前置命令</span></span>
<span class="line">	<span class="token punctuation">-</span> <span class="token string">&quot;rm /home/demo/deploy/gob&quot;</span></span>
<span class="line">  <span class="token key atrule">post_action</span><span class="token punctuation">:</span> <span class="token comment"># 部署后置命令</span></span>
<span class="line">	<span class="token punctuation">-</span> <span class="token string">&quot;chmod 777 /home/demo/deploy/gob&quot;</span></span>
<span class="line">	<span class="token punctuation">-</span> <span class="token string">&quot;/home/demo/deploy/gob app restart&quot;</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>ssh 支持配置：详见 contract/redis.go</p><h2 id="使用方法" tabindex="-1"><a class="header-anchor" href="#使用方法"><span>使用方法：</span></a></h2>`,12),t=[p];function i(c,o){return a(),n("div",null,t)}const r=s(l,[["render",i],["__file","deploy.html.vue"]]),u=JSON.parse('{"path":"/command/deploy.html","title":"部署命令","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"部署命令","description":null},"headers":[{"level":2,"title":"命令介绍：","slug":"命令介绍","link":"#命令介绍","children":[]},{"level":2,"title":"前置需求：","slug":"前置需求","link":"#前置需求","children":[]},{"level":2,"title":"支持命令：","slug":"支持命令","link":"#支持命令","children":[]},{"level":2,"title":"支持配置：","slug":"支持配置","link":"#支持配置","children":[]},{"level":2,"title":"使用方法：","slug":"使用方法","link":"#使用方法","children":[]}],"git":{"updatedTime":1717776942000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"command/deploy.md"}');export{r as comp,u as data};
