package usercenter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	mw "github.com/ra1n6ow/opsx/internal/pkg/middleware/grpc"
	"github.com/ra1n6ow/opsx/internal/pkg/server"
	handler "github.com/ra1n6ow/opsx/internal/usercenter/handler/grpc"
	ucv1 "github.com/ra1n6ow/opsx/pkg/api/usercenter/v1"
)

// grpcServer 定义一个 gRPC 服务器.
type grpcServer struct {
	srv server.Server
	// stop 为优雅关停函数.
	stop func(context.Context)
}

// 确保 *grpcServer 实现了 server.Server 接口.
var _ server.Server = (*grpcServer)(nil)

// NewGRPCServerOr 创建并初始化 gRPC 或者 gRPC +  gRPC-Gateway 服务器.
func (c *ServerConfig) NewGRPCServerOr() (server.Server, error) {
	// 配置 gRPC 服务器选项，包括拦截器链
	serverOptions := []grpc.ServerOption{
		// 注意拦截器顺序！
		grpc.ChainUnaryInterceptor(
			// 请求 ID 拦截器
			mw.RequestIDInterceptor(),
		),
	}

	// 创建 gRPC 服务器
	grpcsrv, err := server.NewGRPCServer(
		c.cfg.GRPCOptions,
		serverOptions,
		func(s grpc.ServiceRegistrar) {
			ucv1.RegisterUsercenterServer(s, handler.NewHandler())
		},
	)
	if err != nil {
		return nil, err
	}

	// 如果配置为 gRPC 服务器模式，则直接返回 gRPC 服务器实例
	if c.cfg.ServerMode == GRPCServerMode {
		return &grpcServer{
			srv: grpcsrv,
			stop: func(ctx context.Context) {
				grpcsrv.GracefulStop(ctx)
			},
		}, nil
	}

	// 先启动 gRPC 服务器，因为 HTTP 服务器依赖 gRPC 服务器.
	go grpcsrv.RunOrDie()

	httpsrv, err := server.NewGRPCGatewayServer(
		c.cfg.HTTPOptions,
		c.cfg.GRPCOptions,
		func(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
			return ucv1.RegisterUsercenterHandler(context.Background(), mux, conn)
		},
	)
	if err != nil {
		return nil, err
	}

	return &grpcServer{
		srv: httpsrv,
		stop: func(ctx context.Context) {
			grpcsrv.GracefulStop(ctx)
			httpsrv.GracefulStop(ctx)
		},
	}, nil
}

// RunOrDie 启动 gRPC 服务器或 HTTP 反向代理服务器，异常时退出.
func (s *grpcServer) RunOrDie() {
	s.srv.RunOrDie()
}

// GracefulStop 优雅停止 HTTP 和 gRPC 服务器.
func (s *grpcServer) GracefulStop(ctx context.Context) {
	s.stop(ctx)
}
