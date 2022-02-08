# tensorflow_service_apis

tensorflow_service的grpc客户端接口封装.

## 使用方法

1. 创建并初始化SDK对象,有3种方式
   1. 可以使用`tensorflow_service_apis.New()`方式,然后调用sdk对象的Init方法传入参数`tensorflow_service_apis.SDKConfig`对象初始化
   2. 构造一个`tensorflow_service_apis.SDKConfig`对象,调用其`NewSDK()`方法创建一个初始化化好的`SDK`对象
   3. 直接使用默认`SDK`对象`tensorflow_service_apis.DefaultSDK`,调用sdk对象的Init方法传入参数`tensorflow_service_apis.SDKConfig`对象初始化
2. 调用SDK对象的`Get{Session|Model|Prediction}ServiceConn`方法获取对应的连接
3. 调用SDK对象的`NewCtx`方法获取请求时用的ctx对象
4. 构造指定方法的请求体
5. 请求指定方法并获得结果

## 例子

此处以调用`ModelService`的`GetModelStatus`方法为例

```go
import (
    tfsapis "github.com/Golang-Tools/tensorflow_service_apis"
    "google.golang.org/protobuf/types/known/wrapperspb"
    log "github.com/Golang-Tools/loggerhelper"
)

func main(){
    tfsapis.DefaultSDK.Init(&tensorflow_service_apis.SDKConfig{
        //你的配置
    })
    conn,err := tfsapis.DefaultSDK.GetModelServiceConn()
    if err != nil{
        panic(err)
    }
    // 获取模型元信息
    ctx,cancel := tfsapis.DefaultSDK.NewCtx()
    defer cancel()
    res, err := conn.GetModelStatus(ctx, &tfsapis.GetModelStatusRequest{
        ModelSpec:&tfsapis.ModelSpec{
            Name:          {modelName},//模型名
            VersionChoice: &tfsapis.ModelSpec_Version{Version: wrapperspb.Int64({version})},//指定版本号
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
2. 执行leave_proto.py文件
3. 使用搜索工具,查找.go文件中的`"tensorflow`,找到import中的内容,前面加上`github.com/Golang-Tools/tensorflow_service_apis/`
4. 在`tensorflow`和`tensorflow_serving`两个文件夹下分别添加一个同名`.go`文件,在其中添加同名`package`声明

    ```go
    package tensorflow|tensorflow_serving
    ```

一般情况下我们只需要修改项目根目录下的4个文件

+ `tensorflow_service_apis.go`,sdk对象的声明,包括各种设置项的处理等
+ `modelserviceconn.go`,`predictionserviceconn.go`和`sessionserviceconn.go`,不同service的连接对象声明