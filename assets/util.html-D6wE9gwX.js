import{_ as n,c as s,o as e,b as a}from"./app-3nXj21aS.js";const i={},l=a(`<h1 id="辅助函数" tabindex="-1"><a class="header-anchor" href="#辅助函数"><span>辅助函数</span></a></h1><p>提供一些辅助函数来帮助更好的进行开发。</p><h2 id="goroutine-相关" tabindex="-1"><a class="header-anchor" href="#goroutine-相关"><span>goroutine 相关</span></a></h2><h3 id="safego" tabindex="-1"><a class="header-anchor" href="#safego"><span>SafeGo</span></a></h3><p>SafeGo 这个函数，提供了一种goroutine安全的函数调用方式。主要适用于业务中需要进行开启异步goroutine业务逻辑调用的场景。</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">// SafeGo 进行安全的goroutine调用</span>
<span class="line">// 第一个参数是context接口，如果还实现了Container接口，且绑定了日志服务，则使用日志服务</span>
<span class="line">// 第二个参数是匿名函数handler, 进行最终的业务逻辑</span>
<span class="line">// SafeGo 函数并不会返回error，panic都会进入的日志服务</span>
<span class="line">func SafeGo(ctx context.Context, handler func())</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>调用方式参照如下的单元测试用例：</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line"></span>
<span class="line">func TestSafeGo(t *testing.T) {</span>
<span class="line">    container := tests.InitBaseContainer()</span>
<span class="line">    container.Bind(&amp;log.TestingLogProvider{})</span>
<span class="line"></span>
<span class="line">    ctx, _ := gin.CreateTestContext(httptest.NewRecorder())</span>
<span class="line">    goroutine.SafeGo(ctx, func() {</span>
<span class="line">        time.Sleep(1 * time.Second)</span>
<span class="line">        return</span>
<span class="line">    })</span>
<span class="line">    t.Log(&quot;safe go main start&quot;)</span>
<span class="line">    time.Sleep(2 * time.Second)</span>
<span class="line">    t.Log(&quot;safe go main end&quot;)</span>
<span class="line"></span>
<span class="line">    goroutine.SafeGo(ctx, func() {</span>
<span class="line">        time.Sleep(1 * time.Second)</span>
<span class="line">        panic(&quot;safe go test panic&quot;)</span>
<span class="line">    })</span>
<span class="line">    t.Log(&quot;safe go2 main start&quot;)</span>
<span class="line">    time.Sleep(2 * time.Second)</span>
<span class="line">    t.Log(&quot;safe go2 main end&quot;)</span>
<span class="line"></span>
<span class="line">}</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="safegoandwait" tabindex="-1"><a class="header-anchor" href="#safegoandwait"><span>SafeGoAndWait</span></a></h3><p>SafeGoAndWait 这个函数，提供安全的多并发调用方式。该函数等待所有函数都结束后才返回。</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line">// SafeGoAndWait 进行并发安全并行调用</span>
<span class="line">// 第一个参数是context接口，如果还实现了Container接口，且绑定了日志服务，则使用日志服务</span>
<span class="line">// 第二个参数是匿名函数handlers数组, 进行最终的业务逻辑</span>
<span class="line">// 返回handlers中任何一个错误（如果handlers中有业务逻辑返回错误）</span>
<span class="line">func SafeGoAndWait(ctx context.Context, handlers ...func() error) error</span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>调用方式参照如下的单元测试用例：</p><div class="language-text line-numbers-mode" data-highlighter="prismjs" data-ext="text" data-title="text"><pre class="language-text"><code><span class="line"></span>
<span class="line">func TestSafeGoAndWait(t *testing.T) {</span>
<span class="line">    container := tests.InitBaseContainer()</span>
<span class="line">    container.Bind(&amp;log.TestingLogProvider{})</span>
<span class="line"></span>
<span class="line">    errStr := &quot;safe go test error&quot;</span>
<span class="line">    t.Log(&quot;safe go and wait start&quot;, time.Now().String())</span>
<span class="line">    ctx, _ := gin.CreateTestContext(httptest.NewRecorder())</span>
<span class="line"></span>
<span class="line">    err := goroutine.SafeGoAndWait(ctx, func() error {</span>
<span class="line">        time.Sleep(1 * time.Second)</span>
<span class="line">        return errors.New(errStr)</span>
<span class="line">    }, func() error {</span>
<span class="line">        time.Sleep(2 * time.Second)</span>
<span class="line">        return nil</span>
<span class="line">    }, func() error {</span>
<span class="line">        time.Sleep(3 * time.Second)</span>
<span class="line">        return nil</span>
<span class="line">    })</span>
<span class="line">    t.Log(&quot;safe go and wait end&quot;, time.Now().String())</span>
<span class="line"></span>
<span class="line">    if err == nil {</span>
<span class="line">        t.Error(&quot;err not be nil&quot;)</span>
<span class="line">    } else if err.Error() != errStr {</span>
<span class="line">        t.Error(&quot;err content not same&quot;)</span>
<span class="line">    }</span>
<span class="line"></span>
<span class="line">    // panic error</span>
<span class="line">    err = goroutine.SafeGoAndWait(ctx, func() error {</span>
<span class="line">        time.Sleep(1 * time.Second)</span>
<span class="line">        return errors.New(errStr)</span>
<span class="line">    }, func() error {</span>
<span class="line">        time.Sleep(2 * time.Second)</span>
<span class="line">        panic(&quot;test2&quot;)</span>
<span class="line">    }, func() error {</span>
<span class="line">        time.Sleep(3 * time.Second)</span>
<span class="line">        return nil</span>
<span class="line">    })</span>
<span class="line">    if err == nil {</span>
<span class="line">        t.Error(&quot;err not be nil&quot;)</span>
<span class="line">    } else if err.Error() != errStr {</span>
<span class="line">        t.Error(&quot;err content not same&quot;)</span>
<span class="line">    }</span>
<span class="line">}</span>
<span class="line"></span>
<span class="line"></span></code></pre><div class="line-numbers" aria-hidden="true" style="counter-reset:line-number 0;"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,13),t=[l];function r(d,c){return e(),s("div",null,t)}const o=n(i,[["render",r],["__file","util.html.vue"]]),u=JSON.parse('{"path":"/guide/util.html","title":"辅助函数","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"辅助函数","description":null},"headers":[{"level":2,"title":"goroutine 相关","slug":"goroutine-相关","link":"#goroutine-相关","children":[{"level":3,"title":"SafeGo","slug":"safego","link":"#safego","children":[]},{"level":3,"title":"SafeGoAndWait","slug":"safegoandwait","link":"#safegoandwait","children":[]}]}],"git":{"updatedTime":1718112320000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":2}]},"filePathRelative":"guide/util.md"}');export{o as comp,u as data};
