package service

import (
	"context"

	pb "github.com/byteflowteam/kratos-vue-admin/api/modelop/service/v1"
)

type ModelopService struct {
	pb.UnimplementedModelopServer
}

func NewModelopService() *ModelopService {
	return &ModelopService{}
}

func (s *ModelopService) CreateModelop(ctx context.Context, req *pb.CreateModelopRequest) (*pb.CreateModelopReply, error) {
	return &pb.CreateModelopReply{}, nil
}
func (s *ModelopService) UpdateModelop(ctx context.Context, req *pb.UpdateModelopRequest) (*pb.UpdateModelopReply, error) {
	return &pb.UpdateModelopReply{}, nil
}
func (s *ModelopService) DeleteModelop(ctx context.Context, req *pb.DeleteModelopRequest) (*pb.DeleteModelopReply, error) {
	return &pb.DeleteModelopReply{}, nil
}
func (s *ModelopService) GetModelop(ctx context.Context, req *pb.GetModelopRequest) (*pb.GetModelopReply, error) {
	return &pb.GetModelopReply{}, nil
}
func (s *ModelopService) ListModelop(ctx context.Context, req *pb.ListModelopRequest) (*pb.ListModelopReply, error) {
	return &pb.ListModelopReply{}, nil
}
