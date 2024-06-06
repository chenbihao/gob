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
            {text: "使用文档", link: "/guide/introduce"},
            {text: "服务提供者", link: "/provider/"},
            {text: "Github", link: "https://github.com/chenbihao/gob"}
        ],
        // 为以下路由添加侧边栏
        sidebar: {
            "/guide/": [
                {
                    title: "指南",
                    collapsable: false,
                    children: [
                        "introduce",
                        "install",
                        "version",
                        "build",
                        "structure",
                        "app",
                        "env",
                        "dev",
                        "command",
                        "cron",
                        "middleware",
                        "swagger",
                        "provider",
                        "model",
                        "deploy",
                        "util",
                        "grpc",
                        "todo",
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
