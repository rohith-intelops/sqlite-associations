package daos

import (
	"errors"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos/clients/sqls"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FollowersDao struct {
	db *gorm.DB
}

func NewFollowersDao() (*FollowersDao, error) {
	sqlClient, err := sqls.InitGORMSQLiteDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Followers{})
	if err != nil {
		return nil, err
	}
	return &FollowersDao{
		db: sqlClient.DB,
	}, nil
}

func (followersDao *FollowersDao) CreateFollowers(m *models.Followers) (*models.Followers, error) {
	if err := followersDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create followers: %v", err)
		return nil, err
	}

	log.Debugf("followers created")
	return m, nil
}

func (followersDao *FollowersDao) ListFollowers() ([]*models.Followers, error) {
	var followers []*models.Followers

	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := followersDao.db.Find(&followers).Error; err != nil {
		log.Debugf("failed to list followers: %v", err)
		return nil, err
	}

	log.Debugf("followers listed")
	return followers, nil
}

func (followersDao *FollowersDao) GetFollowers(id int64) (*models.Followers, error) {
	var m *models.Followers
	if err := followersDao.db.Where("id = ?", id).Preload("Following").First(&m).Error; err != nil {
		log.Debugf("failed to get followers: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("followers retrieved")
	return m, nil
}

func (followersDao *FollowersDao) UpdateFollowers(id int64, m *models.Followers) (*models.Followers, error) {
	if id == 0 {
		return nil, errors.New("invalid followers ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var followers *models.Followers
	if err := followersDao.db.Where("id = ?", id).First(&followers).Error; err != nil {
		log.Debugf("failed to find followers for update: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	if err := followersDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update followers: %v", err)
		return nil, err
	}
	log.Debugf("followers updated")
	return m, nil
}

func (followersDao *FollowersDao) DeleteFollowers(id int64) error {
	var m *models.Followers
	if err := followersDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete followers: %v", err)
		return err
	}

	log.Debugf("followers deleted")
	return nil
}
