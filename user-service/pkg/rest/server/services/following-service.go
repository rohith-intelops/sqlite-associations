package services

import (
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
)

type FollowingService struct {
	followingDao *daos.FollowingDao
}

func NewFollowingService() (*FollowingService, error) {
	followingDao, err := daos.NewFollowingDao()
	if err != nil {
		return nil, err
	}
	return &FollowingService{
		followingDao: followingDao,
	}, nil
}

func (followingService *FollowingService) CreateFollowing(following *models.Following) (*models.Following, error) {
	return followingService.followingDao.CreateFollowing(following)
}

func (followingService *FollowingService) ListFollowings() ([]*models.Following, error) {
	return followingService.followingDao.ListFollowings()
}

func (followingService *FollowingService) GetFollowing(id int64) (*models.Following, error) {
	return followingService.followingDao.GetFollowing(id)
}

func (followingService *FollowingService) UpdateFollowing(id int64, following *models.Following) (*models.Following, error) {
	return followingService.followingDao.UpdateFollowing(id, following)
}

func (followingService *FollowingService) DeleteFollowing(id int64) error {
	return followingService.followingDao.DeleteFollowing(id)
}
