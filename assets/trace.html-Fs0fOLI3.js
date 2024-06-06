import{_ as n,c as a,o as s,a as e}from"./app-yBQH3CcA.js";const t={},c=e(`<h1 id="gob-trace" tabindex="-1"><a class="header-anchor" href="#gob-trace"><span>gob:trace</span></a></h1><h2 id="服务介绍" tabindex="-1"><a class="header-anchor" href="#服务介绍"><span>服务介绍：</span></a></h2><p>提供分布式链路追踪服务，可以用于跟踪分布式服务调用链路。</p><h2 id="支持命令-无" tabindex="-1"><a class="header-anchor" href="#支持命令-无"><span>支持命令：无</span></a></h2><h2 id="支持配置-无" tabindex="-1"><a class="header-anchor" href="#支持配置-无"><span>支持配置：无</span></a></h2><h2 id="使用方法" tabindex="-1"><a class="header-anchor" href="#使用方法"><span>使用方法</span></a></h2><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">type</span> Trace <span class="token keyword">interface</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token comment">// WithTrace register new trace to context</span></span>
<span class="line">	<span class="token function">WithTrace</span><span class="token punctuation">(</span>c context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> trace <span class="token operator">*</span>TraceContext<span class="token punctuation">)</span> context<span class="token punctuation">.</span>Context</span>
<span class="line">	<span class="token comment">// GetTrace From trace context</span></span>
<span class="line">	<span class="token function">GetTrace</span><span class="token punctuation">(</span>c context<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token operator">*</span>TraceContext</span>
<span class="line">	<span class="token comment">// NewTrace generate a new trace</span></span>
<span class="line">	<span class="token function">NewTrace</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token operator">*</span>TraceContext</span>
<span class="line">	<span class="token comment">// StartSpan generate cspan for child call</span></span>
<span class="line">	<span class="token function">StartSpan</span><span class="token punctuation">(</span>trace <span class="token operator">*</span>TraceContext<span class="token punctuation">)</span> <span class="token operator">*</span>TraceContext</span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ToMap traceContext to map for logger</span></span>
<span class="line">	<span class="token function">ToMap</span><span class="token punctuation">(</span>trace <span class="token operator">*</span>TraceContext<span class="token punctuation">)</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// ExtractHTTP GetTrace By Http</span></span>
<span class="line">	<span class="token function">ExtractHTTP</span><span class="token punctuation">(</span>req <span class="token operator">*</span>http<span class="token punctuation">.</span>Request<span class="token punctuation">)</span> <span class="token operator">*</span>TraceContext</span>
<span class="line">	<span class="token comment">// InjectHTTP Set Trace to Http</span></span>
<span class="line">	<span class="token function">InjectHTTP</span><span class="token punctuation">(</span>req <span class="token operator">*</span>http<span class="token punctuation">.</span>Request<span class="token punctuation">,</span> trace <span class="token operator">*</span>TraceContext<span class="token punctuation">)</span> <span class="token operator">*</span>http<span class="token punctuation">.</span>Request</span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,7),p=[c];function l(o,i){return s(),a("div",null,p)}const u=n(t,[["render",l],["__file","trace.html.vue"]]),d=JSON.parse('{"path":"/provider/trace.html","title":"gob:trace","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"gob:trace","description":null},"headers":[{"level":2,"title":"服务介绍：","slug":"服务介绍","link":"#服务介绍","children":[]},{"level":2,"title":"支持命令：无","slug":"支持命令-无","link":"#支持命令-无","children":[]},{"level":2,"title":"支持配置：无","slug":"支持配置-无","link":"#支持配置-无","children":[]},{"level":2,"title":"使用方法","slug":"使用方法","link":"#使用方法","children":[]}],"git":{"updatedTime":1717674707000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":2}]},"filePathRelative":"provider/trace.md"}');export{u as comp,d as data};
