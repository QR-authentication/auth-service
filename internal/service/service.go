package service

import (
	"context"
	"time"

	authproto "github.com/QR-authentication/auth-proto/auth-proto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	authproto.UnimplementedAuthServiceServer
	repository     DBRepo
	authSigningKey string
}

func New(DBRepo DBRepo, authSigningKey string) *Service {
	return &Service{
		repository:     DBRepo,
		authSigningKey: authSigningKey,
	}
}

func (s *Service) Login(_ context.Context, in *authproto.LoginIn) (*authproto.LoginOut, error) {
	exists, err := s.repository.UserExists(in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user existence: %v", err)
	}

	if !exists {
		return &authproto.LoginOut{
			Token:       "",
			LoginStatus: false,
		}, nil
	}

	user, err := s.repository.GetUserData(in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user data: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.Password)); err != nil {
		return &authproto.LoginOut{
			Token:       "",
			LoginStatus: false,
		}, nil
	}

	claims := &jwt.RegisteredClaims{
		Subject:   user.UUID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.authSigningKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &authproto.LoginOut{
		Token:       tokenString,
		LoginStatus: true,
	}, nil
}
