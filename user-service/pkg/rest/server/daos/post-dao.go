package daos

import (
	"errors"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos/clients/sqls"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostDao struct {
	db *gorm.DB
}

func NewPostDao() (*PostDao, error) {
	sqlClient, err := sqls.InitGORMSQLiteDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Post{})
	if err != nil {
		return nil, err
	}
	return &PostDao{
		db: sqlClient.DB,
	}, nil
}

func (postDao *PostDao) CreatePost(m *models.Post) (*models.Post, error) {
	if err := postDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create post: %v", err)
		return nil, err
	}

	log.Debugf("post created")
	return m, nil
}

func (postDao *PostDao) ListPosts() ([]*models.Post, error) {
	var posts []*models.Post

	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := postDao.db.Find(&posts).Error; err != nil {
		log.Debugf("failed to list posts: %v", err)
		return nil, err
	}

	log.Debugf("post listed")
	return posts, nil
}

func (postDao *PostDao) GetPost(id int64) (*models.Post, error) {
	var m *models.Post
	if err := postDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get post: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("post retrieved")
	return m, nil
}

func (postDao *PostDao) UpdatePost(id int64, m *models.Post) (*models.Post, error) {
	if id == 0 {
		return nil, errors.New("invalid post ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var post *models.Post
	if err := postDao.db.Where("id = ?", id).First(&post).Error; err != nil {
		log.Debugf("failed to find post for update: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	if err := postDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update post: %v", err)
		return nil, err
	}
	log.Debugf("post updated")
	return m, nil
}

func (postDao *PostDao) DeletePost(id int64) error {
	var m *models.Post
	if err := postDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete post: %v", err)
		return err
	}

	log.Debugf("post deleted")
	return nil
}
