package daos

import (
	"errors"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos/clients/sqls"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FollowingDao struct {
	db *gorm.DB
}

func NewFollowingDao() (*FollowingDao, error) {
	sqlClient, err := sqls.InitGORMSQLiteDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Following{})
	if err != nil {
		return nil, err
	}
	return &FollowingDao{
		db: sqlClient.DB,
	}, nil
}

func (followingDao *FollowingDao) CreateFollowing(m *models.Following) (*models.Following, error) {
	if err := followingDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create following: %v", err)
		return nil, err
	}

	log.Debugf("following created")
	return m, nil
}

func (followingDao *FollowingDao) ListFollowings() ([]*models.Following, error) {
	var followings []*models.Following

	// TODO populate associations here with your own logic - https://gorm.io/docs/belongs_to.html
	if err := followingDao.db.Find(&followings).Error; err != nil {
		log.Debugf("failed to list followings: %v", err)
		return nil, err
	}

	log.Debugf("following listed")
	return followings, nil
}

func (followingDao *FollowingDao) GetFollowing(id int64) (*models.Following, error) {
	var m *models.Following
	if err := followingDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get following: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("following retrieved")
	return m, nil
}

func (followingDao *FollowingDao) UpdateFollowing(id int64, m *models.Following) (*models.Following, error) {
	if id == 0 {
		return nil, errors.New("invalid following ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var following *models.Following
	if err := followingDao.db.Where("id = ?", id).First(&following).Error; err != nil {
		log.Debugf("failed to find following for update: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	if err := followingDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update following: %v", err)
		return nil, err
	}
	log.Debugf("following updated")
	return m, nil
}

func (followingDao *FollowingDao) DeleteFollowing(id int64) error {
	var m *models.Following
	if err := followingDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete following: %v", err)
		return err
	}

	log.Debugf("following deleted")
	return nil
}
