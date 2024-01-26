import { defaultTheme } from '@vuepress/theme-default'
import { defineUserConfig } from 'vuepress/cli'
import { viteBundler } from '@vuepress/bundler-vite'

export default defineUserConfig({
  lang: 'en-US',

  title: 'Gob 框架',
  description: '一个支持前后端开发的基于协议的框架',

  sidebarDepth: 2,

  theme: defaultTheme({

    // logo
    logo: 'https://vuejs.press/images/hero.png',

    // navbar: ['/', '/get-started'],

    // 添加导航栏
    navbar: [
      {text: "主页", link: "/"}, // 导航条
      {text: "使用文档", link: "/guide/"},
      {text: "服务提供者", link: "/provider/"},
      {text: "Github", link: "https://github.com/chenbihao/gob"}
    ],
  }),

  bundler: viteBundler(),
})
