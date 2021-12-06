package v1alpha2

import (
	"context"
	"unsafe"

	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func (s *service) RemovePodSandbox(
	ctx context.Context, req *pb.RemovePodSandboxRequest,
) (*pb.RemovePodSandboxResponse, error) {
	if err := s.server.RemovePodSandbox(ctx, (*v1.RemovePodSandboxRequest)(unsafe.Pointer(req))); err != nil {
		return nil, err
	}
	return &pb.RemovePodSandboxResponse{}, nil
}
