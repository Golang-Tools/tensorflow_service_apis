# v2.0.2

## 实现优化

+ 更新`github.com/Golang-Tools/grpcsdk`至v0.0.2
+ 更新`google.golang.org/grpc`至v1.46.2
+ 更新`google.golang.org/protobuf`至v1.28.0

# v2.0.1

现在本项目基于[github.com/Golang-Tools/grpcsdk](https://github.com/Golang-Tools/grpcsdk).因此使用方法和`grpcsdk`构造的sdk完全一致.

原本的sdk对象通过3个方法获得3个服务接口对象改为每个服务接口一个对应的默认sdk,推荐直接使用默认sdk对象而不是自己再创建.

# v2.0.0

大版本更新,支持仅支持go 1.18+.使用泛型.接口也有较大变动
