package shopserver1

import (
	"context"
	pb "sniper-api/rpc/shop/v1"
)

// Server 实现 /twirp/shop.v1.Shop 服务
type Server struct{}

// Echo 实现 /twirp/shop.v1.Shop/Echo 接口
func (s *Server) Echo(ctx context.Context, req *pb.EchoReq) (resp *pb.EchoResp, err error) {
	return &pb.EchoResp{Msg: req.Msg}, nil
}
