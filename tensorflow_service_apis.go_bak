package tensorflow_service_apis

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//SDKConfig 的客户端类型
type SDKConfig struct {
	Address                                 []string `json:"address" jsonschema:"description=连接服务的主机和端口"`
	ServiceName                             string   `json:"service_name,omitempty" jsonschema:"description=服务器域名"`
	AppName                                 string   `json:"app_name,omitempty" jsonschema:"description=服务名"`
	AppVersion                              string   `json:"app_version,omitempty" jsonschema:"description=服务版本"`
	BalanceWithZookeeper                    bool     `json:"balance_with_zookeeper,omitempty" jsonschema:"description=是否使用zookeeper做本地负载均衡"`
	CaCert                                  string   `json:"ca_cert,omitempty" jsonschema:"description=如果要使用tls则需要指定根证书位置"`
	InitialWindowSize                       int      `json:"initial_window_size,omitempty" jsonschema:"description=基于Stream的滑动窗口大小"`
	InitialConnWindowSize                   int      `json:"initial_conn_window_size,omitempty" jsonschema:"description=基于Connection的滑动窗口大小"`
	KeepaliveTime                           int      `json:"keepalive_time,omitempty" jsonschema:"description=空闲连接每隔n秒ping一次客户端已确保连接存活"`
	KeepaliveTimeout                        int      `json:"keepalive_timeout,omitempty" jsonschema:"description=ping时长超过n则认为连接已死"`
	KeepaliveEnforcementPermitWithoutStream bool     `json:"keepalive_enforement_permit_without_stream,omitempty" jsonschema:"description=是否当连接空闲时仍然发送PING帧监测"`
	ConnWithBlock                           bool     `json:"conn_with_block,omitempty" jsonschema:"description=同步的连接建立"`
	MaxRecvMsgSize                          int      `json:"max_rec_msg_size,omitempty" jsonschema:"description=允许接收的最大消息长度"`
	MaxSendMsgSize                          int      `json:"max_send_msg_size,omitempty" jsonschema:"description=允许发送的最大消息长度"`
}

//NewModelServiceSDK 创建ModelService客户端对象
func (c *SDKConfig) NewModelServiceSDK() *ModelServiceSDK {
	sdk := NewModelServiceSDK()
	sdk.Init(c)
	return sdk
}

//New PredictionServiceSDK 创建PredictionService客户端对象
func (c *SDKConfig) NewPredictionServiceSDK() *PredictionServiceSDK {
	sdk := NewPredictionServiceSDK()
	sdk.Init(c)
	return sdk
}
