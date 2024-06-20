import{_ as n,c as s,o as a,b as e}from"./app-3nXj21aS.js";const p={},t=e(`<h1 id="gob-orm" tabindex="-1"><a class="header-anchor" href="#gob-orm"><span>gob:orm</span></a></h1><h2 id="服务介绍" tabindex="-1"><a class="header-anchor" href="#服务介绍"><span>服务介绍：</span></a></h2><p>提供ORM服务的服务，可以用于获取数据库连接，获取表结构等。</p><h2 id="支持命令-无" tabindex="-1"><a class="header-anchor" href="#支持命令-无"><span>支持命令：无</span></a></h2><h2 id="支持配置" tabindex="-1"><a class="header-anchor" href="#支持配置"><span>支持配置：</span></a></h2><p>使用之前需要确保已经正确配置了redis服务。</p><p>配置文件为 <code>config/[env]/database.yaml</code>，以下是一个配置的例子：</p><div class="language-yaml line-numbers-mode" data-highlighter="prismjs" data-ext="yml" data-title="yml"><pre class="language-yaml"><code><span class="line"><span class="token comment">##### mysql连接配置</span></span>
<span class="line"><span class="token comment">#mysql:</span></span>
<span class="line"><span class="token comment">#  hostname: 127.0.0.1</span></span>
<span class="line"><span class="token comment">#  username: test</span></span>
<span class="line"><span class="token comment">#  password: env(DB_PASSWORD)</span></span>
<span class="line"><span class="token comment">#  timeout: 1</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">##### 分组下通用配置</span></span>
<span class="line"></span>
<span class="line"><span class="token key atrule">conn_max_idle</span><span class="token punctuation">:</span> <span class="token number">10</span> <span class="token comment"># 通用配置，连接池最大空闲连接数</span></span>
<span class="line"><span class="token key atrule">conn_max_open</span><span class="token punctuation">:</span> <span class="token number">100</span> <span class="token comment"># 通用配置，连接池最大连接数</span></span>
<span class="line"><span class="token key atrule">conn_max_lifetime</span><span class="token punctuation">:</span> 1h <span class="token comment"># 通用配置，连接数最大生命周期</span></span>
<span class="line"><span class="token key atrule">protocol</span><span class="token punctuation">:</span> tcp <span class="token comment"># 通用配置，传输协议</span></span>
<span class="line"><span class="token key atrule">loc</span><span class="token punctuation">:</span> Local <span class="token comment"># 通用配置，时区</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">##### 默认分组下的mysql连接配置</span></span>
<span class="line"><span class="token key atrule">default</span><span class="token punctuation">:</span></span>
<span class="line">  <span class="token key atrule">driver</span><span class="token punctuation">:</span> mysql <span class="token comment"># 连接驱动</span></span>
<span class="line">  <span class="token key atrule">dsn</span><span class="token punctuation">:</span> <span class="token string">&quot;&quot;</span> <span class="token comment"># dsn，如果设置了dsn, 以下的所有设置都不生效</span></span>
<span class="line">  <span class="token key atrule">host</span><span class="token punctuation">:</span> localhost <span class="token comment"># ip地址</span></span>
<span class="line">  <span class="token key atrule">port</span><span class="token punctuation">:</span> <span class="token number">3306</span> <span class="token comment"># 端口</span></span>
<span class="line">  <span class="token key atrule">database</span><span class="token punctuation">:</span> demo <span class="token comment"># 数据库</span></span>
<span class="line">  <span class="token key atrule">username</span><span class="token punctuation">:</span> demo <span class="token comment"># 用户名</span></span>
<span class="line">  <span class="token key atrule">password</span><span class="token punctuation">:</span> <span class="token string">&quot;123456&quot;</span> <span class="token comment"># 密码</span></span>
<span class="line">  <span class="token key atrule">allow_native_passwords</span><span class="token punctuation">:</span> <span class="token boolean important">true</span></span>
<span class="line">  <span class="token key atrule">charset</span><span class="token punctuation">:</span> utf8mb4 <span class="token comment"># 字符集</span></span>
<span class="line">  <span class="token key atrule">collation</span><span class="token punctuation">:</span> utf8mb4_unicode_ci <span class="token comment"># 字符序</span></span>
<span class="line">  <span class="token key atrule">timeout</span><span class="token punctuation">:</span> 10s <span class="token comment"># 连接超时</span></span>
<span class="line">  <span class="token key atrule">read_timeout</span><span class="token punctuation">:</span> 2s <span class="token comment"># 读超时</span></span>
<span class="line">  <span class="token key atrule">write_timeout</span><span class="token punctuation">:</span> 2s <span class="token comment"># 写超时</span></span>
<span class="line">  <span class="token key atrule">parse_time</span><span class="token punctuation">:</span> <span class="token boolean important">true</span> <span class="token comment"># 是否解析时间</span></span>
<span class="line">  <span class="token key atrule">protocol</span><span class="token punctuation">:</span> tcp <span class="token comment"># 传输协议</span></span>
<span class="line">  <span class="token key atrule">loc</span><span class="token punctuation">:</span> Local <span class="token comment"># 时区</span></span>
<span class="line">  <span class="token key atrule">conn_max_idle</span><span class="token punctuation">:</span> <span class="token number">10</span> <span class="token comment"># 连接池最大空闲连接数</span></span>
<span class="line">  <span class="token key atrule">conn_max_open</span><span class="token punctuation">:</span> <span class="token number">20</span> <span class="token comment"># 连接池最大连接数</span></span>
<span class="line">  <span class="token key atrule">conn_max_lifetime</span><span class="token punctuation">:</span> 1h <span class="token comment"># 连接数最大生命周期</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">##### 默认分组下的sqlite连接配置</span></span>
<span class="line"><span class="token comment">#default:</span></span>
<span class="line"><span class="token comment">#  driver: sqlite # 连接驱动</span></span>
<span class="line"><span class="token comment">#  dsn: D:\\dev-project\\0.demo\\out\\box.db</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="提供方法" tabindex="-1"><a class="header-anchor" href="#提供方法"><span>提供方法：</span></a></h2><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token keyword">type</span> ORM <span class="token keyword">interface</span> <span class="token punctuation">{</span></span>
<span class="line">	<span class="token comment">// 获取 DB</span></span>
<span class="line">	<span class="token function">GetDB</span><span class="token punctuation">(</span>option <span class="token operator">...</span>DBOption<span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token operator">*</span>gorm<span class="token punctuation">.</span>DB<span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// CanConnect 是否可以连接</span></span>
<span class="line">	<span class="token function">CanConnect</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> db <span class="token operator">*</span>gorm<span class="token punctuation">.</span>DB<span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token builtin">bool</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line"></span>
<span class="line">	<span class="token comment">// Table 相关</span></span>
<span class="line">	<span class="token function">GetTables</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> db <span class="token operator">*</span>gorm<span class="token punctuation">.</span>DB<span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line">	<span class="token function">HasTable</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> db <span class="token operator">*</span>gorm<span class="token punctuation">.</span>DB<span class="token punctuation">,</span> table <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token builtin">bool</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line">	<span class="token function">GetTableColumns</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> db <span class="token operator">*</span>gorm<span class="token punctuation">.</span>DB<span class="token punctuation">,</span> table <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span>TableColumn<span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span></span>
<span class="line"><span class="token punctuation">}</span></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,10),l=[t];function c(o,i){return a(),s("div",null,l)}const r=n(p,[["render",c],["__file","orm.html.vue"]]),m=JSON.parse('{"path":"/provider/orm.html","title":"gob:orm","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"gob:orm","description":null},"headers":[{"level":2,"title":"服务介绍：","slug":"服务介绍","link":"#服务介绍","children":[]},{"level":2,"title":"支持命令：无","slug":"支持命令-无","link":"#支持命令-无","children":[]},{"level":2,"title":"支持配置：","slug":"支持配置","link":"#支持配置","children":[]},{"level":2,"title":"提供方法：","slug":"提供方法","link":"#提供方法","children":[]}],"git":{"updatedTime":1718886161000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":5}]},"filePathRelative":"provider/orm.md"}');export{r as comp,m as data};
