# dockerfile产生镜像的方式，
# 其实就是在现有的镜像基础上，
# 加入自定义的一些操作后，
# 就成为了自定义的新镜像

# 基础镜像
FROM alpine:3.12

# RUN指令是跑在基础镜像当中的，不是宿主主机中
RUN mkdir -p "/data/app"
COPY service.log /data/app/service.log
RUN ln -sf /dev/stdout /data/app/service.log

# 规定基础镜像中的工作目录
WORKDIR "/data/app"

# COPY指令是将宿主主机中的文件拷贝到基础镜像中，
# COPY参考的相对路径起点就是Dockerfile所在的目录
COPY hello /data/app/hello

# 规定当执行 docker run <本镜像> 后，在container中会执行什么命令，
# 这里的 . 遵从 WORKDIR命令，取 "/data/app"
ENTRYPOINT ["./hello"]