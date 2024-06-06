import{_ as e,c as n,o as a,a as s}from"./app-yBQH3CcA.js";const t={},l=s(`<h1 id="gob-kernel" tabindex="-1"><a class="header-anchor" href="#gob-kernel"><span>gob:kernel</span></a></h1><h2 id="服务介绍" tabindex="-1"><a class="header-anchor" href="#服务介绍"><span>服务介绍：</span></a></h2><p>提供框架最核心的结构，包括 http 和 grpc 的 Engine 结构。</p><h2 id="支持命令-无" tabindex="-1"><a class="header-anchor" href="#支持命令-无"><span>支持命令：无</span></a></h2><h2 id="支持配置-无" tabindex="-1"><a class="header-anchor" href="#支持配置-无"><span>支持配置：无</span></a></h2><h2 id="使用方法" tabindex="-1"><a class="header-anchor" href="#使用方法"><span>使用方法</span></a></h2><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">type</span> Kernel <span class="token keyword">interface</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token comment">// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine</span></span>
<span class="line">	<span class="token function">HttpEngine</span><span class="token punctuation">(</span><span class="token punctuation">)</span> http<span class="token punctuation">.</span>Handler</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,7),i=[l];function r(c,p){return a(),n("div",null,i)}const d=e(t,[["render",r],["__file","kernel.html.vue"]]),h=JSON.parse('{"path":"/provider/kernel.html","title":"gob:kernel","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"gob:kernel","description":null},"headers":[{"level":2,"title":"服务介绍：","slug":"服务介绍","link":"#服务介绍","children":[]},{"level":2,"title":"支持命令：无","slug":"支持命令-无","link":"#支持命令-无","children":[]},{"level":2,"title":"支持配置：无","slug":"支持配置-无","link":"#支持配置-无","children":[]},{"level":2,"title":"使用方法","slug":"使用方法","link":"#使用方法","children":[]}],"git":{"updatedTime":1717674707000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":2}]},"filePathRelative":"provider/kernel.md"}');export{d as comp,h as data};
