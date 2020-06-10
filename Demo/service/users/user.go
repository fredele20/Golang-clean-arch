package users

import (
	"Demo/entity"
	resErr "Demo/errors"
	"Demo/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Service interface {
	Register(user *entity.User) (*entity.User, error)
	UserAlreadyExist(user *entity.User) (*entity.User, error)
	Validate(user *entity.User) error
	Login(user *entity.User) (*entity.User, error)
	ValidatePassword(providedPassword string, user *entity.User) (string, error)
	EncryptPassword(user *entity.User) (string, error)
}

type service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) Validate(user *entity.User) error {
	var emailRegexp = regexp.MustCompile("^[a-zA-z0-9.!#$%&'*+/?^_`{|}~]+@[a-zA-z0-9](?:[a-zA-z0-9]{0," +
		"61}[a-zA-z0-9])?(?:\\.[a-zA-z0-9](?:[a-zA-z0-9]{0,61}[a-zA-z0-9])?)*$")
	if user == nil {
		err := errors.New("form can not be empty")
		return  err
	}

	if user.Username == "" {
		err := errors.New("username can not be empty")
		return err
	}

	if user.Email == "" {
		err := errors.New("email can not be empty")
		return err
	}

	if !emailRegexp.MatchString(user.Email) {
		err := errors.New("email is not valid")
		return err
	}

	if user.Firstname == "" {
		err := errors.New("firstname can not be empty")
		return err
	}

	if user.Lastname == "" {
		err := errors.New("lastname can not be empty")
		return err
	}
	return nil
}

func (s *service) UserAlreadyExist(user *entity.User) (*entity.User, error) {
	//nameExists, _ := s.repo.FindByUsername(user)
	////if err != nil {
	////	return nil, err
	////}
	//if nameExists {
	//	return nil, resErr.ErrUserWithUsernameAlreadyExist
	//}

	emailExists, _ := s.repo.FindByEmail(user)
	//if err != nil {
	//	return nil, err
	//}
	if emailExists {
		return nil, resErr.ErrUserWithEmailAlreadyExists
	}
	return user, nil
}

func (s *service) Register(user *entity.User) (*entity.User, error) {

	return s.repo.Create(user)
}

func (s *service) EncryptPassword(user *entity.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 21)
	if err != nil {
		return "", nil
	}
	user.Password = string(hashedPassword)
	return string(hashedPassword), nil
}

func (s *service) ValidatePassword(providedPassword string, user *entity.User) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	//p, err := newFromH
	if err != nil {
		return "wrong password", err
	}
	return "correct", nil
}

func (s *service) Login(user *entity.User) (*entity.User, error) {
	_, err := s.repo.FindByEmail(user)
	if err != nil {
		err := errors.New("invalid email provided")
		return nil, err
	}
	return user, nil
}
