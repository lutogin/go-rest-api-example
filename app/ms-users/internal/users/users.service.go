package users

import (
	"context"
	userDto "ms-users/internal/users/dto"
	"ms-users/pkg/logging"
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

func (s *Service) GetById(ctx context.Context, payload userDto.GetUserByIdDto) (user UserEntity, err error) {
	result, err := s.repo.GetById(ctx, payload)
	return result, err
}
