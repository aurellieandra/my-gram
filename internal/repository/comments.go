package repository

import (
	"context"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type CommentQuery interface {
	GetCommentsByPhotoId(ctx context.Context, id uint64) ([]model.Comment, error)
	GetCommentById(ctx context.Context, id uint64) (model.Comment, error)
}
type CommentCommand interface {
	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	UpdateCommentById(ctx context.Context, id uint64) (model.Comment, error)
	DeleteCommentById(ctx context.Context, id uint64) error
}

// STRUCT
type commentQueryImpl struct {
	db infrastructure.GormPostgres
}
type commentComandImpl struct {
    db infrastructure.GormPostgres
}

// NEW USER QUERY
func NewCommentQuery(db infrastructure.GormPostgres) CommentQuery {
	return &commentQueryImpl{db:db}
}
func NewCommentCommand(db infrastructure.GormPostgres) CommentCommand {
    return &commentComandImpl{db: db}
}

// USER QUERY IMPL
func (u *commentQueryImpl) GetCommentsByPhotoId(ctx context.Context, id uint64) ([]model.Comment, error) {
	db := u.db.GetConnection()
	comments := []model.Comment{}

	if err := db.WithContext(ctx).Table("comments").Where("user_id = ?", id).Find(&comments).Error; err != nil {
		return []model.Comment{}, nil
	}
	return comments, nil
}
func (u *commentQueryImpl) GetCommentById(ctx context.Context, id uint64) (model.Comment, error) {
	db := u.db.GetConnection()
	comments := model.Comment{}

	if err := db.WithContext(ctx).Table("comments").Where("comment_id = ?", id).Find(&comments).Error; err != nil {
		return model.Comment{}, nil
	}
	return comments, nil
}

// USER COMMAND IMPL
func (u *commentComandImpl) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Create(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}
func (u *commentComandImpl) UpdateCommentById(ctx context.Context, id uint64) (model.Comment, error) {
	db := u.db.GetConnection()
	comment := model.Comment{}

	if err := db.WithContext(ctx).Table("comments").First(&comment, id).Error; err != nil {
		return model.Comment{}, err
	}
	if err := db.WithContext(ctx).Save(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}
func (u *commentComandImpl) DeleteCommentById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("comments").Where("id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}