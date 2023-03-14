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
sudo docker run -it -p 18880:8880 -v /home/bluezealot/work/morebook/cache:/home/moredoc/workspace/cache -v /home/bluezealot/work/morebook/document:/home/moredoc/workspace/documents -v /home/bluezealot/work/morebook/uploads:/home/moredoc/workspace/uploads -e MYSQL_CONNECTION="root:password@tcp(10.50.30.59:32306)/moredoc_test?charset=utf8mb4&loc=Local&parseTime=true" bluezealot/moredoc:Linux_ce_v1.0.0
```