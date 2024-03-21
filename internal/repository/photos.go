package repository

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type PhotoQuery interface {
	GetPhotos(ctx context.Context) ([]model.Photo, error)
	GetPhotosByUserId(ctx context.Context, id uint64) ([]model.Photo, error)
	GetPhotoById(ctx context.Context, id uint64) (model.Photo, error)
}
type PhotoCommand interface {
	UpdatePhotoById(ctx context.Context, id uint64) (model.Photo, error)
	DeletePhotoById(ctx context.Context, id uint64) error
}

// STRUCT
type photoQueryImpl struct {
	db infrastructure.GormPostgres
}
type photoCommandImpl struct {
    db infrastructure.GormPostgres
}

// NEW USER QUERY
func NewPhotoQuery(db infrastructure.GormPostgres) PhotoQuery {
	return &photoQueryImpl{db:db}
}
func NewPhotoCommand(db infrastructure.GormPostgres) PhotoCommand {
    return &photoCommandImpl{db: db}
}

// USER QUERY IMPL
func (u *photoQueryImpl) GetPhotos(ctx context.Context) ([]model.Photo, error) {
	db := u.db.GetConnection()
	photos := []model.Photo{}

	if err := db.WithContext(ctx).Table("photos").Find(&photos).Error; err != nil {
		return nil, nil
	}
	return photos, nil
}
func (u *photoQueryImpl) GetPhotosByUserId(ctx context.Context, id uint64) ([]model.Photo, error) {
	db := u.db.GetConnection()
	photos := []model.Photo{}

	if err := db.WithContext(ctx).Table("photos").Where("user_id = ?", id).Find(&photos).Error; err != nil {
		return []model.Photo{}, nil
	}
	return photos, nil
}
func (u *photoQueryImpl) GetPhotoById(ctx context.Context, id uint64) (model.Photo, error) {
	db := u.db.GetConnection()
	photos := model.Photo{}

	if err := db.WithContext(ctx).Table("photos").Where("photo_id = ?", id).Find(&photos).Error; err != nil {
		return model.Photo{}, nil
	}
	return photos, nil
}

// USER COMMAND IMPL
func (u *photoCommandImpl) UpdatePhotoById(ctx context.Context, id uint64) (model.Photo, error) {
	db := u.db.GetConnection()
	photo := model.Photo{}

	if err := db.WithContext(ctx).Table("photos").First(&photo, id).Error; err != nil {
		return model.Photo{}, err
	}
	if err := db.WithContext(ctx).Save(&photo).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}
func (u *photoCommandImpl) DeletePhotoById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("photos").Where("id = ?", id).Delete(&model.Photo{}).Error; err != nil {
		return err
	}
	return nil
}