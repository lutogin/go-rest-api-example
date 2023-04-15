package user

import (
	"context"
	userDto "ms-gateway/internal/user/dto"
)

type Repository interface {
	Create(ctx context.Context, payload userDto.CreateUserDto) (id string, err error)
	GetById(ctx context.Context, payload userDto.GetUserByIdDto) (user UserEntity, err error)
	GetAll(ctx context.Context, payload userDto.GetUsersDto) (user []UserEntity, err error)
	Update(ctx context.Context, payload userDto.UpdateUserDto) (user UserEntity, err error)
	Delete(ctx context.Context, payload userDto.DeleteUserDto) (err error)
}
