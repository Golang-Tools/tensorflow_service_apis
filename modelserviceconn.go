package tensorflow_service_apis
import (
	tfserv "github.com/Golang-Tools/tensorflow_service_apis/tensorflow/tensorflow_serving"
	grpc "google.golang.org/grpc"
)
//ModelServiceConn 客户端类
type ModelServiceConn struct {
	tfserv.ModelServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
	once bool
}


//ModelServiceConn 建立一个新的连接
func (c *SDK) NewModelServiceConn() (*ModelServiceConn, error) {
	conn, err := newModelServiceConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func newModelServiceConn(sdk *SDK, addr string, opts ...grpc.DialOption) (*ModelServiceConn, error) {
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

