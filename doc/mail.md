# 构建service-mail 镜像

1. 下载ubuntu:18.04镜像
```bash
docker pull ubuntu:18.04
```

2. 进入ubuntu镜像并且安装nodejs跟npm
```
docker run -it ubuntu:18.04 bash
#进入docker 镜像交互
apt-get update
apt-get install
apt-get install curl
#安装nodejs 10
curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash -
apt-get install -y nodejs
#安装npm
apt-get install npm
```

3. nodejs 安装完成后将service-mail 文件移进该镜像

- 1.打开新的终端
- 2.`docker container ls` 查看container ID , 找到最近一个, 我的是`6c5610365160`
- 3.`docker cp ./service-mail/ 6c5610365160:/service-mail` 将nodejs文件考入容器.
- 4.在交互式终端中就可以看见导入的文件了

4. 打包镜像

- 1.退出交互式镜像 `exit`
- 2.将容器打包成镜像 `docker commit 6c5610365160 service-mail:1.0.0`
- 3.查看镜像 `docker image ls`
- 4.开启服务 `docker run -it service-mail:1.0.0 sh -c "cd service-mail && npm install && nodejs server.js"`




## 镜像成功之后启动 mail 服务

1. 进入bash
`docker run --ip=192.169.0.11 -it --network=service-discovery --name=mail service-mail:1.0.0  bash`

2. 启动服务
`node server.js`