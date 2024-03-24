package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// INTERFACE
type UserQuery interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id uint64) (*model.User, error)
}
type UserCommand interface {
	Register(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, user model.User) (model.User, error)

	UpdateUserById(ctx context.Context, updatedUser model.User, id uint64) (model.User, error)
	DeleteUserById(ctx context.Context, id uint64) error

	// MISC
	GenerateUserAccessToken(ctx context.Context, user model.User) (model.User, error)
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

func (u *userQueryImpl) GetUserById(ctx context.Context, id uint64) (*model.User, error) {
	db := u.db.GetConnection()
	var users *model.User

	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Find(&users).Error; err != nil {
		return nil, nil
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

func (u *userCommandImpl) Login(ctx context.Context, credentials model.User) (model.User, error) {
	db := u.db.GetConnection()

	var existingUser model.User
	err := db.WithContext(ctx).Table("users").Where("username = ?", credentials.Username).First(&existingUser).Error
	if err != nil {
		return model.User{}, err
	}

	fmt.Println(existingUser.Password)
	fmt.Println(credentials.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(credentials.Password)); err != nil {
		return model.User{}, err
	}

	return existingUser, nil
}

func (u *userCommandImpl) UpdateUserById(ctx context.Context, updatedUser model.User, id uint64) (model.User, error) {
	db := u.db.GetConnection()

	var user model.User
	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, err
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Dob = updatedUser.Dob

	if err := db.WithContext(ctx).Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userCommandImpl) DeleteUserById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	// HARD DELETE
	// if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
	// 	return err
	// }

	// SOFT DELETE
	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

// MISC
func (u *userCommandImpl) GenerateUserAccessToken(ctx context.Context, user model.User) (model.User, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Preload("users").Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
