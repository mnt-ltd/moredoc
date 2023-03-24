# Version: 1.0.0

# 此Dockerfile为半自动的，制作的镜像包含文库的运行环境、代码和前端开发环境，go代码需要容器启动后自行编译并配置mysql

# 制作镜像：docker build --no-cache -t ubuntu-moredoc:v1.0 . (不包含mysql)
# docker启动：docker run -d --name=moredoc -p8080:3000 ubuntu-moredoc:v1.0
# 进入容器app目录编译moredoc，初始化数据库，启动服务：
#   dcocker exec -it moredoc /bin/bash         #进入容器
#       cd /app/moredoc                        #进入程序目录
#       go build  -ldflags="-w -s" -o moredoc  #编译项目
#       cp app.example.toml app.toml           #初始化配置文件并用vim编辑数据库信息
#       ./moredoc syncdb                       #初始化数据库
#       nohup ./moredoc serve &                #启动服务后exit退出即可，网页访问：http://ip:8080

FROM xiaoyantian03/ubuntu-moredoc-env:v1.0

MAINTAINER xiaoyantian "xiaoyantian03@163.com"

WORKDIR /app

RUN git clone https://gitee.com/mnt-ltd/moredoc.git

#RUN go build  -ldflags="-w -s" -o moredoc && ./moredoc syncdb #此步骤进入容器执行


#编译和启动前端项目
WORKDIR moredoc/web
RUN npm install

# 对外端口
EXPOSE 3000

# 运行前端项目
ENTRYPOINT ["npm", "run"]
CMD ["dev"]
