package users

import (
	"context"
	userDto "ms-users/internal/users/dto"
)

type Repository interface {
	Create(ctx context.Context, payload userDto.CreateUserDto) (id string, err error)
	GetById(ctx context.Context, payload userDto.GetUserByIdDto) (user UserEntity, err error)
	GetByFilter(ctx context.Context, payload userDto.GetUsersDto) (user []UserEntity, err error)
	Update(ctx context.Context, payload userDto.UpdateUserDto) (err error)
	Delete(ctx context.Context, payload userDto.DeleteUserDto) (err error)
}
