package service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gomeet/gomeet/utils/log"

	pb "github.com/gomeet-examples/svc-echo/pb"
)

func (s *echoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Debug(ctx, "service call", log.Fields{"req": req})

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res := &pb.EchoResponse{
		Uuid:    req.GetUuid(),
		Content: req.GetContent(),
	}

	return res, nil
}
