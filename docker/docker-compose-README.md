# Docker Compose 部署魔豆文库

> 您决定使用docker来部署魔豆文库，则默认您对docker有基本的了解。


**当前教程，所有操作指令，均为在`~/moredoc`目录下执行，特此说明。**

## 适用 docker compose 部署魔豆文库步骤

### 1. 安装docker及docker compose

略

### 2. 下载魔豆文库

如果您的docker所在服务器CPU架构是arm，则下载`xxx_linux_arm64.tar.gz`，否则下载`xxx_linux_amd64.tar.gz`。

**下载地址：**
- Gitee: https://gitee.com/mnt-ltd/moredoc/releases
- Github: https://github.com/mnt-ltd/moredoc/releases

下载之后，将程序解压到相应目录，这里解压目录为`~/moredoc/server`，如下：

```
-rw-r--r-- app.example.toml
drwxr-xr-x dictionary
drwxr-xr-x dist
-rwxr-xr-x moredoc
```

需要注意的是，如果`moredoc`没有可执行权限，则需要设置下权限：
```
[sudo] chmod +x moredoc
```

### 3. 修改程序配置


在`~/moredoc/server`中，重命名`app.example.toml`为`app.toml`，并修改`app.toml`中的数据库用户名和密码。如下(数据库配置的`dsn`，照搬即可)：
```
...
# 数据库配置
[database]
    driver = "mysql"
    dsn = "root:moredoc@tcp(moredoc-mysql:3306)/moredoc?charset=utf8mb4&loc=Local&parseTime=true"
    ## 是否显示 SQL 语句，生产环境下请设置为false
    showSQL = true
    maxOpen = 10
    maxIdle = 10
...
```

然后修改`[jwt]`的配置，把`secret`修改为其他值：
```
...
## JSON Web Token 配置
[jwt]
    ## 生成 token 的密钥，请务必修改！！！
    secret = "hello world"
    ## token 过期时间
    expireDays = 365
...
```

### 4. 启动文库

下载当前目录下的[`Dockerfile`](Dockerfile)和[`docker-compose.yml`](docker-compose.yml)两个文件，放到`~/moredoc`目录下。

然后执行下述指令：
```
cd ~/moredoc
[sudo] docker-compose -f docker-compose.yml up
```

当前容器名为`moredoc-server`，监听`8880`端口，挂载目录为`~/moredoc/server`。

这时就可以通过 `IP:8880` 来访问文库服务了。

文库初始账号密码为：
```
账号： admin
密码： mnt.ltd
```


