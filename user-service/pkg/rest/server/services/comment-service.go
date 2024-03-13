package services

import (
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
)

type CommentService struct {
	commentDao *daos.CommentDao
}

func NewCommentService() (*CommentService, error) {
	commentDao, err := daos.NewCommentDao()
	if err != nil {
		return nil, err
	}
	return &CommentService{
		commentDao: commentDao,
	}, nil
}

func (commentService *CommentService) CreateComment(comment *models.Comment) (*models.Comment, error) {
	return commentService.commentDao.CreateComment(comment)
}

func (commentService *CommentService) ListComments() ([]*models.Comment, error) {
	return commentService.commentDao.ListComments()
}

func (commentService *CommentService) GetComment(id int64) (*models.Comment, error) {
	return commentService.commentDao.GetComment(id)
}

func (commentService *CommentService) UpdateComment(id int64, comment *models.Comment) (*models.Comment, error) {
	return commentService.commentDao.UpdateComment(id, comment)
}

func (commentService *CommentService) DeleteComment(id int64) error {
	return commentService.commentDao.DeleteComment(id)
}
