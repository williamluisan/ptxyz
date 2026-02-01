package user

import (
	"context"
	"errors"
	"fmt"

	entity "ptxyz/customer-product-service/internal/domain/entity/user"
	repository "ptxyz/customer-product-service/internal/domain/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, u *entity.UserInput) error
	GetByNik(ctx context.Context, id string) (*entity.User, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func New(r repository.UserRepository) UserService {
	if r == nil {
		panic("user repository cannot be nil")
	}
	return &userServiceImpl{
		repo: r,
	}
}

func (s *userServiceImpl) GetByNik(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.repo.GetByNik(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) Create(ctx context.Context, input *entity.UserInput) error {
	// check duplicate email
	existingUser, _ := s.repo.GetByNik(ctx, input.Nik)
	if (existingUser != nil) {
		return errors.New("the nik is already in use")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if (err != nil) {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// create the domain entity
	data := &entity.UserInput{
		Nik: input.Nik,
		Password: string(hashedPassword),
		FullName: input.FullName,
		LegalName: input.LegalName,
		TempatLahir: input.TempatLahir,
		TanggalLahir: input.TanggalLahir,
		Gaji: input.Gaji,
		FotoKtp: input.FotoKtp,
		FotoSelfie: input.FotoSelfie,
	}
	newUser := entity.NewUser(data)
	fmt.Println(newUser)

	if err := s.repo.Create(ctx, newUser); err != nil {
		return err
	}

	return nil
}