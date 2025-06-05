package repository

import (
	"database/sql"
	"skycompare-backend-main/internal/models"
)


type UserRepositoryInterface interface {
	GetUserByCredentials(nickname, password string) (*models.User, error)
	GetUserByNickname(nickname string) (*models.User, error)
	InsertUser(user models.User) error
	GetNextUserID() (int, error)
	UpdateFavourite(nickname, fav string) error
	GetFavourite(nickname string) (string, error)
}


type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUserByCredentials(nickname, password string) (*models.User, error) {
	query := "SELECT id, nickname, full_name, email, pass, fav_airport FROM users WHERE nickname=? AND pass=?"
	row := r.DB.QueryRow(query, nickname, password)

	var u models.User
	err := row.Scan(&u.ID, &u.Nickname, &u.FullName, &u.Email, &u.Password, &u.FavAirport)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByNickname(nickname string) (*models.User, error) {
	query := "SELECT id, nickname, full_name, email, pass, fav_airport FROM users WHERE nickname=?"
	row := r.DB.QueryRow(query, nickname)

	var u models.User
	err := row.Scan(&u.ID, &u.Nickname, &u.FullName, &u.Email, &u.Password, &u.FavAirport)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) InsertUser(user models.User) error {
	query := "INSERT INTO users (id, nickname, full_name, email, pass, fav_airport) VALUES(?,?,?,?,?,?)"
	_, err := r.DB.Exec(query, user.ID, user.Nickname, user.FullName, user.Email, user.Password, user.FavAirport)
	return err
}

func (r *UserRepository) GetNextUserID() (int, error) {
	var id int
	row := r.DB.QueryRow("SELECT COALESCE(MAX(id), 0) FROM users")
	err := row.Scan(&id)
	return id + 1, err
}

func (r *UserRepository) UpdateFavourite(nickname, fav string) error {
	query := "UPDATE users SET fav_airport=? WHERE nickname=?"
	_, err := r.DB.Exec(query, fav, nickname)
	return err
}

func (r *UserRepository) GetFavourite(nickname string) (string, error) {
	var fav string
	row := r.DB.QueryRow("SELECT fav_airport FROM users WHERE nickname=?", nickname)
	err := row.Scan(&fav)
	return fav, err
}
