## 启动mysql镜像

`docker pull papandadj/mysql:1.0.0`

1. mysql镜像已经配置好了, 每次进去只需要执行`service mysql restart`可以其它同网段访问


## 在局域网启动 registry 服务

依赖golang, 需要dockerfile编译

1. 进入bash
`docker run --ip=192.169.0.10 -it --network=service-discovery --name=registry service-registry:1.0.0  bash`

2. 执行

`./app



## 在局域网启动 balance 服务

依赖golang, 需要dockerfile编译

1. 进入bash
`docker run -p 5050:5050 --ip=192.169.0.12 -it --network=service-discovery --name=balance service-balance:1.0.0  bash`

2. 执行

`./app`