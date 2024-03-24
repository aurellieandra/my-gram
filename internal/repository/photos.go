package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type PhotoQuery interface {
	GetPhotos(ctx context.Context, user_id *uint64) ([]model.Photo, error)
	GetPhotoById(ctx context.Context, id uint64) (*model.Photo, error)
}
type PhotoCommand interface {
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	UpdatePhotoById(ctx context.Context, updatedPhoto model.Photo, id uint64) (model.Photo, error)
	DeletePhotoById(ctx context.Context, id uint64) error
}

// STRUCT
type photoQueryImpl struct {
	db infrastructure.GormPostgres
}
type photoCommandImpl struct {
	db infrastructure.GormPostgres
}

// NEW PHOTO QUERY
func NewPhotoQuery(db infrastructure.GormPostgres) PhotoQuery {
	return &photoQueryImpl{db: db}
}
func NewPhotoCommand(db infrastructure.GormPostgres) PhotoCommand {
	return &photoCommandImpl{db: db}
}

// PHOTO QUERY IMPL
func (u *photoQueryImpl) GetPhotos(ctx context.Context, user_id *uint64) ([]model.Photo, error) {
	db := u.db.GetConnection()
	photos := []model.Photo{}

	query := db.WithContext(ctx).Table("photos").Where("deleted_at IS NULL")
	fmt.Println(query)

	if user_id != nil && *user_id != 0 {
		query = query.Where("user_id = ?", *user_id)
	}

	if err := query.Find(&photos).Error; err != nil {
		return nil, err
	}

	return photos, nil
}

func (u *photoQueryImpl) GetPhotoById(ctx context.Context, id uint64) (*model.Photo, error) {
	db := u.db.GetConnection()
	var photo *model.Photo

	if err := db.WithContext(ctx).Table("photos").Where("id = ?", id).Where("deleted_at IS NULL").Find(&photo).Error; err != nil {
		return nil, nil
	}

	return photo, nil
}

// PHOTO COMMAND IMPL
func (u *photoCommandImpl) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Create(&photo).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (u *photoCommandImpl) UpdatePhotoById(ctx context.Context, updatedPhoto model.Photo, id uint64) (model.Photo, error) {
	db := u.db.GetConnection()

	photo := model.Photo{}
	if err := db.WithContext(ctx).Table("photos").Where("id = ?", id).Where("deleted_at IS NULL").First(&photo).Error; err != nil {
		return model.Photo{}, err
	}

	photo.Title = updatedPhoto.Title
	photo.Url = updatedPhoto.Url
	photo.Caption = updatedPhoto.Caption

	if err := db.WithContext(ctx).Save(&photo).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}

func (u *photoCommandImpl) DeletePhotoById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	// HARD DELETE
	// if err := db.WithContext(ctx).Table("photos").Where("id = ?", id).Delete(&model.Photo{}).Error; err != nil {
	// 	return err
	// }

	// SOFT DELETE
	if err := db.WithContext(ctx).Table("photos").Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
