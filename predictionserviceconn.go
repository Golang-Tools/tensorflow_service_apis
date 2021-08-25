package tensorflow_service_apis
import (
	tfserv "github.com/Golang-Tools/tensorflow_service_apis/tensorflow/tensorflow_serving"
	grpc "google.golang.org/grpc"
)
//PredictionServiceConn 客户端类
type PredictionServiceConn struct {
	tfserv.PredictionServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
	once bool
}


//ModelServiceConn 建立一个新的连接
func (c *SDK) NewPredictionServiceConn() (*PredictionServiceConn, error) {
	conn, err := newPredictionServiceConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}



func newPredictionServiceConn(sdk *SDK, addr string, opts ...grpc.DialOption) (*PredictionServiceConn, error) {
	c := new(PredictionServiceConn)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.PredictionServiceClient = tfserv.NewPredictionServiceClient(conn)
	return c, nil
}

//Close 断开连接
func (c *PredictionServiceConn) Close() error {
	return c.conn.Close()
}

