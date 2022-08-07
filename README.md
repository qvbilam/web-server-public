## 依赖
```shell
# 框架
go get -u github.com/gin-gonic/gin

# oss
go get -u github.com/aliyun/aliyun-oss-go-sdk/oss

# 日志
go get -u go.uber.org/zap

# 配置
go get -u github.com/spf13/viper
```

## 部署
```shell
# 构建镜像
docker build -t qvbilam/oss-web-server-alpine:0.0.1 .

# 启动容器
docker container run -d -p 9501:9501 qvbilam/oss-web-server-alpine:0.0.1

# 上传镜像 -p 容器ID
docker commit -a "qvbilam" -m "go oss web server 0.0.1" -p 973b6fb8001694 qvbilam/oss-web-server-alpine:0.0.1

docker tag qvbilam/oss-web-server-alpine:0.0.1 docker.qvbilam.xin:8081/qvbilam/oss-web-server-alpine:0.0.1

# 推送镜像
docker push docker.qvbilam.xin:8081/qvbilam/oss-web-server-alpine:0.0.1
```