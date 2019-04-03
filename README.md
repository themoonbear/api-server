# API 服务器

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
