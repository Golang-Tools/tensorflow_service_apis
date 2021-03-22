package tensorflow_service_apis

import (
	"errors"
	"fmt"
	"strings"
	"time"

	tfserv "Golang-Tools/tensorflow_service_apis/tensorflow/tensorflow_serving"

	log "github.com/Golang-Tools/loggerhelper"
	"github.com/liyue201/grpc-lb/balancer"
	registry "github.com/liyue201/grpc-lb/registry/zookeeper"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	resolver "google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	_ "google.golang.org/grpc/xds"
)

//ModelServiceSDK ModelService的客户端类型
type ModelServiceSDK struct {
	*SDKConfig
	opts          []grpc.DialOption
	serviceconfig map[string]interface{}
	addr          string
}

//New 创建客户端对象
func NewModelServiceSDK() *ModelServiceSDK {
	c := new(ModelServiceSDK)
	c.opts = make([]grpc.DialOption, 0, 10)
	return c
}

//InitCallOpts 初始化连接选项
func (c *ModelServiceSDK) InitCallOpts() {
	callopts := []grpc.CallOption{}
	if c.MaxRecvMsgSize != 0 {
		callopts = append(callopts, grpc.MaxCallRecvMsgSize(c.MaxRecvMsgSize))
	}
	if c.MaxSendMsgSize != 0 {
		callopts = append(callopts, grpc.MaxCallSendMsgSize(c.MaxSendMsgSize))
	}
	if len(callopts) > 0 {
		c.opts = append(c.opts, grpc.WithDefaultCallOptions(callopts...))
	}
}

//InitOpts 初始化连接选项
func (c *ModelServiceSDK) InitOpts() error {
	c.opts = append(c.opts)
	if c.CaCert != "" {
		creds, err := credentials.NewClientTLSFromFile(c.CaCert, "")
		if err != nil {
			log.Error("failed to load credentials", log.Dict{"err": err.Error()})
			return err
		}
		c.opts = append(c.opts, grpc.WithTransportCredentials(creds))
	} else {
		c.opts = append(c.opts, grpc.WithInsecure())
	}
	if c.KeepaliveTime != 0 || c.KeepaliveTimeout != 0 || c.KeepaliveEnforcementPermitWithoutStream == true {
		kacp := keepalive.ClientParameters{
			Time:                time.Duration(c.KeepaliveTime) * time.Second,
			Timeout:             time.Duration(c.KeepaliveTimeout) * time.Second,
			PermitWithoutStream: c.KeepaliveEnforcementPermitWithoutStream, // send pings even without active streams
		}
		c.opts = append(c.opts, grpc.WithKeepaliveParams(kacp))
	}
	if c.ConnWithBlock == true {
		c.opts = append(c.opts, grpc.WithBlock())
	}
	if c.InitialWindowSize != 0 {
		c.opts = append(c.opts, grpc.WithInitialWindowSize(int32(c.InitialWindowSize)))
	}
	if c.InitialConnWindowSize != 0 {
		c.opts = append(c.opts, grpc.WithInitialConnWindowSize(int32(c.InitialConnWindowSize)))
	}

	return nil
}

//Init 初始化ModelService客户端的连接信息
func (c *ModelServiceSDK) Init(conf *SDKConfig) error {
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
			if conf.BalanceWithZookeeper {
				c.initWithZooKeeperBalance()
			} else {
				c.initStandalone()
			}
		}
	default:
		{
			if conf.BalanceWithZookeeper {
				c.initWithZooKeeperBalance()
			} else {
				c.initWithLocalBalance()
			}
		}
	}
	err := c.InitOpts()
	if err != nil {
		return err
	}
	c.InitCallOpts()
	if c.serviceconfig != nil {
		serviceconfig, err := json.MarshalToString(c.serviceconfig)
		if err != nil {
			return err
		}
		c.opts = append(c.opts, grpc.WithDefaultServiceConfig(serviceconfig))
	}
	return nil
}

//InitStandalone 初始化单机服务的连接配置
func (c *ModelServiceSDK) initStandalone() error {
	c.addr = c.Address[0]
	return nil
}

//InitWithLocalBalance 初始化本地负载均衡的连接配置
func (c *ModelServiceSDK) initWithLocalBalance() error {
	serverName := ""
	if c.AppName != "" {
		if c.AppVersion != "" {
			serverName = fmt.Sprintf("%s-%s", c.AppName, strings.ReplaceAll(c.AppVersion, ".", "_"))
		} else {
			serverName = c.AppName
		}

	}
	if c.serviceconfig == nil {
		c.serviceconfig = map[string]interface{}{
			"loadBalancingPolicy": "round_robin",
			"healthCheckConfig":   map[string]string{"serviceName": c.ServiceName},
		}
	} else {
		c.serviceconfig["loadBalancingPolicy"] = "round_robin"
		c.serviceconfig["healthCheckConfig"] = map[string]string{"serviceName": serverName}
	}
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
	return nil
}

//InitWithZooKeeperBalance 初始zokeeper负载均衡的连接配置
func (c *ModelServiceSDK) initWithZooKeeperBalance() error {
	registry.RegisterResolver("zk", c.Address, "/backend/services", c.AppName, c.AppVersion)
	c.addr = "zk:///"
	c.opts = append(c.opts, grpc.WithBalancerName(balancer.RoundRobin))
	return nil
}

//NewConn 建立一个新的连接
func (c *ModelServiceSDK) NewConn() (*ModelServiceConn, error) {
	conn, err := newModelServiceConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

//ModelServiceConn ModelService客户端类
type ModelServiceConn struct {
	tfserv.ModelServiceClient
	conn *grpc.ClientConn
	sdk  *ModelServiceSDK
	once bool
}

func newModelServiceConn(sdk *ModelServiceSDK, addr string, opts ...grpc.DialOption) (*ModelServiceConn, error) {
	c := new(ModelServiceConn)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.ModelServiceClient = tfserv.NewModelServiceClient(conn)
	return c, nil
}

//Close 断开连接
func (c *ModelServiceConn) Close() error {
	return c.conn.Close()
}

//DefaultModelServiceSDK 默认的DefaultModelServiceSDK客户端
var DefaultModelServiceSDK = NewModelServiceSDK()
