package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	entity "ptxyz/customer-product-service/internal/domain/entity/auth"

	userRepsitory "ptxyz/customer-product-service/internal/domain/repository/user"
)

type AuthService interface {
	VerifyCredential(ctx context.Context, input entity.VerifyInput) (*entity.VerifyOutput, error)
}

type authServiceImpl struct {
	userRepo userRepsitory.UserRepository
}

func New(userRepo userRepsitory.UserRepository) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
	}
}

func (s *authServiceImpl) VerifyCredential(ctx context.Context, input entity.VerifyInput) (*entity.VerifyOutput, error) {

	user, err := s.userRepo.GetByNik(ctx, input.Nik)
	if err != nil {
		return nil, errors.New("salah email or password")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {
		return nil, errors.New("salah email or password")
	}

	return &entity.VerifyOutput{
		UserPublicId: user.PublicId,
	}, nil
}
