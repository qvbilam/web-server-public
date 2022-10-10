## 依赖
```shell
# 框架
go get -u github.com/gin-gonic/gin

# oss
go get -u github.com/aliyun/aliyun-oss-go-sdk/oss
# 点播
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk

# 日志
go get -u go.uber.org/zap

# 配置(后者用来读取系统参数
# go get -u github.com/spf13/viper
go get -u github.com/namsral/flag

go get -u google.golang.org/grpc

# grpc 重连
go get -u github.com/grpc-ecosystem/go-grpc-middleware/retry
```

## 部署
```shell
# 构建镜像
docker build -t qvbilam/oss-web-server-alpine:[TAG] .

# 启动容器
docker container run -d -p 9501:9501 qvbilam/oss-web-server-alpine:[TAG]

# 登陆阿里云
docker login --username=13501294164 registry.cn-hangzhou.aliyuncs.com

# 镜像ID
docker tag [IMAGE ID] registry.cn-hangzhou.aliyuncs.com/qvbilam/oss-web:[TAG]
docker push registry.cn-hangzhou.aliyuncs.com/qvbilam/oss-web:[TAG]
```

## 启动
```shell
# 创建访问令牌
kubectl create secret docker-registry ali-image-key --docker-server=registry.cn-hangzhou.aliyuncs.com --docker-username=13501294164 --docker-password=kenan123456789

# 申请配置
kubectl apply -f oss.config.yaml

# 启动服务
kubectl apply -f oss.deployment.yaml
kubectl apply -f oss.server.yaml

# 查看pods
kubectl get pods

# 查看服务
kubectl get svc

# 开放端口
kubectl port-forward service/oss-web-server 9501:9501 -n default
```