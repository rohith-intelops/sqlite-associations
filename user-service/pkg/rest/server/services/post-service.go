package services

import (
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/daos"
	"github.com/rohith-intelops/socialmedia/user-service/pkg/rest/server/models"
)

type PostService struct {
	postDao *daos.PostDao
}

func NewPostService() (*PostService, error) {
	postDao, err := daos.NewPostDao()
	if err != nil {
		return nil, err
	}
	return &PostService{
		postDao: postDao,
	}, nil
}

func (postService *PostService) CreatePost(post *models.Post) (*models.Post, error) {
	return postService.postDao.CreatePost(post)
}

func (postService *PostService) ListPosts() ([]*models.Post, error) {
	return postService.postDao.ListPosts()
}

func (postService *PostService) GetPost(id int64) (*models.Post, error) {
	return postService.postDao.GetPost(id)
}

func (postService *PostService) UpdatePost(id int64, post *models.Post) (*models.Post, error) {
	return postService.postDao.UpdatePost(id, post)
}

func (postService *PostService) DeletePost(id int64) error {
	return postService.postDao.DeletePost(id)
}
