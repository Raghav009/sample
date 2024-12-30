package services

import (
	"database/sql"
	"sample/internal/db"
	"sample/internal/models"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetUserDetails(username string) (*models.User, error) {
	return db.GetUser(username, s.DB)
}
