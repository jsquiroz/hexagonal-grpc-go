package handler

import (
	"context"

	app_role "github.com/jsquiroz/hexagonal-grpc-go/pkg/application/role"
	"github.com/jsquiroz/hexagonal-grpc-go/pkg/domain/role"
	grpc_role "github.com/jsquiroz/hexagonal-grpc-go/pkg/infrastructure/delivery/grpc/proto/role"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Service is a WorkShift Interface, contains the same methods
// that protobuf
type Service struct {
	serv app_role.Service
}

// NewRoleServerGrpc ...
func NewRoleServerGrpc(gserver *grpc.Server, addserv app_role.Service) {

	attenserver := &Service{
		serv: addserv,
	}

	grpc_role.RegisterRoleServiceServer(gserver, attenserver)
	reflection.Register(gserver)
}

func (s *Service) parseToGRPC(a *role.Role) *grpc_role.Role {

	at := &grpc_role.Role{
		Idcompany: uint32(a.IDCompany),
		Name:      a.Name,
	}

	return at
}

func (s *Service) parseToData(a *grpc_role.Role) *role.Role {

	ad := &role.Role{
		IDCompany: uint(a.Idcompany),
		Name:      a.Name,
	}

	return ad
}

// Create ...
func (s *Service) Create(ctx context.Context, u *grpc_role.CreateRequest) (*grpc_role.CreateResponse, error) {
	// FIXME: implement errors
	a := s.parseToData(u.Role)

	s.serv.AddRole(*a)

	return &grpc_role.CreateResponse{}, nil
}
