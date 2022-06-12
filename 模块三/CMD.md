- 构建本地镜像
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化

```bash
🐧 ls
dockerfile  go.mod  go.sum  main.go

# 基于 dockerfile 构建镜像
docker build -t httpserver:1.0 .
```

- 将镜像推送至 docker 官方镜像仓库

```bash
# 登录 dockerhub.com，创建名为 enjoylinux 的公共镜像仓库

# 修改镜像标签
docker tag httpserver:1.0 enjoylinux/httpserver:1.0

# 登录 docker hub
docker login

# 推送镜像
docker push enjoylinux/httpserver:1.0
```

- 通过 docker 命令本地启动 httpserver

````bash
# 创建并启动容器
docker run -itd -p 80:80 httpserver:1.0
````

- 通过 nsenter 进入容器查看 IP 配置

```bash
# 查看容器 id 
🐧 docker ps | grep httpserver | awk '{print $1}'
afae47e702d3

# 查看容器在宿主机上对应的 pid
🐧 docker inspect afae47e702d3 | grep -i pid
            "Pid": 123269,
            "PidMode": "",
            "PidsLimit": null,

# 通过 nsenter 查看容器 ip 配置
🐧 nsenter -t 123269 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
9055: eth0@if9056: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:03 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.3/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

