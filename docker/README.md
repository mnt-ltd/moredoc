# Docker部署魔豆文库

当前`dockerfile`，在[@bluezealot](https://github.com/bluezealot)之前给魔豆文库提交的dockerfile PR的基础上修改调整而来。

> 您决定使用docker来部署魔豆文库，则默认您对docker有基本的了解。


**当前教程，所有操作指令，均为在`~/moredoc`目录下执行，特此说明。**

## 安装docker

略

## 构建镜像

下载当前`dockerfile`文件，将文件存放到`~/moredoc`下。或复制当前dockerfile文件内容，粘贴到`~/moredoc/dockerfile`文件中。

```
cd ~/moredoc
[sudo] docker build -t moredoc:latest .
```

## 下载魔豆文库

当前魔豆文库`dockerfile`，是基于`ubuntu:22.04`构建镜像。如果您的docker所在服务器CPU架构是arm，则下载`xxx_linux_arm64.tar.gz`，否则下载`xxx_linux_amd64.tar.gz`。

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
chmod +x moredoc
```


## 安装MySQL

```
cd ~/moredoc
[sudo] docker run --name moredoc-mysql -d -p 127.0.0.1:33060:3306 -e MYSQL_ROOT_PASSWORD=moredoc --restart=always --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci -v $(pwd)/database:/var/lib/mysql
```

## 修改程序配置

查看MySQL容器IP地址：
```
[sudo] docker inspect moredoc-mysql
```

这里，我们修改



