// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/echo.proto
package service

import (
	"golang.org/x/net/context"

	"github.com/gomeet/gomeet/utils/log"

	pb "github.com/gomeet-examples/svc-echo/pb"
)

func (s *echoServer) Version(ctx context.Context, req *pb.EmptyMessage) (*pb.VersionResponse, error) {
	log.Debug(ctx, "message call", log.Fields{"req": req})

	v := &pb.VersionResponse{
		Name:    name,
		Version: version,
	}

	return v, nil
}
