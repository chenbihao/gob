import {defaultTheme} from '@vuepress/theme-default'
import {defineUserConfig} from 'vuepress/cli'
import {viteBundler} from '@vuepress/bundler-vite'
import {searchPlugin} from '@vuepress/plugin-search'
import {mdEnhancePlugin} from "vuepress-plugin-md-enhance"

export default defineUserConfig({
    lang: 'en-US',
    title: 'Gob 框架',
    description: '一个支持前后端开发的基于协议的框架',

    sidebarDepth: 2,
    base: '/gob/',

    head: [["link", {rel: "icon", href: "/images/logo.png"}]],
    theme: defaultTheme({

        // logo
        logo: '/images/logo.png',

        // 添加导航栏
        navbar: [
            {text: "主页", link: "/"}, // 导航条
            {text: "使用文档", link: "/guide/"},
            {text: "服务提供者", link: "/provider/"},
            {text: "提供命令", link: "/command/"},
            {text: "Github", link: "https://github.com/chenbihao/gob"}
        ],
        // 为以下路由添加侧边栏
        sidebar: {
            "/guide/": [
                {
                    title: "指南",
                    collapsable: false,
                    children: [
                        "",             // 介绍
                        "introduce",    // 快速上手
                        "install",      // 安装
                        "structure",    // 目录结构
                        "app",          // 运行
                        "version",      // 版本
                        "build",        // 编译
                        "env",          // 环境变量
                        "dev",          // 调试模式
                        "command",      // 命令
                        "cron",         // 定时任务
                        "middleware",   // 中间件
                        "swagger",      // swagger
                        "provider",     // 服务提供者
                        "model",        // 模型
                        "deploy",       // 自动部署
                        "util",         // 辅助函数
                        "grpc",         // grpc 支持
                        "TODO",         // 代办
                    ],
                },
            ],
            "/command/": [
                {
                    title: "提供命令",
                    collapsable: false,
                    children: [
                        "app",
                        "config",
                        "new",
                        "build",
                        "env",
                        "dev",
                        "cmd",
                        "cmd_go",
                        "cmd_npm",
                        "cron",
                        "middleware",
                        "swagger",
                        "provider",
                        "deploy",
                        "version",
                        "model",
                    ],
                },
            ],
            "/provider/": [
                {
                    title: "服务提供者",
                    collapsable: false,
                    children: [
                        "app",
                        "cache",
                        "config",
                        "distributed",
                        "env",
                        "id",
                        "kernel",
                        "log",
                        "orm",
                        "redis",
                        "sls",
                        "ssh",
                        "trace",
                    ],
                },
            ],
        },
    }),
    plugins: [
        searchPlugin({
            // 配置项
        }),
        mdEnhancePlugin({
            codetabs: true,
            mermaid: true,
        }),
    ],

    bundler: viteBundler(),
})
