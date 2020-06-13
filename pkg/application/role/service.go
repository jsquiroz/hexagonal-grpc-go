package role

import "github.com/jsquiroz/hexagonal-grpc-go/pkg/domain/role"

// Service provides role adding operation
type Service interface {
	AddRole(...role.Role)
}

type service struct {
	aR role.Repository
}

// NewService create an adding service with the necesary dependencies
func NewService(r role.Repository) Service {
	return &service{r}
}

func (s *service) AddRole(a ...role.Role) {
	// TODO: Validate something
	for _, role := range a {
		_ = s.aR.AddRole(role) // TODO: Error handler
	}
}
