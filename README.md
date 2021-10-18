# 作业二
  1.构建本地镜像 
  
    $ docker build -t zhangxin666/httpserver .
    
  2.推送镜像到 DockerHub
  
    $ docker push zhangxin666/httpserver:v0.0.1
    
  3.Docker启动本地httpserver镜像
  
    $ docker run -d -p 8080:8080 zhangxin666/httpserver
    
  4.nsenter 进入容器内查看 IP 配置
    
    $ nsenter -t <pid> -n ip addr
