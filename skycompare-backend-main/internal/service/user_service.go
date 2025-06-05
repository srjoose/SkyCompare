package service

import (
	"skycompare-backend-main/internal/models"
	"skycompare-backend-main/internal/repository"
)

type UserService struct {
	Repo repository.UserRepositoryInterface
}

func (s *UserService) Login(nickname, password string) (bool, error) {
	user, err := s.Repo.GetUserByCredentials(nickname, password)
	if err != nil || user == nil {
		return false, err
	}
	return true, nil
}

func (s *UserService) Register(user models.User) (bool, error) {
	existing, err := s.Repo.GetUserByNickname(user.Nickname)
	if err == nil && existing != nil {
		return false, nil // ya existe el usuario
	}

	userID, err := s.Repo.GetNextUserID()
	if err != nil {
		return false, err
	}
	user.ID = userID

	err = s.Repo.InsertUser(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *UserService) UpdateFavourite(nickname, fav string) error {
	return s.Repo.UpdateFavourite(nickname, fav)
}

func (s *UserService) GetFavourite(nickname string) (string, error) {
	return s.Repo.GetFavourite(nickname)
}
