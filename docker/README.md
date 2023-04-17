# 关于Docker

## 端口
docker 的-p参数，目前向外暴露了8880端口。

## 文件
docker 的-v参数，目前向外暴露了三个路径
- /home/moredoc/workspace/cache
- /home/moredoc/workspace/documents
- /home/moredoc/workspace/uploads

## 数据库链接地址
提供环境变量 MYSQL_CONNECTION 来设置

## docker run 命令启动例子

```
sudo docker run -it -p 18880:8880 -v /home/bluezealot/work/morebook/cache:/home/moredoc/workspace/cache -v /home/bluezealot/work/morebook/document:/home/moredoc/workspace/documents -v /home/bluezealot/work/morebook/uploads:/home/moredoc/workspace/uploads -e MYSQL_CONNECTION="root:password@tcp(10.50.30.59:32306)/moredoc_test?charset=utf8mb4&loc=Local&parseTime=true" bluezealot/moredoc:Linux_ce_v1.3.0
```

## 关于升级

我测试过从V1.0.0经过所有中间版本升级到V1.3.0,不知道是否可以跳级升级，但是docker的启动脚本里面包含了moredoc要求的升级script，理论上应该没有问题。另外moredoc支持的云存储方式在这个docker镜像上没有反映，
这个image只有mount到本地文件夹的方式,我本地使用的时候是直接使用k8s mount NFS的方式，也是没有问题的。
如果使用中有问题，可以继续探讨。