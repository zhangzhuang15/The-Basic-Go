本项目源自Golang语言开发栈frank  
部分内容经过适度修改  

## 你该如何做
1. 执行`GOOS=linux GOARCH=amd64 build -o hello` 生成可执行文件`hello`;

2. 执行`docker build -t hello:v1.0.0 .` 生成镜像`hello:v1.0.0`;  

3. 执行`docker images`, 检查是否生成镜像`hello:v1.0.0`;

4. 如果生成了，执行 `docker run -d -p 8090:8089 hello:v1.0.0`;

5. 执行`docker ps -a`, 应该可以看到 `hello`容器是运行状态的；

6. 执行 `curl http://localhost:8090/hello`, 收到返回值`hello world`;

*注意：*
> 运行第一条指令时，一定要指明 GOOS 和 GOARCH，因为基础镜像是
> amd64位的linux操作系统，生成的可执行文件hello也必须是这个平台
> 上的。因为有的人会在arm芯片上的电脑尝试这个项目，所以不得不考虑
> 交叉编译情形。


## 项目中的文件是干啥的
* main.go 文件写了一个http服务程序，访问`http://localhost:8089/hello` 时，直接返回
一个`hello world`;  

* Dockerfile说明如何产生一个镜像，里边有注释，如果不明白，
可以参考`https://docs.docker.com/engine/reference/commandline/build/#build-with--`,
相信里边会有你要的解释;   

* `service.log`, 用于将容器的标准输出写入到自身中，在Dockerfile中也有提到；

* 至于为啥访问8090端口会收到返回，是因为通过`-p` 选项，将宿主主机端口8090映射
为容器中的端口8089，访问宿主主机的端口8090，就会转为访问容器的8089端口，自然
就会收到hello world。