package service

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
)

// INTERFACE
type SocialMediaService interface {
	GetSocialMedias(ctx context.Context, user_id *uint64) ([]model.SocialMedia, error)
	GetSocialMediaById(ctx context.Context, id uint64) (*model.SocialMedia, error)

	CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error)
	UpdateSocialMediaById(ctx context.Context, social_media model.SocialMedia, id uint64) (model.SocialMedia, error)
	DeleteSocialMediaById(ctx context.Context, id uint64) error
}

// STRUCT
type socialMediaServiceImpl struct {
	repo    repository.SocialMediaQuery
	command repository.SocialMediaCommand
}

// NEW USER SERVICE
func NewSocialMediaService(repo repository.SocialMediaQuery, command repository.SocialMediaCommand) SocialMediaService {
	return &socialMediaServiceImpl{repo: repo, command: command}
}

// USER SERVICE IMPL
func (u *socialMediaServiceImpl) GetSocialMedias(ctx context.Context, user_id *uint64) ([]model.SocialMedia, error) {
	id := *user_id

	social_medias, err := u.repo.GetSocialMedias(ctx, &id)
	if err != nil {
		return nil, err
	}
	return social_medias, err
}

func (u *socialMediaServiceImpl) GetSocialMediaById(ctx context.Context, id uint64) (*model.SocialMedia, error) {
	social_media, err := u.repo.GetSocialMediaById(ctx, id)
	if err != nil {
		return nil, err
	}

	return social_media, err
}

func (u *socialMediaServiceImpl) CreateSocialMedia(ctx context.Context, social_media model.SocialMedia) (model.SocialMedia, error) {
	createdSocialMedia, err := u.command.CreateSocialMedia(ctx, social_media)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return createdSocialMedia, nil
}

func (u *socialMediaServiceImpl) UpdateSocialMediaById(ctx context.Context, social_media model.SocialMedia, id uint64) (model.SocialMedia, error) {
	updatedSocialMedia, err := u.command.UpdateSocialMediaById(ctx, social_media, id)
	if err != nil {
		return model.SocialMedia{}, err
	}
	return updatedSocialMedia, nil
}

func (u *socialMediaServiceImpl) DeleteSocialMediaById(ctx context.Context, id uint64) error {
	err := u.command.DeleteSocialMediaById(ctx, id)
	return err
}
