package service

import "github.com/QR-authentication/auth-service/internal/model"

type DBRepo interface {
	UserExists(cardNumber string) (bool, error)
	GetUserData(cardNumber string) (*model.User, error)
}
