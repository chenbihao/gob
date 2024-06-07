package main

var providerMD = `---
lang: zh-CN
title: {{.key}}
description:
---
# {{.key}}

{{.remark}}

## 提供方法：
{{.code}}
`

var commandMD = `---
lang: zh-CN
title: {{.key}}
description:
---
# {{.key}}

{{.remark}}

## 使用方法：
{{.code}}
`
