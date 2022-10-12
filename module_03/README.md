# build
>make build

# build docker image
>make release

# 将镜像推送至 docker 官方镜像仓库
>make push

# 通过 docker 命令本地启动 httpserver
>docker run -d --restart=always -v /tmp:/tmp -p 8080:8080 --name="httpserver" hhsd/httpserver:v1.0

# 通过 nsenter 进入容器查看 IP 配置
>PID=$(docker inspect --format {{.State.Pid}} httpserver)

>sudo nsenter -t $PID -n ip a
