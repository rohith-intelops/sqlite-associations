package daos

import (
	"errors"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos/clients/sqls"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao() (*CommentDao, error) {
	sqlClient, err := sqls.InitGORMSQLiteDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Comment{})
	if err != nil {
		return nil, err
	}
	return &CommentDao{
		db: sqlClient.DB,
	}, nil
}

func (commentDao *CommentDao) CreateComment(m *models.Comment) (*models.Comment, error) {
	if err := commentDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create comment: %v", err)
		return nil, err
	}

	log.Debugf("comment created")
	return m, nil
}

func (commentDao *CommentDao) ListComments() ([]*models.Comment, error) {
	var comments []*models.Comment

	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := commentDao.db.Find(&comments).Error; err != nil {
		log.Debugf("failed to list comments: %v", err)
		return nil, err
	}

	log.Debugf("comment listed")
	return comments, nil
}

func (commentDao *CommentDao) GetComment(id int64) (*models.Comment, error) {
	var m *models.Comment
	if err := commentDao.db.Preload("Post").Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get comment: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("comment retrieved")
	return m, nil
}

func (commentDao *CommentDao) UpdateComment(id int64, m *models.Comment) (*models.Comment, error) {
	if id == 0 {
		return nil, errors.New("invalid comment ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var comment *models.Comment
	if err := commentDao.db.Where("id = ?", id).First(&comment).Error; err != nil {
		log.Debugf("failed to find comment for update: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	if err := commentDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update comment: %v", err)
		return nil, err
	}
	log.Debugf("comment updated")
	return m, nil
}

func (commentDao *CommentDao) DeleteComment(id int64) error {
	var m *models.Comment
	if err := commentDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete comment: %v", err)
		return err
	}

	log.Debugf("comment deleted")
	return nil
}
