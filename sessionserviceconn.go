package tensorflow_service_apis

import (
	tfserv "github.com/Golang-Tools/tensorflow_service_apis/tensorflow_serving/apis"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

//SessionServiceConn 客户端类
type SessionServiceConn struct {
	tfserv.SessionServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
}

func newSessionServiceConn(sdk *SDK, addr string, opts ...grpc.DialOption) (*SessionServiceConn, error) {
	c := new(SessionServiceConn)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.SessionServiceClient = tfserv.NewSessionServiceClient(conn)
	c.sdk = sdk
	return c, nil
}

//NewSessionServiceConn 建立一个新的连接
func (c *SDK) NewSessionServiceConn() (*SessionServiceConn, error) {
	conn, err := newSessionServiceConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	c.getConnLock.Lock()
	defer c.getConnLock.Unlock()
	c.sessionServiceConn = conn
	return conn, nil
}

func (c *SDK) getSessionServiceConn() *SessionServiceConn {
	c.getConnLock.RLock()
	defer c.getConnLock.RUnlock()
	if c.sessionServiceConn != nil {
		if c.sessionServiceConn.conn.GetState() == connectivity.Shutdown {
			return nil
		}
		return c.sessionServiceConn
	}
	return nil
}

//GetPredictionServiceConn 获取PredictionServiceClient客户端连接
func (c *SDK) GetSessionServiceConn() (*SessionServiceConn, error) {
	conn := c.getSessionServiceConn()
	if conn != nil {
		return conn, nil
	}
	return c.NewSessionServiceConn()
}

//Close 断开连接
func (c *SessionServiceConn) Close() error {
	return c.conn.Close()
}
