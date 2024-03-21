package repository

import (
	"context"
	"errors"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// INTERFACE
type UserQuery interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id uint64) (model.User, error)
}
type UserCommand interface {
	Register(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, user model.User) (model.User, error)
	UpdateUserById(ctx context.Context, user model.User, id uint64) (model.User, error)
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
	return &userQueryImpl{db: db}
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
func (u *userCommandImpl) Login(ctx context.Context, user model.User) (model.User, error) {
	db := u.db.GetConnection()

	var foundUser model.User
	err := db.WithContext(ctx).Where("username = ?", user.Username).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, errors.New("Invalid username or password")
		}
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return model.User{}, errors.New("Invalid username or password")
	}

	return foundUser, nil
}
func (u *userCommandImpl) UpdateUserById(ctx context.Context, user model.User, id uint64) (model.User, error) {
	db := u.db.GetConnection()

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
