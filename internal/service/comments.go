package service

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
)

// INTERFACE
type CommentService interface {
	GetCommentsByPhotoId(ctx context.Context, id uint64) ([]model.Comment, error)
	GetCommentById(ctx context.Context, id uint64) (model.Comment, error)
	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	UpdateCommentById(ctx context.Context, id uint64) (model.Comment, error)
	DeleteCommentById(ctx context.Context, id uint64) error
}

// STRUCT
type commentServiceImpl struct {
	repo repository.CommentQuery
	command repository.CommentCommand
}

// NEW USER SERVICE
func NewCommentService(repo repository.CommentQuery, command repository.CommentCommand) CommentService {
	return &commentServiceImpl{repo:repo, command:command}
}

// USER SERVICE IMPL
func (u *commentServiceImpl) GetCommentsByPhotoId(ctx context.Context, id uint64) ([]model.Comment, error) {
	comments, err := u.repo.GetCommentsByPhotoId(ctx, id)
	if err != nil {
		return nil, err
	}
	return comments, err
}
func (u *commentServiceImpl) GetCommentById(ctx context.Context, id uint64) (model.Comment, error) {
	comment, err := u.repo.GetCommentById(ctx, id)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, err
}
func (u *commentServiceImpl) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	photo, err := u.command.CreateComment(ctx, comment)
	if err != nil {
		return model.Comment{}, err
	}
	return photo, err
}
func (u *commentServiceImpl) UpdateCommentById(ctx context.Context, id uint64) (model.Comment, error) {
    updatedPhoto, err := u.command.UpdateCommentById(ctx, id)
    if err != nil {
        return model.Comment{}, err
    }
    return updatedPhoto, nil
}
func (u *commentServiceImpl) DeleteCommentById(ctx context.Context, id uint64) error {
    err := u.command.DeleteCommentById(ctx, id)
    return err
}