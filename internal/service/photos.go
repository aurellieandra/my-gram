package service

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
)

// INTERFACE
type PhotoService interface {
	GetPhotos(ctx context.Context) ([]model.Photo, error)
	GetPhotosByUserId(ctx context.Context, id uint64) ([]model.Photo, error)
	GetPhotoById(ctx context.Context, id uint64) (*model.Photo, error)

	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	UpdatePhotoById(ctx context.Context, photo model.Photo, id uint64) (model.Photo, error)
	DeletePhotoById(ctx context.Context, id uint64) error
}

// STRUCT
type photoServiceImpl struct {
	repo    repository.PhotoQuery
	command repository.PhotoCommand
}

// NEW USER SERVICE
func NewPhotoService(repo repository.PhotoQuery, command repository.PhotoCommand) PhotoService {
	return &photoServiceImpl{repo: repo, command: command}
}

// USER SERVICE IMPL
func (u *photoServiceImpl) GetPhotos(ctx context.Context) ([]model.Photo, error) {
	photos, err := u.repo.GetPhotos(ctx)
	if err != nil {
		return nil, err
	}
	return photos, err
}

func (u *photoServiceImpl) GetPhotosByUserId(ctx context.Context, id uint64) ([]model.Photo, error) {
	photos, err := u.repo.GetPhotosByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	return photos, err
}

func (u *photoServiceImpl) GetPhotoById(ctx context.Context, id uint64) (*model.Photo, error) {
	photo, err := u.repo.GetPhotoById(ctx, id)
	if err != nil {
		return nil, err
	}

	return photo, err
}

func (u *photoServiceImpl) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	createdPhoto, err := u.command.CreatePhoto(ctx, photo)
	if err != nil {
		return model.Photo{}, err
	}
	return createdPhoto, nil
}

func (u *photoServiceImpl) UpdatePhotoById(ctx context.Context, photo model.Photo, id uint64) (model.Photo, error) {
	updatedPhoto, err := u.command.UpdatePhotoById(ctx, photo, id)
	if err != nil {
		return model.Photo{}, err
	}
	return updatedPhoto, nil
}

func (u *photoServiceImpl) DeletePhotoById(ctx context.Context, id uint64) error {
	err := u.command.DeletePhotoById(ctx, id)
	return err
}
