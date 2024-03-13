package services

import (
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
)

type FollowersService struct {
	followersDao *daos.FollowersDao
}

func NewFollowersService() (*FollowersService, error) {
	followersDao, err := daos.NewFollowersDao()
	if err != nil {
		return nil, err
	}
	return &FollowersService{
		followersDao: followersDao,
	}, nil
}

func (followersService *FollowersService) CreateFollowers(followers *models.Followers) (*models.Followers, error) {
	return followersService.followersDao.CreateFollowers(followers)
}

func (followersService *FollowersService) ListFollowers() ([]*models.Followers, error) {
	return followersService.followersDao.ListFollowers()
}

func (followersService *FollowersService) GetFollowers(id int64) (*models.Followers, error) {
	return followersService.followersDao.GetFollowers(id)
}

func (followersService *FollowersService) UpdateFollowers(id int64, followers *models.Followers) (*models.Followers, error) {
	return followersService.followersDao.UpdateFollowers(id, followers)
}

func (followersService *FollowersService) DeleteFollowers(id int64) error {
	return followersService.followersDao.DeleteFollowers(id)
}
