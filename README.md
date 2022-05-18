# tensorflow_service_apis/v2

tensorflow_service的grpc客户端接口封装.

v2版本将只支持go 1.18+ 低版本的请使用v0版本.



## 使用方法

本项目基于[github.com/Golang-Tools/grpcsdk](https://github.com/Golang-Tools/grpcsdk).因此使用方法和`grpcsdk`构造的sdk完全一致.

本项目只是提供了

1. `pb`模块,`tenserflow`和`tensorflow_serving`的protobuf文件生成的go模块

2. 3个默认的sdk对象
    + `DefaultPredictionSDK`对应`PredictionService`
    + `DefaultModelServiceSDK`对应`ModelService`
    + `DefaultSessionServiceSDK`对应`SessionService`

## 例子

此处以调用`ModelService`的`GetModelStatus`方法为例

```go
import (
    tfsapis "github.com/Golang-Tools/tensorflow_service_apis"
    apipb "github.com/Golang-Tools/tensorflow_service_apis/pb/tensorflow_serving/apis"
    "google.golang.org/protobuf/types/known/wrapperspb"
    "github.com/Golang-Tools/grpcsdk"
)

func main(){
    err := tfsapis.DefaultSDK.Init(grpcsdk.WithQueryAddresses("localhost:5000"))
    if err != nil {
        tfsapis.DefaultSDK.Logger.Error("sdk.Init get error", log.Dict{"err": err.Error()})
    }
    defer tfsapis.DefaultSDK.Close()
    tfsapis.DefaultSDK.Logger.Info("setup sdk init ok")
    Conn, release := tfsapis.DefaultSDK.GetClient()
    defer release()
    tfsapis.DefaultSDK.Logger.Info("setup sdk GetClient ok")
    tfsapis.DefaultSDK.Logger.Info("setup ok")
    // 获取模型元信息
    ctx, cancel := tfsapis.DefaultSDK.NewCtx()
    defer cancel()
    // 获取模型元信息
    res, err := conn.GetModelStatus(ctx, &apipb.GetModelStatusRequest{
        ModelSpec:&apipb.ModelSpec{
            Name:          {modelName},//模型名
            VersionChoice: &apipb.ModelSpec_Version{Version: wrapperspb.Int64({version})},//指定版本号
        },
    })
    if err != nil{
        panic(err)
    }
    log.Info("get model status",log.Dict{"res":res})
}
```

## 注意事项

### `tensorflow.serving.PredictionService/GetModelMetadata`常用来查看模型的元信息

1. 请求这个方法必须填写参数`MetadataField: []string{"signature_def"}`
2. 这个方法的返回中有`any`类型,其对应的是`tensorflow_serving.SignatureDefMap`使用如下方式获取:

    ```golang
    import (
        "github.com/golang/protobuf/ptypes"
        apispb "github.com/Golang-Tools/tensorflow_service_apis/tensorflow_serving"
    )

    func main(){
        sd := apispb.SignatureDefMap{}
        err = ptypes.UnmarshalAny(meta.Metadata["signature_def"], &sd)
    }
    ```

## 开发方式

1. 下载指定版本的tensorflow和tfserving,将其中有用的文件夹(tensorflow/core和tensorflow_serving)留下其他都删除.
2. 执行gen_proto.py文件
3. 使用搜索工具,查找.go文件中的`"tensorflow`,找到import中的内容,前面加上`github.com/Golang-Tools/tensorflow_service_apis/v2`
4. 在`tensorflow`和`tensorflow_serving`两个文件夹下分别添加一个同名`.go`文件,在其中添加同名`package`声明

    ```go
    package tensorflow|tensorflow_serving
    ```

一般情况下我们只需要修改项目根目录下的4个文件

+ `tensorflow_service_apis.go`,sdk对象的声明,包括各种设置项的处理等
+ `modelserviceconn.go`,`predictionserviceconn.go`和`sessionserviceconn.go`,不同service的连接对象声明