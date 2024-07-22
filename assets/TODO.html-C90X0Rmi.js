import{_ as l,c as i,o as e,b as a}from"./app-CiuEt72j.js";const r={},n=a('<h1 id="待办" tabindex="-1"><a class="header-anchor" href="#待办"><span>待办</span></a></h1><h2 id="框架支持功能" tabindex="-1"><a class="header-anchor" href="#框架支持功能"><span>框架支持功能</span></a></h2><h3 id="框架模块优化" tabindex="-1"><a class="header-anchor" href="#框架模块优化"><span>框架模块优化</span></a></h3><ul><li><p>[ ] 日志</p><ul><li>[ ] 同时支持多个日志输出，并且接管gin的日志输出</li><li>[ ] 统一优化日志打印格式</li><li>[ ] 可选固定 json 字段顺序打印配置（强迫症可选）</li></ul></li><li><p>[ ] model 代码生成功能优化</p><ul><li>[ ] 生成更符合业务调用场景</li><li>[ ] 其他数据源如sqlite的字段优化</li></ul></li><li><p>[ ] 调试模式</p><ul><li>[ ] 文件监控比服务启动还提前，编译完成前修改文件可能会导致空指针</li><li>[ ] swagger 热更新？热读取版本配置？</li></ul></li><li><p>[ ] 部署模式</p><ul><li>[ ] 当断开 ssh 时服务会停止（特定问题，Manjaro 系统特有）</li><li>[ ] 发布自动化完善（部署备份、部署回滚）</li></ul></li><li><p>[ ] 将过时或停止维护的三方库换掉</p><ul><li>[ ] survey 换成 bubbletea <ul><li>[ ] 命令行支持静默运行参数</li></ul></li><li>[ ] github 的调用限制重构复用</li><li>[ ] gspt 构建出错 ， 需要交叉编译</li></ul></li><li><p>[ ] 前端文件夹配置可选？</p><ul><li>[ ] frontendFolder 配置可选</li></ul></li><li><p>[ ] 数据库重连重试机制</p></li><li><p>[ ] 其他优化</p><ul><li>[ ] 把其他命令适配到纯工具模式（<code>go install</code>）</li><li>[ ] win 下不支持 Daemon，兼容成后台运行（appDaemon）</li><li>[ ] cache 服务当配置了 redis，并且有 redis 相关配置时，优先读取 redis 配置</li></ul></li></ul><h3 id="框架模块新增" tabindex="-1"><a class="header-anchor" href="#框架模块新增"><span>框架模块新增</span></a></h3><ul><li><p>[ ] 远程配置中心</p></li><li><p>[ ] 引入数据库迁移，方便后续蓝图，选型：</p><ul><li>简单 gorm 迁移增强 <ul><li><a href="https://github.com/go-gormigrate/gormigrate" target="_blank" rel="noopener noreferrer">gormigrate</a></li></ul></li><li>驱动多功能强大，无建表 <ul><li><a href="https://github.com/golang-migrate/migrate" target="_blank" rel="noopener noreferrer">migrate</a></li></ul></li><li>驱动少功能强大，有建表，有导出架构 <ul><li><a href="https://github.com/amacneil/dbmate" target="_blank" rel="noopener noreferrer">dbmate</a></li></ul></li></ul></li><li><p>[ ] 目前考虑 dbmate</p><ul><li>需要重新思考model生成相关逻辑</li><li>先手写model后表？</li><li>先手写sql后model和api？</li><li>都兼容？</li></ul></li></ul><h3 id="蓝图模块功能" tabindex="-1"><a class="header-anchor" href="#蓝图模块功能"><span>蓝图模块功能</span></a></h3><ul><li><p>[ ] 初始化蓝图流程</p><ul><li>[ ] 蓝图定义，包括依赖关系、版本等</li><li>[ ] 定义拉取蓝图模块流程</li><li>[ ] 拉取后执行表迁移工作</li></ul></li><li><p>[ ] 后台管理基础</p><ul><li>[ ] 低代码快速搭建？</li></ul></li><li><p>[ ] 用户注册登录</p><ul><li>[ ] RBAC 权限</li><li>[ ] 多租户模块</li></ul></li><li><p>[ ] 博客</p></li><li><p>[ ] ...</p></li><li><p>[ ] ...</p></li></ul><h2 id="其他功能优化" tabindex="-1"><a class="header-anchor" href="#其他功能优化"><span>其他功能优化</span></a></h2><ul><li><p>[ ] 部分 linux 内容未测试</p><ul><li>[ ] 条件编译</li><li>[ ] 守护进程模式 <code>app start --daemon=true</code></li><li>[ ] gspt 库（<code>CGO_ENABLED=1</code>）</li></ul></li><li><p>[ ] 业务单测的构建</p></li></ul><h2 id="已完成归档" tabindex="-1"><a class="header-anchor" href="#已完成归档"><span>已完成归档</span></a></h2><h3 id="梳理相关" tabindex="-1"><a class="header-anchor" href="#梳理相关"><span>梳理相关</span></a></h3><ul><li><p>[x] 梳理使用框架</p><ul><li>[x] 梳理源码引入 <ul><li>cobra</li><li>gin v1.9.1 + middleware</li></ul></li><li>[x] 梳理三方库引入 <ul><li>fsnotify、go-daemon、goconvey、swaggo、cast</li><li>survey/v2、go-git/v5、go-github/v62、go-redis/v9、cron/v3、gorm + gen</li><li>gotree、uuid、xid、ratelimit、file-rotatelogs、mapstructure</li><li>jennifer/jen、jianfengye/collection、kr/pretty</li></ul></li><li>[x] 梳理三方框架使用 <ul><li>vue、vuepress</li></ul></li></ul></li><li><p>[x] 梳理新版本 go 废弃 API，换成新的</p><ul><li><code>io/ioutil</code> -&gt; <code>os</code>、<code>io</code></li><li><code>strings.Title</code> -&gt; <code>cases.Title</code></li><li><code>math/rand</code> -&gt; <code>rand.Rand</code></li></ul></li></ul><h3 id="统一代码" tabindex="-1"><a class="header-anchor" href="#统一代码"><span>统一代码</span></a></h3><ul><li>[x] 统一 provider 注册方法 （<code>func (provider *GormProvider) Register</code> 里的调用 new 命名）</li><li>[x] 补充 command 、contract 文件开头说明文档，方便查看（甚至改成支持 doc） <ul><li>[x] command：包括命令说明、可选配置项</li><li>[x] contract：包括对应命令、配置项说明</li></ul></li></ul><h3 id="框架模块优化-1" tabindex="-1"><a class="header-anchor" href="#框架模块优化-1"><span>框架模块优化</span></a></h3><ul><li>[x] 脚手架优化 <ul><li>[x] <code>go install</code> 使用的优化（新增纯工具模式）</li></ul></li></ul>',17),t=[n];function o(d,s){return e(),i("div",null,t)}const u=l(r,[["render",o],["__file","TODO.html.vue"]]),p=JSON.parse('{"path":"/guide/TODO.html","title":"待办","lang":"zh-CN","frontmatter":{"lang":"zh-CN","title":"待办","description":null},"headers":[{"level":2,"title":"框架支持功能","slug":"框架支持功能","link":"#框架支持功能","children":[{"level":3,"title":"框架模块优化","slug":"框架模块优化","link":"#框架模块优化","children":[]},{"level":3,"title":"框架模块新增","slug":"框架模块新增","link":"#框架模块新增","children":[]},{"level":3,"title":"蓝图模块功能","slug":"蓝图模块功能","link":"#蓝图模块功能","children":[]}]},{"level":2,"title":"其他功能优化","slug":"其他功能优化","link":"#其他功能优化","children":[]},{"level":2,"title":"已完成归档","slug":"已完成归档","link":"#已完成归档","children":[{"level":3,"title":"梳理相关","slug":"梳理相关","link":"#梳理相关","children":[]},{"level":3,"title":"统一代码","slug":"统一代码","link":"#统一代码","children":[]},{"level":3,"title":"框架模块优化","slug":"框架模块优化-1","link":"#框架模块优化-1","children":[]}]}],"git":{"updatedTime":1721578757000,"contributors":[{"name":"陈壁浩","email":"chenbihao@qljy.com","commits":8}]},"filePathRelative":"guide/TODO.md"}');export{u as comp,p as data};
