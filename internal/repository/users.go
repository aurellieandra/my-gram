package repository

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type UserQuery interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id uint64) (model.User, error)
}
type UserCommand interface {
	// Register(ctx context.Context, user model.User) (model.User, error)
	// Login(ctx context.Context) (model.User, error)
	// Logout(ctx context.Context) error
	UpdateUserById(ctx context.Context, id uint64) (model.User, error)
	DeleteUserById(ctx context.Context, id uint64) error
}

// STRUCT
type userQueryImpl struct {
	db infrastructure.GormPostgres
}
type userCommandImpl struct {
    db infrastructure.GormPostgres
}

// NEW USER QUERY
func NewUserQuery(db infrastructure.GormPostgres) UserQuery {
	return &userQueryImpl{db:db}
}
func NewUserCommand(db infrastructure.GormPostgres) UserCommand {
    return &userCommandImpl{db: db}
}

// USER QUERY IMPL
func (u *userQueryImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	db := u.db.GetConnection()
	users := []model.User{}

	if err := db.WithContext(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}
func (u *userQueryImpl) GetUserById(ctx context.Context, id uint64) (model.User, error) {
	db := u.db.GetConnection()
	users := model.User{}

	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Find(&users).Error; err != nil {
		return model.User{}, nil
	}
	return users, nil
}

// USER COMMAND IMPL
func (u *userCommandImpl) Register(ctx context.Context, user model.User) (model.User, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (u *userCommandImpl) Login(ctx context.Context) (model.User, error) {
	// login
	user := model.User{}

	return user, nil
}
func (u *userCommandImpl) Logout(ctx context.Context) error {
	// logout
	return nil
}
func (u *userCommandImpl) UpdateUserById(ctx context.Context, id uint64) (model.User, error) {
	db := u.db.GetConnection()
	user := model.User{}

	if err := db.WithContext(ctx).Table("users").First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	if err := db.WithContext(ctx).Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (u *userCommandImpl) DeleteUserById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}