package service

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
)

// INTERFACE
type UserService interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id uint64) (model.User, error)
	// Register(ctx context.Context, user model.User) (model.User, error)
	// Login(ctx context.Context) (model.User, error)
	// Logout(ctx context.Context) error
	UpdateUserById(ctx context.Context, id uint64) (model.User, error)
	DeleteUserById(ctx context.Context, id uint64) error
}

// STRUCT
type userServiceImpl struct {
	repo repository.UserQuery
	command repository.UserCommand
}

// NEW USER SERVICE
func NewUserService(repo repository.UserQuery, command repository.UserCommand) UserService {
	return &userServiceImpl{repo:repo, command:command}
}

// USER SERVICE IMPL
func (u *userServiceImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}
func (u *userServiceImpl) GetUserById(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}
func (u *userServiceImpl) Register(ctx context.Context, user model.User) (model.User, error) {
	registeredUser, err := u.command.Register(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	return registeredUser, nil
}
func (u *userServiceImpl) Login(ctx context.Context) (model.User, error) {
    loggedInUser, err := u.command.Login(ctx)
    if err != nil {
        return model.User{}, err
    }
    return loggedInUser, nil
}
func (u *userServiceImpl) Logout(ctx context.Context) error {
    err := u.command.Logout(ctx)
    return err
}
func (u *userServiceImpl) UpdateUserById(ctx context.Context, id uint64) (model.User, error) {
    updatedUser, err := u.command.UpdateUserById(ctx, id)
    if err != nil {
        return model.User{}, err
    }
    return updatedUser, nil
}
func (u *userServiceImpl) DeleteUserById(ctx context.Context, id uint64) error {
    err := u.command.DeleteUserById(ctx, id)
    return err
}