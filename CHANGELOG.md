# 0.1.0

## 依赖更新

+ `tensorflow` ?->2.6.2
+ `tensorflow_serving` ?->2.6.2
+ `grpc` 1.36.0->1.44.0,同步改用`grpc.WithTransportCredentials(insecure.NewCredentials())`替换了`grpc.WithInsecure()`

更新了tensorflow和tensorflow_serving的版本,同时更新了grpc和protobuf的版本

## 新增工具

新增gen_proto.py用于提取tf和tf-serv中的pb并生成对应的go模块

## 新增接口

1. 增加了对`SessionServiceClient`的支持

## 接口变动

1. 将`NewCtx`改到`SDK`上,并且添加了可配置参数(具体见`ctxopt.go`文件)
2. 增加了`Get{Session|Model|Prediction}ServiceConn`方法,用于获取或创建对应的连接,现在的sdk会保留连接

# 0.0.4

## 依赖更新

+ github.com/Golang-Tools/loggerhelper@v0.0.4
+ github.com/golang/protobuf@v1.5.2
+ github.com/json-iterator/go@v1.1.12
+ google.golang.org/grpc@v1.40.0
+ google.golang.org/protobuf@v1.27.1

# 0.0.3

## 改进

+ 增加对xds负载均衡的支持
+ 优化结构提高可维护性

# 0.0.2

## 改进

+ 增加了对dns的支持,现在dns的请求会做负载均衡
+ 删除了对zk服务注册服务发现的支持
+ 改进了结构便于扩展

# v0.0.1

项目创建
