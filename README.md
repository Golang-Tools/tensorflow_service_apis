# tensorflow_service_apis

tensorflow_service的grpc客户端接口封装.比较常用的接口如下,下面是对应的使用注意点

## `tensorflow.serving.PredictionService/GetModelMetadata`常用来查看模型的元信息

1. 请求这个方法必须填写参数`MetadataField: []string{"signature_def"}`
2. 这个方法的返回中有`any`类型,其对应的是`tensorflow_serving.SignatureDefMap`使用如下方式获取:

    ```golang
    import (
        "github.com/golang/protobuf/ptypes"
        apispb "github.com/Golang-Tools/tensorflow_service_apis/tensorflow/tensorflow_serving"
    )

    func main(){
        sd := apispb.SignatureDefMap{}
        err = ptypes.UnmarshalAny(meta.Metadata["signature_def"], &sd)
    }
    ```

