package auth

import (
	"context"

	authEntity "ptxyz/main-service/internal/domain/entity/auth"
	authRepository "ptxyz/main-service/internal/domain/repository/auth"
)

type LoginService interface {
	Login(ctx context.Context, input authEntity.LoginInput) (*authEntity.LoginOutput, error)
}

type loginServiceImpl struct {
	repo  authRepository.AuthRepository
	token *TokenService
}

func NewLoginService(repo authRepository.AuthRepository, token *TokenService) LoginService {
	return &loginServiceImpl{
		repo:  repo,
		token: token,
	}
}

func (s *loginServiceImpl) Login(ctx context.Context, input authEntity.LoginInput) (*authEntity.LoginOutput, error) {
	identity, err := s.repo.VerifyCredential(
		ctx,
		input.Nik,
		input.Password,
	)
	if err != nil {
		return nil, err
	}

	// generate JWT token
	jwt, err := s.token.Generate(
		identity.UserPublicID,
	)
	if err != nil {
		return nil, err
	}

	return &authEntity.LoginOutput{
		AccessToken: jwt,
	}, nil
}
