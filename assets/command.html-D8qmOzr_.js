import{_ as e,r as l,c as i,a as n,e as s,d as c,w as t,b as d,o}from"./app-BSoEOH6w.js";const p={},r=n("h1",{id:"命令",tabindex:"-1"},[n("a",{class:"header-anchor",href:"#命令"},[n("span",null,"命令")])],-1),m=n("h2",{id:"指南",tabindex:"-1"},[n("a",{class:"header-anchor",href:"#指南"},[n("span",null,"指南")])],-1),u=d(`<p>gob 允许自定义命令，挂载到 gob 上。并且提供了<code>./gob command</code> 系列命令。</p><div class="language-bash line-numbers-mode" data-highlighter="prismjs" data-ext="sh" data-title="sh"><pre><code><span class="line"><span class="token operator">&gt;</span> ./gob <span class="token builtin class-name">command</span></span>
<span class="line">控制台命令相关</span>
<span class="line"></span>
<span class="line">Usage:</span>
<span class="line">  gob <span class="token builtin class-name">command</span> <span class="token punctuation">[</span>flags<span class="token punctuation">]</span></span>
<span class="line">  gob <span class="token builtin class-name">command</span> <span class="token punctuation">[</span>command<span class="token punctuation">]</span></span>
<span class="line"></span>
<span class="line">Available Commands:</span>
<span class="line">  list        列出所有控制台命令</span>
<span class="line">  new         创建一个控制台命令</span>
<span class="line"></span>
<span class="line">Flags:</span>
<span class="line">  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> <span class="token builtin class-name">command</span></span>
<span class="line"></span>
<span class="line">Use <span class="token string">&quot;gob command [command] --help&quot;</span> <span class="token keyword">for</span> <span class="token function">more</span> information about a command.</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="创建" tabindex="-1"><a class="header-anchor" href="#创建"><span>创建</span></a></h2><p>创建一个新命令，可以使用 <code>./gob command new</code></p><p>这是一个交互式的命令行工具。</p><p>创建完成之后，会在应用的 app/console/command/ 目录下创建一个新的文件。</p><h2 id="自定义" tabindex="-1"><a class="header-anchor" href="#自定义"><span>自定义</span></a></h2><p>gob 中的命令使用的是 cobra 库。 https://github.com/spf13/cobra</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre><code><span class="line">package command</span>
<span class="line"></span>
<span class="line">import (</span>
<span class="line">        &quot;fmt&quot;</span>
<span class="line"></span>
<span class="line">        &quot;github.com/chenbihao/gob/framework/cobra&quot;</span>
<span class="line">        &quot;github.com/chenbihao/gob/framework/command/util&quot;</span>
<span class="line">)</span>
<span class="line"></span>
<span class="line">var TestCommand = &amp;cobra.Command{</span>
<span class="line">        Use:   &quot;test&quot;,</span>
<span class="line">        Short: &quot;test&quot;,</span>
<span class="line">        RunE: func(c *cobra.Command, args []string) error {</span>
<span class="line">                container := util.GetContainer(c.Root())</span>
<span class="line">                fmt.Println(container)</span>
<span class="line">                return nil</span>
<span class="line">        },</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>基本上，我们要求实现</p><ul><li>Use // 命令行的关键字</li><li>Short // 命令行的简短说明</li><li>RunE // 命令行实际运行的程序</li></ul><p>更多配置和参数可以参考 <a href="https://github.com/spf13/cobra" target="_blank" rel="noopener noreferrer">cobra 的 github 页面</a></p><h2 id="挂载" tabindex="-1"><a class="header-anchor" href="#挂载"><span>挂载</span></a></h2><p>编写完自定义命令后，请记得挂载到 <code>console/kernel.go</code> 中。</p><div class="language-golang line-numbers-mode" data-highlighter="prismjs" data-ext="golang" data-title="golang"><pre><code><span class="line">func RunCommand(container framework.Container) error {</span>
<span class="line">	var rootCmd = &amp;cobra.Command{</span>
<span class="line">		Use:   &quot;gob&quot;,</span>
<span class="line">		Short: &quot;main&quot;,</span>
<span class="line">		Long:  &quot;gob commands&quot;,</span>
<span class="line">	}</span>
<span class="line"></span>
<span class="line">	ctx := commandUtil.RegiestContainer(container, rootCmd)</span>
<span class="line">	gobCommand.AddKernelCommands(rootCmd)</span>
<span class="line">	rootCmd.AddCommand(command.DemoCommand)</span>
<span class="line">	return rootCmd.ExecuteContext(ctx)</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,15);function v(b,h){const a=l("RouteLink");return o(),i("div",null,[r,m,n("p",null,[s("相关的命令详见："),c(a,{to:"/command/cmd.html"},{default:t(()=>[s("command")]),_:1})]),u])}const k=e(p,[["render",v],["__file","command.html.vue"]]),f=JSON.parse('{"path":"/guide/command.html","title":"命令","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"命令","description":null},"headers":[{"level":2,"title":"指南","slug":"指南","link":"#指南","children":[]},{"level":2,"title":"创建","slug":"创建","link":"#创建","children":[]},{"level":2,"title":"自定义","slug":"自定义","link":"#自定义","children":[]},{"level":2,"title":"挂载","slug":"挂载","link":"#挂载","children":[]}],"git":{"updatedTime":1718016868000,"contributors":[{"name":"被水淹没","email":"994523036@qq.com","commits":1},{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":1}]},"filePathRelative":"guide/command.md"}');export{k as comp,f as data};
