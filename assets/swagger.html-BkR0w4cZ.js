import{_ as s,c as n,o as a,b as e}from"./app-CvecxDTg.js";const l={},i=e(`<h1 id="swagger" tabindex="-1"><a class="header-anchor" href="#swagger"><span>swagger</span></a></h1><h2 id="命令" tabindex="-1"><a class="header-anchor" href="#命令"><span>命令</span></a></h2><p>相关的命令详见：<a href="../command/swagger">swagger</a></p><p>gob 使用 <a href="https://github.com/swaggo/swag" target="_blank" rel="noopener noreferrer">swaggo</a> 集成了 swagger 生成和服务项目。</p><p>并且封装了 <code>./gob swagger</code> 命令。</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre class="language-bash"><code><span class="line"><span class="token operator">&gt;</span> ./gob swagger</span>
<span class="line">swagger对应命令</span>
<span class="line"></span>
<span class="line">Usage:</span>
<span class="line">  gob swagger <span class="token punctuation">[</span>flags<span class="token punctuation">]</span></span>
<span class="line">  gob swagger <span class="token punctuation">[</span>command<span class="token punctuation">]</span></span>
<span class="line"></span>
<span class="line">Available Commands:</span>
<span class="line">  gen         生成对应的swagger文件, contain swagger.yaml, doc.go</span>
<span class="line"></span>
<span class="line">Flags:</span>
<span class="line">  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> swagger</span>
<span class="line"></span>
<span class="line">Use <span class="token string">&quot;gob swagger [command] --help&quot;</span> <span class="token keyword">for</span> <span class="token function">more</span> information about a command.</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="注释" tabindex="-1"><a class="header-anchor" href="#注释"><span>注释</span></a></h2><p>gob 使用 <a href="https://github.com/swaggo/swag" target="_blank" rel="noopener noreferrer">swaggo</a> 来实现注释生成 swagger 功能。</p><p>全局注释在文件 <code>app/http/swagger.go</code> 中:</p><div class="language-go line-numbers-mode" data-highlighter="prismjs" data-ext="go" data-title="go"><pre class="language-go"><code><span class="line"><span class="token comment">// Package http API.</span></span>
<span class="line"><span class="token comment">// @title gob</span></span>
<span class="line"><span class="token comment">// @version 0.1.11</span></span>
<span class="line"><span class="token comment">// @description gob框架</span></span>
<span class="line"><span class="token comment">// @termsOfService https://github.com/swaggo/swag</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">// @contact.name chenbihao</span></span>
<span class="line"><span class="token comment">// @contact.email chenbihao@foxmail.com</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">// @license.name Apache 2.0</span></span>
<span class="line"><span class="token comment">// @license.url http://www.apache.org/licenses/LICENSE-2.0.html</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">// @BasePath /</span></span>
<span class="line"><span class="token comment">// @query.collection.format multi</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">// @securityDefinitions.basic BasicAuth</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">// @securityDefinitions.apikey ApiKeyAuth</span></span>
<span class="line"><span class="token comment">// @in header</span></span>
<span class="line"><span class="token comment">// @name Authorization</span></span>
<span class="line"></span>
<span class="line"><span class="token comment">// @x-extension-openapi {&quot;example&quot;: &quot;value on a json format&quot;}</span></span>
<span class="line"></span>
<span class="line"><span class="token keyword">package</span> http</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>接口注释请写在各自模块的 <code>api.go</code> 中</p><div class="language-golang line-numbers-mode" data-highlighter="prismjs" data-ext="golang" data-title="golang"><pre class="language-golang"><code><span class="line">// Demo godoc</span>
<span class="line">// @Summary 获取所有用户</span>
<span class="line">// @Description 获取所有用户</span>
<span class="line">// @Produce  json</span>
<span class="line">// @Tags demo</span>
<span class="line">// @Success 200 array []UserDTO</span>
<span class="line">// @Router /demo/demo [get]</span>
<span class="line">func (api *DemoApi) Demo(c *gin.Context) {</span>
<span class="line">	users := api.service.GetUsers()</span>
<span class="line">	usersDTO := UserModelsToUserDTOs(users)</span>
<span class="line">	c.JSON(200, usersDTO)</span>
<span class="line">}</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>swagger 注释的格式和关键词可以参考：<a href="https://github.com/swaggo/swag" target="_blank" rel="noopener noreferrer">swaggo</a></p><h2 id="生成" tabindex="-1"><a class="header-anchor" href="#生成"><span>生成</span></a></h2><p>使用命令 <code>./gob swagger gen</code></p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">&gt; ./gob swagger gen</span>
<span class="line">2024/06/10 18:35:22 Generate swagger docs....</span>
<span class="line">2024/06/10 18:35:22 Generate general API Info, search dir:D:\\DevProjects\\自己库\\gob\\app\\http</span>
<span class="line">2024/06/10 18:35:23 Generating demo.UserDTO</span>
<span class="line">2024/06/10 18:35:23 create docs.go at D:\\DevProjects\\自己库\\gob\\app\\http\\swagger/docs.go</span>
<span class="line">2024/06/10 18:35:23 create swagger.json at D:\\DevProjects\\自己库\\gob\\app\\http\\swagger/swagger.json</span>
<span class="line">2024/06/10 18:35:23 create swagger.yaml at D:\\DevProjects\\自己库\\gob\\app\\http\\swagger/swagger.yaml</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>在目录 <code>app/http/swagger/</code> 下自动生成swagger相关文件。</p><h2 id="服务" tabindex="-1"><a class="header-anchor" href="#服务"><span>服务</span></a></h2><p>可以使用命令 <code>./gob swagger serve</code> 启动当前应用的 swagger ui 服务。</p><blockquote><p>如果你的 swagger 服务已经启动，更新 swagger 只需要重新运行 <code>./gob swagger gen</code> 就能更新。</p><p>因为 swagger 服务读取的是生成的 <code>swagger.json</code> 这个文件。</p></blockquote><p>服务端口，我们也可以通过配置文件 <code>config/[env]/swagger.yaml</code> 中的配置来配置swagger serve 启动的服务:</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">url: http://127.0.0.1:8069</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div></div></div>`,22),c=[i];function p(r,t){return a(),n("div",null,c)}const o=s(l,[["render",p],["__file","swagger.html.vue"]]),g=JSON.parse('{"path":"/guide/swagger.html","title":"swagger","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"swagger","description":null},"headers":[{"level":2,"title":"命令","slug":"命令","link":"#命令","children":[]},{"level":2,"title":"注释","slug":"注释","link":"#注释","children":[]},{"level":2,"title":"生成","slug":"生成","link":"#生成","children":[]},{"level":2,"title":"服务","slug":"服务","link":"#服务","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/swagger.md"}');export{o as comp,g as data};
