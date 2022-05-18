package tensorflow_service_apis

import (
	"github.com/Golang-Tools/grpcsdk"
	"github.com/Golang-Tools/tensorflow_service_apis/v2/pb/tensorflow_serving/apis"
)

var DefaultPredictionSDK = grpcsdk.New(apis.NewPredictionServiceClient, &apis.PredictionService_ServiceDesc)
var DefaultModelServiceSDK = grpcsdk.New(apis.NewModelServiceClient, &apis.ModelService_ServiceDesc)
var DefaultSessionServiceSDK = grpcsdk.New(apis.NewSessionServiceClient, &apis.SessionService_ServiceDesc)
