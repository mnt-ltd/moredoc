# Dockerfile 部署魔豆文库

当前[`Dockerfile`](Dockerfile)，在[@bluezealot](https://github.com/bluezealot)之前给魔豆文库提交的Dockerfile PR的基础上修改调整而来。且同样适用于ARM架构的服务器。

> 您决定使用docker来部署魔豆文库，则默认您对docker有基本的了解。


**当前教程，所有操作指令，均为在`~/moredoc`目录下执行，特此说明。**

## 使用 Dockerfile 部署魔豆文库步骤

### 1. 安装docker

略


### 2. 下载魔豆文库

当前魔豆文库`Dockerfile`，是基于`ubuntu:22.04`构建镜像。如果您的docker所在服务器CPU架构是arm，则下载`xxx_linux_arm64.tar.gz`，否则下载`xxx_linux_amd64.tar.gz`。

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

### 3. 构建文库镜像

下载当前[`Dockerfile`](Dockerfile)文件，将文件存放到`~/moredoc`下。或复制当前Dockerfile文件内容，粘贴到`~/moredoc/Dockerfile`文件中。

```
cd ~/moredoc
[sudo] docker build -t moredoc:latest .
```

### 4. 安装MySQL

这里用docker安装，将数据库挂载到`~/moredoc/database`下。

```
mkdir ~/moredoc/database
cd ~/moredoc
[sudo] docker run --name moredoc-mysql -d -p 127.0.0.1:33060:3306 -e MYSQL_ROOT_PASSWORD=moredoc --restart=always -v $(pwd)/database:/var/lib/mysql mysql:8.0.36
```

数据库容器名为`moredoc-mysql`，密码为`moredoc`，对宿主机暴露`33060`端口。需要注意的是，这里监听的IP地址是 `127.0.0.1`，不建议监听`0.0.0.0`。

### 5. 修改程序配置

查看MySQL容器IP地址：
```
[sudo] docker inspect moredoc-mysql
```

显示如下：
```
...
"bridge": {
    "IPAMConfig": null,
    "Links": null,
    "Aliases": null,
    "MacAddress": "02:42:ac:11:00:02",
    "NetworkID": "e5f2f39ebf864af4c81195c09e664fd4f072489f299812e6c27d95cf632dc120",
    "EndpointID": "02cb3b8b1dac0484bcb72809d67247b65c28139af024854eb71cb67380685aac",
    "Gateway": "172.17.0.1",
    "IPAddress": "172.17.0.2",
    "IPPrefixLen": 16,
    "IPv6Gateway": "",
    "GlobalIPv6Address": "",
    "GlobalIPv6PrefixLen": 0,
    "DriverOpts": null
}
...
```

也就是地址为：`172.17.0.2`。

在`~/moredoc/server`中，重命名`app.example.toml`为`app.toml`，并修改`app.toml`中的数据库用户名和密码。如下：
```
...
# 数据库配置
[database]
    driver = "mysql"
    dsn = "root:moredoc@tcp(172.17.0.2:3306)/moredoc?charset=utf8mb4&loc=Local&parseTime=true"
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

### 6. 启动文库容器
```
cd ~/moredoc
[sudo] docker run -it -d --name moredoc-server  --restart=always -p 8880:8880 -v $(pwd)/server:/home/moredoc moredoc:latest
```

当前容器名为`moredoc-server`，监听`8880`端口，挂载目录为`~/moredoc/server`。

这时就可以通过 `IP:8880` 来访问文库服务了。

文库初始账号密码为：
```
账号： admin
密码： mnt.ltd
```

## FAQ

### 后续文库程序怎么升级更新？

如果后续文库系统发布了新版本，则直接将新版本文库系统解压覆盖到`~/moredoc/server`下，然后重启容器`moredoc-server`即可。

### 开机启动后，文库启动失败，怎么办？

在上述操作指令中，魔豆文库容器`moredoc-server`和MySQL容器`moredoc-mysql`，都是设置了`--restart=always`的，但是`moredoc-server`依赖`moredoc-mysql`，如果此时MySQL容器启动慢了，会导致文库服务访问数据库失败，从而启动不成功。此时，您可以手动启动，也可以通过`docker-compose`来解决这个问题，详见[docker-compose-README.md](docker-compose-README.md)。

当然，也要注意docker是否跟着开机启动了。

### 为什么不把容器更新到docker hub ?

个人在本地构建的容器镜像，因为安装了`libreoffice`、`calibre`等，镜像大小高达`3.37GB`，因此没必要更新上去。且`Dockerfile`中已替换成阿里云的源，在本地构建镜像，也不比你直接拉取镜像慢。