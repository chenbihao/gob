import{_ as n,c as e,o as a,a as l}from"./app-MeNRN-Ja.js";const s={},i=l(`<h1 id="编译" tabindex="-1"><a class="header-anchor" href="#编译"><span>编译</span></a></h1><h2 id="命令" tabindex="-1"><a class="header-anchor" href="#命令"><span>命令</span></a></h2><p>应用分为前端（frontend）和后端（backend），所以编译也分为三类</p><ul><li>编译前端</li><li>编译后端</li><li>自编译</li><li>同时编译</li></ul><p>相关的命令详见：<a href="../command/build">build</a></p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre><code><span class="line">&gt; ./gob build</span>
<span class="line">编译相关命令</span>
<span class="line"></span>
<span class="line">Usage:</span>
<span class="line">  gob build [flags]</span>
<span class="line">  gob build [command]</span>
<span class="line"></span>
<span class="line">Available Commands:</span>
<span class="line">  all         同时编译前端和后端</span>
<span class="line">  backend     使用 go 编译后端</span>
<span class="line">  frontend    使用 npm 编译前端</span>
<span class="line">  self        编译 gob 命令</span>
<span class="line"></span>
<span class="line">Flags:</span>
<span class="line">  -h, --help   help for build</span>
<span class="line"></span>
<span class="line">Use &quot;gob build [command] --help&quot; for more information about a command.</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="编译前端" tabindex="-1"><a class="header-anchor" href="#编译前端"><span>编译前端</span></a></h2><p>要求当前编译机器安装 npm 软件，并且当前项目已经运行了 npm install，安装完成前端依赖。</p><p>运行命令 <code>./gob build frontend</code></p><p>编译后的前端文件在 dist 目录中</p><p>实际上 build 就是调用 <code>npm build</code> 来编译前端项目。</p><h2 id="编译后端" tabindex="-1"><a class="header-anchor" href="#编译后端"><span>编译后端</span></a></h2><p>要求当前编译机器安装 go 软件，版本 &gt; 1.3。</p><p>运行命令： <code>./gob build backend</code></p><p>在项目根目录下就看到生成的可执行文件 gob。 后续可以通过 ./gob 直接运行。</p><h2 id="自编译" tabindex="-1"><a class="header-anchor" href="#自编译"><span>自编译</span></a></h2><p>在项目根目录下，gob 可以通过 gob 命令编译出 gob 命令自己。</p><p>运行命令 <code>gob build self</code></p><p>在项目根目录下就看到生成的可执行文件 gob。 后续可以通过 ./gob 直接运行。</p><blockquote><p>其实自编译和后端编译是同样效果，但是为了命令语义化，增加了自编译的命令。</p></blockquote><h2 id="同时编译" tabindex="-1"><a class="header-anchor" href="#同时编译"><span>同时编译</span></a></h2><p>顾名思义，同时编译前端和后端，命令为 <code>./gob build all</code></p>`,22),d=[i];function c(p,o){return a(),e("div",null,d)}const r=n(s,[["render",c],["__file","build.html.vue"]]),b=JSON.parse('{"path":"/guide/build.html","title":"编译","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"编译","description":null},"headers":[{"level":2,"title":"命令","slug":"命令","link":"#命令","children":[]},{"level":2,"title":"编译前端","slug":"编译前端","link":"#编译前端","children":[]},{"level":2,"title":"编译后端","slug":"编译后端","link":"#编译后端","children":[]},{"level":2,"title":"自编译","slug":"自编译","link":"#自编译","children":[]},{"level":2,"title":"同时编译","slug":"同时编译","link":"#同时编译","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/build.md"}');export{r as comp,b as data};
