# moredoc - 魔豆文库

moredoc - 魔豆文库，由 深圳市摩枫网络科技(https://mnt.ltd) 基于golang开发的类似百度文库的开源文库系统，支持TXT、PDF、EPUB、office等格式文档的在线预览与管理，为 `dochub文库`( [github](https://github.com/truthhun/dochub), [gitee](https://gitee.com/truthhun/DocHub) )的重构版本。

## 特性

- 支持多种格式文档的在线预览，包括：TXT、PDF、EPUB、office等格式文档。
- 前后端分离，前端基于vue.js，后端基于golang，支持多平台部署
- 基于apache2.0开源协议，源码开放，可自由修改、二次开发

## 技术栈

- Golang   ：gin + gRPC
- Vue.js   : nuxt2 + element-ui
- Database : MySQL 5.7

## 开源地址

- Github - https://github.com/mnt-ltd/moredoc
- Gitee - https://gitee.com/mnt-ltd/moredoc
- MNT.Ltd - https://git.mnt.ltd/mnt/moredoc

## 初始化

```
# 安装go依赖
go mod tidy

# 初始化工程依赖
make init

# 编译proto api
make api
```

## 开发说明

- 除了文件上传相关的接口，其他接口一律使用proto进行定义。

## 管理员初始账号密码

```
账号：admin
密码：mnt.ltd
```