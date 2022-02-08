package tensorflow_service_apis

import (
	tfserv "github.com/Golang-Tools/tensorflow_service_apis/tensorflow_serving/apis"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

//ModelServiceConn ModelServiceClient客户端的连接类
type ModelServiceConn struct {
	tfserv.ModelServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
}

func newModelServiceConn(sdk *SDK, addr string, opts ...grpc.DialOption) (*ModelServiceConn, error) {
	c := new(ModelServiceConn)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.ModelServiceClient = tfserv.NewModelServiceClient(conn)
	c.sdk = sdk
	return c, nil
}

//NewModelServiceConn 建立一个新的ModelServiceClient客户端的连接类
func (c *SDK) NewModelServiceConn() (*ModelServiceConn, error) {
	conn, err := newModelServiceConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	c.getConnLock.Lock()
	defer c.getConnLock.Unlock()
	c.modelServiceConn = conn
	return conn, nil
}

func (c *SDK) getModelServiceConn() *ModelServiceConn {
	c.getConnLock.RLock()
	defer c.getConnLock.RUnlock()
	if c.modelServiceConn != nil {
		if c.modelServiceConn.conn.GetState() == connectivity.Shutdown {
			return nil
		}
		return c.modelServiceConn
	}
	return nil
}

//GetModelServiceConn 获取ModelServiceClient客户端连接
func (c *SDK) GetModelServiceConn() (*ModelServiceConn, error) {
	conn := c.getModelServiceConn()
	if conn != nil {
		return conn, nil
	}
	return c.NewModelServiceConn()
}

//Close 断开连接
func (c *ModelServiceConn) Close() error {
	return c.conn.Close()
}
