package service

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/model"
	"github.com/aurellieandra/my-gram/internal/repository"
)

// INTERFACE
type CommentService interface {
	GetComments(ctx context.Context, photo_id *uint64) ([]model.Comment, error)
	GetCommentById(ctx context.Context, id uint64) (*model.Comment, error)

	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	UpdateCommentById(ctx context.Context, comment model.Comment, id uint64) (model.Comment, error)
	DeleteCommentById(ctx context.Context, id uint64) error
}

// STRUCT
type commentServiceImpl struct {
	repo    repository.CommentQuery
	command repository.CommentCommand
}

// NEW USER SERVICE
func NewCommentService(repo repository.CommentQuery, command repository.CommentCommand) CommentService {
	return &commentServiceImpl{repo: repo, command: command}
}

// USER SERVICE IMPL
func (u *commentServiceImpl) GetComments(ctx context.Context, photo_id *uint64) ([]model.Comment, error) {
	id := *photo_id

	comments, err := u.repo.GetComments(ctx, &id)
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (u *commentServiceImpl) GetCommentById(ctx context.Context, id uint64) (*model.Comment, error) {
	comment, err := u.repo.GetCommentById(ctx, id)
	if err != nil {
		return nil, err
	}

	return comment, err
}

func (u *commentServiceImpl) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	createdComment, err := u.command.CreateComment(ctx, comment)
	if err != nil {
		return model.Comment{}, err
	}
	return createdComment, nil
}

func (u *commentServiceImpl) UpdateCommentById(ctx context.Context, comment model.Comment, id uint64) (model.Comment, error) {
	updatedComment, err := u.command.UpdateCommentById(ctx, comment, id)
	if err != nil {
		return model.Comment{}, err
	}
	return updatedComment, nil
}

func (u *commentServiceImpl) DeleteCommentById(ctx context.Context, id uint64) error {
	err := u.command.DeleteCommentById(ctx, id)
	return err
}
