package repository

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type SocialMediaQuery interface {
	GetSocialMediasByUserId(ctx context.Context, id uint64) ([]model.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error)
}
type SocialMediaCommand interface {
	CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error)
	UpdateSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error)
	DeleteSocialMediaById(ctx context.Context, id uint64) error
}

// STRUCT
type socialMediaQueryImpl struct {
	db infrastructure.GormPostgres
}
type socialMediaCommandImpl struct {
    db infrastructure.GormPostgres
}

// NEW USER QUERY
func NewSocialMediaQuery(db infrastructure.GormPostgres) SocialMediaQuery {
	return &socialMediaQueryImpl{db:db}
}
func NewSocialMediaCommand(db infrastructure.GormPostgres) SocialMediaCommand {
    return &socialMediaCommandImpl{db: db}
}

// USER QUERY IMPL
func (u *socialMediaQueryImpl) GetSocialMediasByUserId(ctx context.Context, id uint64) ([]model.SocialMedia, error) {
	db := u.db.GetConnection()
	social_medias := []model.SocialMedia{}

	if err := db.WithContext(ctx).Table("social_medias").Where("user_id = ?", id).Find(&social_medias).Error; err != nil {
		return []model.SocialMedia{}, nil
	}
	return social_medias, nil
}
func (u *socialMediaQueryImpl) GetSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error) {
	db := u.db.GetConnection()
	social_medias := model.SocialMedia{}

	if err := db.WithContext(ctx).Table("social_medias").Where("id = ?", id).Find(&social_medias).Error; err != nil {
		return model.SocialMedia{}, nil
	}
	return social_medias, nil
}

// USER COMMAND IMPL
func (u *socialMediaCommandImpl) CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Create(&social_media).Error; err != nil {
		return model.SocialMedia{}, err
	}
	return social_media, nil
}
func (u *socialMediaCommandImpl) UpdateSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error) {
	db := u.db.GetConnection()
	social_media := model.SocialMedia{}

	if err := db.WithContext(ctx).Table("social_medias").First(&social_media, id).Error; err != nil {
		return model.SocialMedia{}, err
	}
	if err := db.WithContext(ctx).Save(&social_media).Error; err != nil {
		return model.SocialMedia{}, err
	}
	return social_media, nil
}
func (u *socialMediaCommandImpl) DeleteSocialMediaById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("social_medias").Where("id = ?", id).Delete(&model.SocialMedia{}).Error; err != nil {
		return err
	}
	return nil
}