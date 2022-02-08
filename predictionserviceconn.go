package tensorflow_service_apis

import (
	"context"
	"time"

	tfserv "github.com/Golang-Tools/tensorflow_service_apis/tensorflow_serving/apis"
	grpc "google.golang.org/grpc"
)

//PredictionServiceConn 客户端类
type PredictionServiceConn struct {
	tfserv.PredictionServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
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
	c.sdk = sdk
	return c, nil
}

//Close 断开连接
func (c *PredictionServiceConn) Close() error {
	return c.conn.Close()
}

func (c *PredictionServiceConn) NewCtx() (ctx context.Context, cancel context.CancelFunc) {
	if c.sdk.SDKConfig.Query_Timeout > 0 {
		timeout := time.Duration(c.sdk.SDKConfig.Query_Timeout) * time.Millisecond
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	return
}
