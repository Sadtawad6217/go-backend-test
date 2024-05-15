package service

import (
	"gobackend/pkg/core/model"
	"gobackend/pkg/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) GetPostAll(limit, offset int, searchTitle string) ([]model.Posts, error) {
	return s.repository.GetPostAll(limit, offset, searchTitle)
}

func (s *Service) GetTotalPostCount(searchTitle string) (int, error) {
	return s.repository.GetTotalPostCount(searchTitle)
}

func (s *Service) GetPostID(id string) ([]model.Posts, error) {
	return s.repository.GetPostID(id)
}

func (s *Service) IncrementViewCount(id string) error {
	return s.repository.IncrementViewCount(id)
}

func (s *Service) CreatePosts(post model.Posts) (model.Posts, error) {
	return s.repository.CreatePosts(post)
}

func (s *Service) UpdatePost(id string, updateData model.Posts) (model.Posts, error) {
	return s.repository.UpdatePost(id, updateData)
}

func (s *Service) DeletePost(id string) error {
	return s.repository.DeletePost(id)
}
