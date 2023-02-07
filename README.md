![魔豆文库](web/static/static/images/logo.png)

# MOREDOC - 魔豆文库

`moredoc` - 魔豆文库，由 深圳市摩枫网络科技(**M**orefun **N**etwork **T**echnology Co., **Ltd**，https://mnt.ltd ) 使用 Golang 开发的类似[百度文库](https://wenku.baidu.com/)、[新浪爱问文库](http://ishare.iask.sina.com.cn/)的开源文库系统，支持 `TXT`、`PDF`、`EPUB`、`MOBI`、`Office` 等格式文档的在线预览与管理，为 `dochub文库`( [github](https://github.com/truthhun/dochub), [gitee](https://gitee.com/truthhun/DocHub) )的重构版本。

## 文库介绍

## 技术栈

- Golang ：gin + gRPC + GORM
- Vue.js : nuxt2 + element-ui
- Database : MySQL 5.7

## 开源地址

- Github - https://github.com/mnt-ltd/moredoc
- Gitee - https://gitee.com/mnt-ltd/moredoc
- MNT.Ltd - https://git.mnt.ltd/mnt/moredoc

## 使用手册

关于魔豆文库安装部署以及使用和二次开发等更详细的教程，详见书栈网[《魔豆文库使用手册》](https://www.bookstack.cn/books/moredoc)

## 演示站点

**网址**

- https://moredoc.mnt.ltd

**账号密码**

```
账号：admin
密码：mnt.ltd
```

## 微信交流群

魔豆文库微信交流群，请添加`进击的皇虫`的微信，备注`魔豆文库加群`，以便进群。

**微信二维码**

![魔豆文库微信交流群](docs/wx-qrcode.jpeg)

## 二次开发

除了文件上传相关的接口，其他接口统一使用 proto 进行定义。

### 环境要求

- Golang 1.18+
- Node.js 14.16.0 (可用 nvm 管理)
- MySQL 5.7+

**请自行配置相应环境。如在此过程中遇到错误，请根据错误提示自行通过 Google 或者百度解决。**

### 目录结构

> 部分目录，在程序运行时自动生成，不需要手动创建

```bash
.
├── LICENSE                 # 开源协议
├── Makefile                # 编译脚本
├── README.md               # 项目说明
├── api                     # proto api， API协议定义
├── app.example.toml        # 配置文件示例，需要复制为 app.toml
├── biz                     # 业务逻辑层，主要处理业务逻辑，实现api接口
├── cmd                     # 命令行工具
├── cache                   # 缓存相关
├── conf                    # 配置定义
├── dict                    # 结巴分词字典，用于给文档自动进行分词
├── dist                    # 前端打包后的文件
├── docs                    # API文档等
├── documents               # 用户上传的文档存储目录
├── go.mod                  # go依赖管理
├── go.sum                  # go依赖管理
├── main.go                 # 项目入口
├── middleware              # 中间件
├── model                   # 数据库模型，使用gorm对数据库进行操作
├── release                 # 版本发布生成的版本会放到这里
├── service                 # 服务层，衔接cmd与biz
├── sitemap                 # 站点地图
├── third_party             # 第三方依赖，主要是proto文件
├── uploads                 # 文档文件之外的其他文件存储目录
├── util                    # 工具函数
└── web                     # 前端Web
```

### app.toml

```
# 程序运行级别：debug、info、warn、error
level="debug"

# 日志编码方式，支持：json、console
logEncoding="console"

# 后端监听端口
port="8880"

# 数据库配置
[database]
    driver="mysql"
    dsn="root:root@tcp(localhost:3306)/moredoc?charset=utf8mb4&loc=Local&parseTime=true"
    showSQL=true
    maxOpen=10
    maxIdle=10

# jwt 配置
[jwt]
    secret="moredoc"
    expireDays=365
```

### 初始化

**后端初始化**

```
# 安装go依赖
go mod tidy

# 初始化工程依赖
make init

# 编译proto api
make api

# 修改 app.toml 文件配置
cp app.example.toml app.toml

# 编译后端
go build -o moredoc main.go

# 初始化数据库结构
./moredoc syncdb

# 运行后端(可用其他热编译工具)，监听8880端口
go run main.go serve
```

**前端初始化**

```bash
# 切换到web目录
cd web

# 安装依赖
npm install

# 运行前端，监听3000端口，浏览器访问 http://localhost:3000
npm run dev
```

### 管理员初始账号密码

```
admin
mnt.ltd
```

### 发布版本

以下为示例

```
# 打标签
git tag -a v1.0.0 -m "release v1.0.0"

# 推送标签
git push origin v1.0.0

# 编译前端
cd web && npm run generate

# 编译后端，编译好了的版本会放到release目录下
# 编译linux版本（Windows版本用 make buildwin）
make buildlinux
```

## License

开源版本基于 [Apache License 2.0](./LICENSE) 协议发布。

## 鸣谢

感谢各开源项目为魔豆文库的开发奠定了基础。相关开源项目，后端依赖，详见 [go.mod](./go.mod)；前端依赖，详见 [web/package.json](./web/package.json)。
