package service

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
)

// INTERFACE
type SocialMediaService interface {
	GetSocialMediasByUserId(ctx context.Context, id uint64) ([]model.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error)
	CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error)
	UpdateSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error)
	DeleteSocialMediaById(ctx context.Context, id uint64) error
}

// STRUCT
type socialMediaServiceImpl struct {
	repo repository.SocialMediaQuery
	command repository.SocialMediaCommand
}

// NEW USER SERVICE
func NewSocialMediaService(repo repository.SocialMediaQuery, command repository.SocialMediaCommand) SocialMediaService {
	return &socialMediaServiceImpl{repo:repo, command:command}
}

// USER SERVICE IMPL
func (u *socialMediaServiceImpl) GetSocialMediasByUserId(ctx context.Context, id uint64) ([]model.SocialMedia, error) {
	photo, err := u.repo.GetSocialMediasByUserId(ctx, id)
	if err != nil {
		return []model.SocialMedia{}, err
	}
	return photo, err
}
func (u *socialMediaServiceImpl) GetSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error) {
	photo, err := u.repo.GetSocialMediaById(ctx, id)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return photo, err
}
func (u *socialMediaServiceImpl) CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error) {
	photo, err := u.command.CreateSocialMedia(ctx, social_media)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return photo, err
}
func (u *socialMediaServiceImpl) UpdateSocialMediaById(ctx context.Context, id uint64) (model.SocialMedia, error) {
    updatedPhoto, err := u.command.UpdateSocialMediaById(ctx, id)
    if err != nil {
        return model.SocialMedia{}, err
    }
    return updatedPhoto, nil
}
func (u *socialMediaServiceImpl) DeleteSocialMediaById(ctx context.Context, id uint64) error {
    err := u.command.DeleteSocialMediaById(ctx, id)
    return err
}