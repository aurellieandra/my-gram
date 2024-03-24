package service

import (
	"context"
	"fmt"
	"time"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
	"github.com/aurellieandra/my-gram/pkg/helper"
)

// INTERFACE
type UserService interface {
	Register(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, credentials model.User) (model.User, error)

	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id uint64) (*model.User, error)

	UpdateUserById(ctx context.Context, user model.User, id uint64) (model.User, error)
	DeleteUserById(ctx context.Context, id uint64) error

	// MISC
	GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error)
}

// STRUCT
type userServiceImpl struct {
	repo    repository.UserQuery
	command repository.UserCommand
}

// NEW USER SERVICE
func NewUserService(repo repository.UserQuery, command repository.UserCommand) UserService {
	return &userServiceImpl{repo: repo, command: command}
}

// USER SERVICE IMPL
func (u *userServiceImpl) Register(ctx context.Context, user model.User) (model.User, error) {
	user = model.User{
		Username: user.Username,
		Email:    user.Email,
		Dob:      user.Dob,
	}

	pass, err := helper.GenerateHash(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = pass

	registeredUser, err := u.command.Register(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return registeredUser, nil
}

func (u *userServiceImpl) Login(ctx context.Context, credentials model.User) (model.User, error) {
	credentials = model.User{
		Username: credentials.Username,
	}

	pass, err := helper.GenerateHash(credentials.Password)
	if err != nil {
		return model.User{}, err
	}
	credentials.Password = pass

	loggedInUser, err := u.command.Login(ctx, credentials)
	if err != nil {
		return model.User{}, err
	}
	return loggedInUser, nil
}

func (u *userServiceImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (u *userServiceImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	user, err := u.repo.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userServiceImpl) UpdateUserById(ctx context.Context, user model.User, id uint64) (model.User, error) {
	user = model.User{
		Username: user.Username,
		Email:    user.Email,
		Dob:      user.Dob,
	}

	updatedUser, err := u.command.UpdateUserById(ctx, user, id)
	if err != nil {
		return model.User{}, err
	}
	return updatedUser, nil
}

func (u *userServiceImpl) DeleteUserById(ctx context.Context, id uint64) error {
	err := u.command.DeleteUserById(ctx, id)
	return err
}

// MISC
func (u *userServiceImpl) GenerateUserAccessToken(ctx context.Context, user model.User) (token string, err error) {
	now := time.Now()

	claim := model.StandardClaim{
		Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
		Iss: "my-gram",
		Aud: "golang-006",
		Sub: "access-token",
		Exp: uint64(now.Add(time.Hour).Unix()),
		Iat: uint64(now.Unix()),
		Nbf: uint64(now.Unix()),
	}

	data, err := u.command.GenerateUserAccessToken(ctx, user)
	if err != nil {
		return token, nil
	}

	userResponse := model.UserResponse{
		ID:        data.ID,
		Username:  data.Username,
		Email:     data.Email,
		Dob:       data.Dob,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}

	userClaim := model.AccessClaim{
		StandardClaim: claim,
		UserId:        user.ID,
		Data:          userResponse,
	}

	token, err = helper.GenerateToken(userClaim)
	return
}
