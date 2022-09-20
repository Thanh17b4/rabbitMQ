package service

import (
	model "github.com/Thanh17b4/practice/model"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	ListUser(page int64, limit int64) ([]*model.User, error)
	DetailUser(userID int64) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	Delete(userID int64) (int64, error)
	Create(user *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CountUsers() (int64, error)
}
type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}
func generatePassword(password string) ([]byte, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPass, nil
}

func (s UserService) CreateUser(user *model.User) (*model.User, error) {
	if user.Username == "" || user.Name == "" || user.Email == "" || user.Password == "" || user.Address == "" {
		return nil, errors.New("required field can not empty")
	}
	_, err1 := s.userRepo.GetUserByEmail(user.Email)
	if err1 == nil {
		return nil, errors.New("Email had been used")
	}
	_, err2 := s.userRepo.GetUserByUsername(user.Username)
	if err2 == nil {
		return nil, errors.New("Username had been used")
	}

	hashedPassword, err3 := generatePassword(user.Password)
	if err3 != nil {
		return nil, err3
	}
	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}

func (s UserService) GetListUser(page int64, limit int64) ([]*model.User, error) {
	if page == 0 || limit == 0 {
		return nil, errors.New("page and perPage must be bigger than 0")
	}
	u, err := s.userRepo.CountUsers()
	if err != nil {
		return nil, errors.New("Could not count users")
	}
	p := u % limit
	if p == 0 {
		pages := u / limit
		if page > pages {
			return nil, errors.New("Number of pages is too large, page is not exist")
		}
	}
	if p != 0 {
		pages := u/limit + 1
		if page > pages {
			return nil, errors.New("Number of pages is too large, page is not exist")
		}
	}

	return s.userRepo.ListUser(page, limit)
}

func (s UserService) GetDetailUser(userID int64) (*model.User, error) {
	return s.userRepo.DetailUser(userID)
}

func (s UserService) UpdateUserService(user *model.User) (*model.User, error) {
	if user.Username == "" || user.Name == "" || user.Email == "" || user.Password == "" || user.Address == "" {
		return nil, errors.New("required field can not empty")
	}

	_, err := s.GetDetailUser(int64(user.ID))
	if err != nil {
		return nil, err
	}
	hashedPassword, err := generatePassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.UpdateUser(user)
}
func (s UserService) DeleteUser(userID int64) (int64, error) {
	_, err := s.GetDetailUser(userID)
	if err != nil {
		return 0, err
	}
	return s.userRepo.Delete(userID)
}
