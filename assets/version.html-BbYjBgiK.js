import{_ as s,c as n,o as e,b as a}from"./app-DRNFix5Q.js";const i={},l=a(`<h1 id="版本" tabindex="-1"><a class="header-anchor" href="#版本"><span>版本</span></a></h1><h2 id="命令" tabindex="-1"><a class="header-anchor" href="#命令"><span>命令</span></a></h2><p>相关的命令详见：<a href="../command/version">version</a></p><p>gob 提供了查询当前版本和获取最新版本日志的命令</p><h2 id="查询当前版本" tabindex="-1"><a class="header-anchor" href="#查询当前版本"><span>查询当前版本</span></a></h2><p>使用命令 <code>gob version</code></p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">&gt; ./gob version</span>
<span class="line">gob version: 1.0.0</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="获取最新的版本" tabindex="-1"><a class="header-anchor" href="#获取最新的版本"><span>获取最新的版本</span></a></h2><p>使用命令 <code>gob version list</code></p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">&gt; ./gob version list       </span>
<span class="line">===============前置条件检测===============</span>
<span class="line">gob源码从github.com中下载，正在检测到github.com的连接</span>
<span class="line">github.Rate{Limit:60, Remaining:59, Reset:github.Timestamp{2024-06-10 19:12:45 +0800 CST}}</span>
<span class="line">gob源码从github.com中下载，github.com的连接正常</span>
<span class="line">===============前置条件检测结束===============</span>
<span class="line"></span>
<span class="line">最新的1个版本</span>
<span class="line">-v0.1.11</span>
<span class="line">  发布时间：2024-06-10 08:56:12</span>
<span class="line">  修改说明：</span>
<span class="line">    集成初始化脚手架，可通过以下命令在本地构建应用：</span>
<span class="line"></span>
<span class="line">    使用 go install github.com/chenbihao/gob@latest 来安装 gob 命令。</span>
<span class="line"></span>
<span class="line">    运行初始化脚手架 gob new 并根据命令行互动输入对应的应用名与模块名。</span>
<span class="line"></span>
<span class="line">    进入对应的文件夹，使用 go mod tidy 安装相关依赖，</span>
<span class="line">    随后可以通过引用 github.com/chenbihao/gob/framework 来引用框架相关模块</span>
<span class="line"></span>
<span class="line">更多历史版本请参考 https://github.com/chenbihao/gob/releases</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,10),c=[l];function t(d,r){return e(),n("div",null,c)}const o=s(i,[["render",t],["__file","version.html.vue"]]),v=JSON.parse('{"path":"/guide/version.html","title":"版本","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"版本","description":null},"headers":[{"level":2,"title":"命令","slug":"命令","link":"#命令","children":[]},{"level":2,"title":"查询当前版本","slug":"查询当前版本","link":"#查询当前版本","children":[]},{"level":2,"title":"获取最新的版本","slug":"获取最新的版本","link":"#获取最新的版本","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/version.md"}');export{o as comp,v as data};
