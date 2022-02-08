package tensorflow_service_apis

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	log "github.com/Golang-Tools/loggerhelper"
	jsoniter "github.com/json-iterator/go"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	resolver "google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	_ "google.golang.org/grpc/xds"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//SDKConfig 的客户端类型
type SDKConfig struct {
	Address      []string `json:"address" jsonschema:"description=连接服务的主机和端口"`
	Service_Name string   `json:"service_name,omitempty" jsonschema:"description=服务器域名"`
	App_Name     string   `json:"app_name,omitempty" jsonschema:"description=服务名"`
	App_Version  string   `json:"app_version,omitempty" jsonschema:"description=服务版本"`

	// 性能设置
	Initial_Window_Size                         int  `json:"initial_window_size,omitempty" jsonschema:"description=基于Stream的滑动窗口大小"`
	Initial_Conn_Window_Size                    int  `json:"initial_conn_window_size,omitempty" jsonschema:"description=基于Connection的滑动窗口大小"`
	Keepalive_Time                              int  `json:"keepalive_time,omitempty" jsonschema:"description=空闲连接每隔n秒ping一次客户端已确保连接存活"`
	Keepalive_Timeout                           int  `json:"keepalive_timeout,omitempty" jsonschema:"description=ping时长超过n则认为连接已死"`
	Keepalive_Enforcement_Permit_Without_Stream bool `json:"keepalive_enforement_permit_without_stream,omitempty" jsonschema:"description=是否当连接空闲时仍然发送PING帧监测"`
	Conn_With_Block                             bool `json:"conn_with_block,omitempty" jsonschema:"description=同步的连接建立"`
	Max_Recv_Msg_Size                           int  `json:"max_rec_msg_size,omitempty" jsonschema:"description=允许接收的最大消息长度"`
	Max_Send_Msg_Size                           int  `json:"max_send_msg_size,omitempty" jsonschema:"description=允许发送的最大消息长度"`

	//压缩设置,目前只支持gzip
	Compression string `json:"compression,omitempty" jsonschema:"description=使用哪种方式压缩发送的消息,enum=gzip"`

	// TLS设置
	Ca_Cert_Path     string `json:"ca_cert_path,omitempty" jsonschema:"description=如果要使用tls则需要指定根证书位置"`
	Client_Cert_Path string `json:"client_cert_path,omitempty" jsonschema:"description=客户端整数位置"`
	Client_Key_Path  string `json:"client_key_path,omitempty" jsonschema:"description=客户端证书对应的私钥位置"`

	// XDS设置
	XDS_CREDS bool `json:"xds_creds,omitempty" jsonschema:"description=当address的schema是xds时是否使用xds的令牌加密访问"`

	// 请求超时设置
	Query_Timeout int `json:"query_timeout,omitempty" jsonschema:"description=请求服务的最大超时时间单位ms"`
}

//NewSDK 创建客户端对象
func (c *SDKConfig) NewSDK() *SDK {
	sdk := New()
	sdk.Init(c)
	return sdk
}

//SDK 的客户端类型
type SDK struct {
	*SDKConfig
	opts          []grpc.DialOption
	callopts      []grpc.CallOption
	serviceconfig map[string]interface{}
	addr          string
}

//New 创建客户端对象
func New() *SDK {
	c := new(SDK)
	c.opts = []grpc.DialOption{}
	c.callopts = []grpc.CallOption{}
	c.serviceconfig = map[string]interface{}{}
	return c
}

//initMsgSize 初始化消息大小设置
func (c *SDK) initMsgSize() {
	if c.Max_Recv_Msg_Size != 0 {
		c.callopts = append(c.callopts, grpc.MaxCallRecvMsgSize(c.Max_Recv_Msg_Size))
	}
	if c.Max_Send_Msg_Size != 0 {
		c.callopts = append(c.callopts, grpc.MaxCallSendMsgSize(c.Max_Send_Msg_Size))
	}

}

//initCompression 初始化压缩设置
func (c *SDK) initCompression() {
	switch c.Compression {
	case "gzip":
		{
			c.callopts = append(c.callopts, grpc.UseCompressor(gzip.Name))
		}
	}
}

//initKeepalive 初始化keepalive的相关设置
func (c *SDK) initKeepalive() {
	if c.Keepalive_Time != 0 || c.Keepalive_Timeout != 0 || c.Keepalive_Enforcement_Permit_Without_Stream {
		kacp := keepalive.ClientParameters{
			Time:                time.Duration(c.Keepalive_Time) * time.Second,
			Timeout:             time.Duration(c.Keepalive_Timeout) * time.Second,
			PermitWithoutStream: c.Keepalive_Enforcement_Permit_Without_Stream, // send pings even without active streams
		}
		c.opts = append(c.opts, grpc.WithKeepaliveParams(kacp))
	}
}

//initPerformanceOpts 初始化连接的性能选项
func (c *SDK) initPerformanceOpts() {
	if c.Conn_With_Block {
		c.opts = append(c.opts, grpc.WithBlock())
	}
	if c.Initial_Window_Size != 0 {
		c.opts = append(c.opts, grpc.WithInitialWindowSize(int32(c.Initial_Window_Size)))
	}
	if c.Initial_Conn_Window_Size != 0 {
		c.opts = append(c.opts, grpc.WithInitialConnWindowSize(int32(c.Initial_Conn_Window_Size)))
	}
}

//Init 初始化sdk客户端的连接信息
func (c *SDK) Init(conf *SDKConfig) error {
	c.SDKConfig = conf
	if conf.Address == nil {
		return errors.New("必须至少有一个地址")
	}
	switch len(conf.Address) {
	case 0:
		{
			return errors.New("必须至少有一个地址")
		}
	case 1:
		{
			c.addr = c.Address[0]
			if strings.HasPrefix(c.addr, "dns:///") {
				err := c.initWithDNSLB()
				if err != nil {
					return err
				}
			} else if strings.HasPrefix(c.addr, "xds:///") {
				err := c.initWithXDSLB()
				if err != nil {
					return err
				}
			} else {
				err := c.initWithoutBL()
				if err != nil {
					return err
				}
			}
		}
	default:
		{
			err := c.initWithLocalBalance()
			if err != nil {
				return err
			}
		}
	}
	c.initMsgSize()
	c.initCompression()
	c.initKeepalive()
	c.initPerformanceOpts()
	if len(c.serviceconfig) != 0 {
		serviceconfig, err := json.MarshalToString(c.serviceconfig)
		if err != nil {
			return err
		}
		c.opts = append(c.opts, grpc.WithDefaultServiceConfig(serviceconfig))
	}
	if len(c.callopts) != 0 {
		c.opts = append(c.opts, grpc.WithDefaultCallOptions(c.callopts...))
	}
	return nil
}

// initTLS 初始化TLS设置
func (c *SDK) initTLS() error {
	if c.Ca_Cert_Path != "" {
		if c.Client_Cert_Path != "" && c.Client_Key_Path != "" {
			cert, err := tls.LoadX509KeyPair(c.Client_Cert_Path, c.Client_Key_Path)
			if err != nil {
				log.Error("read client pem file error:", log.Dict{"err": err.Error(), "Cert_path": c.Client_Cert_Path, "Key_Path": c.Client_Key_Path})
				return err
			}
			capool := x509.NewCertPool()
			caCrt, err := ioutil.ReadFile(c.Ca_Cert_Path)
			if err != nil {
				log.Error("read ca pem file error:", log.Dict{"err": err.Error(), "path": c.Ca_Cert_Path})
				return err
			}
			capool.AppendCertsFromPEM(caCrt)
			tlsconf := &tls.Config{
				RootCAs:      capool,
				Certificates: []tls.Certificate{cert},
			}
			creds := credentials.NewTLS(tlsconf)
			c.opts = append(c.opts, grpc.WithTransportCredentials(creds))
		} else {
			creds, err := credentials.NewClientTLSFromFile(c.Ca_Cert_Path, "")
			if err != nil {
				log.Error("failed to load credentials", log.Dict{"err": err.Error()})
				return err
			}
			c.opts = append(c.opts, grpc.WithTransportCredentials(creds))
		}
	} else {
		c.opts = append(c.opts, grpc.WithInsecure())
	}
	return nil
}

//initWithoutBL 初始化没有堵在均衡设置的服务
func (c *SDK) initWithoutBL() error {
	return c.initTLS()
}

//initWithDNSLB 初始化使用外部dns做负载均衡的设置
func (c *SDK) initWithDNSLB() error {
	c.serviceconfig["loadBalancingPolicy"] = "round_robin"
	return c.initTLS()
}

//InitWithStandalone 初始化使用XDS协议做负载均衡的设置
func (c *SDK) initWithXDSLB() error {
	creds := insecure.NewCredentials()
	var err error
	if c.XDS_CREDS {
		creds, err = xdscreds.NewClientCredentials(xdscreds.ClientOptions{FallbackCreds: insecure.NewCredentials()})
	}
	if err != nil {
		return err
	}
	c.opts = append(c.opts, grpc.WithTransportCredentials(creds))
	return nil
}

//InitWithLocalBalance 初始化本地负载均衡的连接配置
func (c *SDK) initWithLocalBalance() error {
	serverName := ""
	if c.App_Name != "" {
		if c.App_Version != "" {
			serverName = fmt.Sprintf("%s-%s", c.App_Name, strings.ReplaceAll(c.App_Version, ".", "_"))
		} else {
			serverName = c.App_Name
		}
	}
	c.serviceconfig["loadBalancingPolicy"] = "round_robin"
	c.serviceconfig["healthCheckConfig"] = map[string]string{"serviceName": serverName}

	r := manual.NewBuilderWithScheme("localbalancer")
	addresses := []resolver.Address{}
	for _, addr := range c.Address {
		addresses = append(addresses, resolver.Address{Addr: addr})
	}
	r.InitialState(resolver.State{
		Addresses: addresses,
	})
	c.addr = fmt.Sprintf("%s:///%s", r.Scheme(), serverName)
	c.opts = append(c.opts, grpc.WithResolvers(r))
	return c.initTLS()
}

//DefaultSDK 默认的sdk客户端
var DefaultSDK = New()
