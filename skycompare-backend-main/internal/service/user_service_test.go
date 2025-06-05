package service

import (
	"errors"
	"skycompare-backend-main/internal/models"
	"testing"
)

type MockUserRepo struct {
	GetUserByCredentialsFunc func(nickname, password string) (*models.User, error)
	GetUserByNicknameFunc    func(nickname string) (*models.User, error)
	InsertUserFunc           func(user models.User) error
	GetNextUserIDFunc        func() (int, error)
	UpdateFavouriteFunc      func(nickname, fav string) error
	GetFavouriteFunc         func(nickname string) (string, error)
}

func (m *MockUserRepo) GetUserByCredentials(nickname, password string) (*models.User, error) {
	return m.GetUserByCredentialsFunc(nickname, password)
}
func (m *MockUserRepo) GetUserByNickname(nickname string) (*models.User, error) {
	return m.GetUserByNicknameFunc(nickname)
}
func (m *MockUserRepo) InsertUser(user models.User) error {
	return m.InsertUserFunc(user)
}
func (m *MockUserRepo) GetNextUserID() (int, error) {
	return m.GetNextUserIDFunc()
}
func (m *MockUserRepo) UpdateFavourite(nickname, fav string) error {
	return m.UpdateFavouriteFunc(nickname, fav)
}
func (m *MockUserRepo) GetFavourite(nickname string) (string, error) {
	return m.GetFavouriteFunc(nickname)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := &MockUserRepo{
		GetUserByCredentialsFunc: func(nickname, password string) (*models.User, error) {
			return &models.User{ID: 1, Nickname: "testuser"}, nil
		},
	}

	service := UserService{Repo: mockRepo}
	ok, err := service.Login("testuser", "password")

	if err != nil || !ok {
		t.Errorf("Expected login to succeed, got err: %v", err)
	}
}

func TestLogin_Fail(t *testing.T) {
	mockRepo := &MockUserRepo{
		GetUserByCredentialsFunc: func(nickname, password string) (*models.User, error) {
			return nil, errors.New("user not found")
		},
	}

	service := UserService{Repo: mockRepo}
	ok, err := service.Login("wronguser", "wrongpass")

	if err == nil || ok {
		t.Errorf("Expected login to fail, got success")
	}
}
