import{_ as e,o as a,c as n,a as s}from"./app-SHUsDq-_.js";const c={},i=s(`<h1 id="gob-cache" tabindex="-1"><a class="header-anchor" href="#gob-cache"><span>gob:cache</span></a></h1><h2 id="说明" tabindex="-1"><a class="header-anchor" href="#说明"><span>说明</span></a></h2><p>gob:cache 是直接微框架提供缓存服务，目前支持的缓存驱动有两种：</p><ul><li>redis</li><li>memory</li></ul><p>通过配置文件 <code>config/[env]/cache.yaml</code> 可以配置缓存服务的驱动和参数，如下是一个配置示例：</p><div class="language-yaml line-numbers-mode" data-ext="yml" data-title="yml"><pre class="language-yaml"><code><span class="token comment">#driver: redis # 连接驱动</span>
<span class="token comment">#host: 127.0.0.1 # ip地址</span>
<span class="token comment">#port: 6379 # 端口</span>
<span class="token comment">#db: 0 #db</span>
<span class="token comment">#timeout: 10s # 连接超时</span>
<span class="token comment">#read_timeout: 2s # 读超时</span>
<span class="token comment">#write_timeout: 2s # 写超时</span>
<span class="token comment">#</span>
<span class="token key atrule">driver</span><span class="token punctuation">:</span> memory <span class="token comment"># 连接驱动</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="使用方法" tabindex="-1"><a class="header-anchor" href="#使用方法"><span>使用方法</span></a></h2><p>cache 服务提供丰富的接口，可以通过接口来操作缓存，如下是接口定义：</p><div class="language-go line-numbers-mode" data-ext="go" data-title="go"><pre class="language-go"><code>

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div>`,9),l=[i];function d(t,o){return a(),n("div",null,l)}const m=e(c,[["render",d],["__file","cache.html.vue"]]);export{m as default};
