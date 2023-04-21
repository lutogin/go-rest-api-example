package user

import (
	"context"
	userDto "ms-gateway/internal/user/dto"
	"ms-gateway/pkg/logging"
)

type Service struct {
	repo   Repository
	logger *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *Service {
	logger.Logger.Infoln("Registering service.")
	return &Service{repo: repository, logger: logger}
}

func (s *Service) Create(ctx context.Context, payload userDto.CreateUserDto) (id string, err error) {
	result, err := s.repo.Create(ctx, payload)
	return result, err
}
