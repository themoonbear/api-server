# API 服务器

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cfd41b89e17142369aa1e9df1bf963bb)](https://app.codacy.com/app/contact_114/api-server?utm_source=github.com&utm_medium=referral&utm_content=themoonbear/api-server&utm_campaign=Badge_Grade_Settings)
[![CircleCI](https://circleci.com/gh/themoonbear/api-server/tree/master.svg?style=svg)](https://circleci.com/gh/themoonbear/api-server/tree/master)

提供 API 服务

## 接口

+ 代理接口，用于访问屏蔽网址

```shell
curl api.moonbear.cn/proxy/?address=base64(address)
```

+ 弹幕接口，用于处理直播聊天

```shell
curl api.moonbear.cn/danmu/base64(address)
```
