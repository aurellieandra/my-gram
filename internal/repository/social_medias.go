package repository

import (
	"context"
	"time"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type SocialMediaQuery interface {
	GetSocialMedias(ctx context.Context, user_id *uint64) ([]model.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id uint64) (*model.SocialMedia, error)
}
type SocialMediaCommand interface {
	CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error)
	UpdateSocialMediaById(ctx context.Context, updatedSocialMedia model.SocialMedia, id uint64) (model.SocialMedia, error)
	DeleteSocialMediaById(ctx context.Context, id uint64) error
}

// STRUCT
type socialMediaQueryImpl struct {
	db infrastructure.GormPostgres
}
type socialMediaCommandImpl struct {
	db infrastructure.GormPostgres
}

// NEW SOCIAL MEDIA QUERY
func NewSocialMediaQuery(db infrastructure.GormPostgres) SocialMediaQuery {
	return &socialMediaQueryImpl{db: db}
}
func NewSocialMediaCommand(db infrastructure.GormPostgres) SocialMediaCommand {
	return &socialMediaCommandImpl{db: db}
}

// SOCIAL MEDIA QUERY IMPL
func (u *socialMediaQueryImpl) GetSocialMedias(ctx context.Context, user_id *uint64) ([]model.SocialMedia, error) {
	db := u.db.GetConnection()
	social_medias := []model.SocialMedia{}

	query := db.WithContext(ctx).Table("social_medias").Where("deleted_at IS NULL")

	if user_id != nil && *user_id != 0 {
		query = query.Where("user_id = ?", *user_id)
	}

	if err := query.Find(&social_medias).Error; err != nil {
		return nil, err
	}

	return social_medias, nil
}

func (u *socialMediaQueryImpl) GetSocialMediaById(ctx context.Context, id uint64) (*model.SocialMedia, error) {
	db := u.db.GetConnection()
	var social_media *model.SocialMedia

	if err := db.WithContext(ctx).Table("social_medias").Where("id = ?", id).Where("deleted_at IS NULL").Find(&social_media).Error; err != nil {
		return nil, nil
	}

	return social_media, nil
}

// PHOTO COMMAND IMPL
func (u *socialMediaCommandImpl) CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("social_medias").Create(&social_media).Error; err != nil {
		return model.SocialMedia{}, err
	}
	return social_media, nil
}

func (u *socialMediaCommandImpl) UpdateSocialMediaById(ctx context.Context, updatedSocialMedia model.SocialMedia, id uint64) (model.SocialMedia, error) {
	db := u.db.GetConnection()

	social_media := model.SocialMedia{}
	if err := db.WithContext(ctx).Table("social_medias").Where("id = ?", id).Where("deleted_at IS NULL").First(&social_media).Error; err != nil {
		return model.SocialMedia{}, err
	}

	social_media.Name = updatedSocialMedia.Name
	social_media.Url = updatedSocialMedia.Url

	if err := db.WithContext(ctx).Save(&social_media).Error; err != nil {
		return model.SocialMedia{}, err
	}
	return social_media, nil
}

func (u *socialMediaCommandImpl) DeleteSocialMediaById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	// HARD DELETE
	// if err := db.WithContext(ctx).Table("social_medias").Where("id = ?", id).Delete(&model.SocialMedia{}).Error; err != nil {
	// 	return err
	// }

	// SOFT DELETE
	if err := db.WithContext(ctx).Table("social_medias").Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
