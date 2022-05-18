package tensorflow_service_apis

import (
	tfserv "github.com/Golang-Tools/tensorflow_service_apis/tensorflow_serving/apis"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

//PredictionServiceConn PredictionServiceClient客户端的连接类
type PredictionServiceConn struct {
	tfserv.PredictionServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
}

func newPredictionServiceConn(sdk *SDK, addr string, opts ...grpc.DialOption) (*PredictionServiceConn, error) {
	c := new(PredictionServiceConn)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.PredictionServiceClient = tfserv.NewPredictionServiceClient(conn)
	c.sdk = sdk
	return c, nil
}

//NewPredictionServiceConn 建立一个新的连接
func (c *SDK) NewPredictionServiceConn() (*PredictionServiceConn, error) {
	conn, err := newPredictionServiceConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	c.getConnLock.Lock()
	defer c.getConnLock.Unlock()
	c.predictionServiceConn = conn
	return conn, nil
}

func (c *SDK) getPredictionServiceConn() *PredictionServiceConn {
	c.getConnLock.RLock()
	defer c.getConnLock.RUnlock()
	if c.predictionServiceConn != nil {
		if c.predictionServiceConn.conn.GetState() == connectivity.Shutdown {
			return nil
		}
		return c.predictionServiceConn
	}
	return nil
}

//GetPredictionServiceConn 获取PredictionServiceClient客户端连接
func (c *SDK) GetPredictionServiceConn() (*PredictionServiceConn, error) {
	conn := c.getPredictionServiceConn()
	if conn != nil {
		return conn, nil
	}
	return c.NewPredictionServiceConn()
}

//Close 断开连接
func (c *PredictionServiceConn) Close() error {
	return c.conn.Close()
}
