package repository

import (
	"context"
	"time"

	"github.com/aurellieandra/my-gram/internal/infrastructure"
	"github.com/aurellieandra/my-gram/internal/model"
)

// INTERFACE
type CommentQuery interface {
	GetComments(ctx context.Context, photo_id *uint64) ([]model.Comment, error)
	GetCommentById(ctx context.Context, id uint64) (*model.Comment, error)
}
type CommentCommand interface {
	CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error)
	UpdateCommentById(ctx context.Context, updatedComment model.Comment, id uint64) (model.Comment, error)
	DeleteCommentById(ctx context.Context, id uint64) error
}

// STRUCT
type commentQueryImpl struct {
	db infrastructure.GormPostgres
}
type commentCommandImpl struct {
	db infrastructure.GormPostgres
}

// NEW SOCIAL MEDIA QUERY
func NewCommentQuery(db infrastructure.GormPostgres) CommentQuery {
	return &commentQueryImpl{db: db}
}
func NewCommentCommand(db infrastructure.GormPostgres) CommentCommand {
	return &commentCommandImpl{db: db}
}

// SOCIAL MEDIA QUERY IMPL
func (u *commentQueryImpl) GetComments(ctx context.Context, photo_id *uint64) ([]model.Comment, error) {
	db := u.db.GetConnection()
	comments := []model.Comment{}

	query := db.WithContext(ctx).Table("comments").Where("deleted_at IS NULL")

	if photo_id != nil && *photo_id != 0 {
		query = query.Where("photo_id = ?", *photo_id)
	}

	if err := query.Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (u *commentQueryImpl) GetCommentById(ctx context.Context, id uint64) (*model.Comment, error) {
	db := u.db.GetConnection()
	var comment *model.Comment

	if err := db.WithContext(ctx).Table("comments").Where("id = ?", id).Where("deleted_at IS NULL").Find(&comment).Error; err != nil {
		return nil, nil
	}

	return comment, nil
}

// PHOTO COMMAND IMPL
func (u *commentCommandImpl) CreateComment(ctx context.Context, comment model.Comment) (model.Comment, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Create(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (u *commentCommandImpl) UpdateCommentById(ctx context.Context, updatedComment model.Comment, id uint64) (model.Comment, error) {
	db := u.db.GetConnection()

	comment := model.Comment{}
	if err := db.WithContext(ctx).Table("comments").Where("id = ?", id).Where("deleted_at IS NULL").First(&comment).Error; err != nil {
		return model.Comment{}, err
	}

	comment.Message = updatedComment.Message

	if err := db.WithContext(ctx).Save(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

func (u *commentCommandImpl) DeleteCommentById(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()

	// HARD DELETE
	// if err := db.WithContext(ctx).Table("comments").Where("id = ?", id).Delete(&model.Comment{}).Error; err != nil {
	// 	return err
	// }

	// SOFT DELETE
	if err := db.WithContext(ctx).Table("comments").Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
