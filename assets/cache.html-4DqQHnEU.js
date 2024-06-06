import{_ as n,c as s,o as a,a as t}from"./app-yBQH3CcA.js";const e={},p=t(`<h1 id="gob-cache" tabindex="-1"><a class="header-anchor" href="#gob-cache"><span>gob:cache</span></a></h1><h2 id="服务介绍" tabindex="-1"><a class="header-anchor" href="#服务介绍"><span>服务介绍：</span></a></h2><p>cache 服务提供丰富的接口，可以通过接口来操作缓存，目前支持的缓存驱动有两种：</p><ul><li>redis</li><li>memory</li></ul><h2 id="支持命令-无" tabindex="-1"><a class="header-anchor" href="#支持命令-无"><span>支持命令：无</span></a></h2><h2 id="支持配置" tabindex="-1"><a class="header-anchor" href="#支持配置"><span>支持配置：</span></a></h2><p>通过配置文件 <code>config/[env]/cache.yaml</code> 可以配置缓存服务的驱动和参数，如下是一个配置示例：</p><div class="language-yaml line-numbers-mode" data-highlighter="prismjs" data-ext="yml" data-title="yml"><pre class="language-yaml"><code><span class="line"><span class="token key atrule">driver</span><span class="token punctuation">:</span> memory 	<span class="token comment"># 连接驱动，可选 redis/memory</span></span>
<span class="line"><span class="token punctuation">...</span> 			<span class="token comment"># 如果 driver: redis，则可配置项与redis服务一致</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="使用方法" tabindex="-1"><a class="header-anchor" href="#使用方法"><span>使用方法</span></a></h2><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">type</span> CacheService <span class="token keyword">interface</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token comment">// Get 获取某个key对应的值</span></span>
<span class="line">	<span class="token function">Get</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token builtin">string</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line">	<span class="token comment">// GetObj 获取某个key对应的对象, 对象必须实现 https://pkg.go.dev/encoding#BinaryUnMarshaler</span></span>
<span class="line">	<span class="token function">GetObj</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> model <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// GetMany 获取某些key对应的值</span></span>
<span class="line">	<span class="token function">GetMany</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> keys <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// Set 设置某个key和值到缓存，带超时时间</span></span>
<span class="line">	<span class="token function">Set</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> val <span class="token builtin">string</span><span class="token punctuation">,</span> timeout time<span class="token punctuation">.</span>Duration<span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// SetObj 设置某个key和对象到缓存, 对象必须实现 https://pkg.go.dev/encoding#BinaryMarshaler</span></span>
<span class="line">	<span class="token function">SetObj</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> val <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> timeout time<span class="token punctuation">.</span>Duration<span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// SetMany 设置多个key和值到缓存</span></span>
<span class="line">	<span class="token function">SetMany</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> data <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> timeout time<span class="token punctuation">.</span>Duration<span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// SetForever 设置某个key和值到缓存，不带超时时间</span></span>
<span class="line">	<span class="token function">SetForever</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> val <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// SetForeverObj 设置某个key和对象到缓存，不带超时时间，对象必须实现 https://pkg.go.dev/encoding#BinaryMarshaler</span></span>
<span class="line">	<span class="token function">SetForeverObj</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> val <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// SetTTL 设置某个key的超时时间</span></span>
<span class="line">	<span class="token function">SetTTL</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> timeout time<span class="token punctuation">.</span>Duration<span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// GetTTL 获取某个key的超时时间</span></span>
<span class="line">	<span class="token function">GetTTL</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>time<span class="token punctuation">.</span>Duration<span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// Remember 实现缓存的Cache-Aside模式, 先去缓存中根据key获取对象，如果有的话，返回，如果没有，调用RememberFunc 生成</span></span>
<span class="line">	<span class="token function">Remember</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> timeout time<span class="token punctuation">.</span>Duration<span class="token punctuation">,</span> rememberFunc RememberFunc<span class="token punctuation">,</span> model <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// Calc 往key对应的值中增加step计数</span></span>
<span class="line">	<span class="token function">Calc</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">,</span> step <span class="token builtin">int64</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token builtin">int64</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line">	<span class="token comment">// Increment 往key对应的值中增加1</span></span>
<span class="line">	<span class="token function">Increment</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token builtin">int64</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line">	<span class="token comment">// Decrement 往key对应的值中减去1</span></span>
<span class="line">	<span class="token function">Decrement</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token builtin">int64</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// Del 删除某个key</span></span>
<span class="line">	<span class="token function">Del</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> key <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line">	<span class="token comment">// DelMany 删除某些key</span></span>
<span class="line">	<span class="token function">DelMany</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> keys <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,10),c=[p];function i(l,o){return a(),s("div",null,c)}const k=n(e,[["render",i],["__file","cache.html.vue"]]),r=JSON.parse('{"path":"/provider/cache.html","title":"gob:cache","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"gob:cache","description":null},"headers":[{"level":2,"title":"服务介绍：","slug":"服务介绍","link":"#服务介绍","children":[]},{"level":2,"title":"支持命令：无","slug":"支持命令-无","link":"#支持命令-无","children":[]},{"level":2,"title":"支持配置：","slug":"支持配置","link":"#支持配置","children":[]},{"level":2,"title":"使用方法","slug":"使用方法","link":"#使用方法","children":[]}],"git":{"updatedTime":1717674707000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":2}]},"filePathRelative":"provider/cache.md"}');export{k as comp,r as data};