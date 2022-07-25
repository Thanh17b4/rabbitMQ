package service

import (
	"golang.org/x/crypto/bcrypt"
	//goMail "gopkg.in/mail.v2"
	//"math/rand"
	model "github.com/Thanh17b4/practice/model"
	//"time"
)

type UserRepo interface {
	ListUser(page int64, limit int64) ([]*model.User, error)
	DetailUser(userID int64) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	Delete(userID int64) (int64, error)
	Creat(user *model.User) (*model.User, error)
	//CreatOTP(otp *model.UserOTP) (*model.UserOTP, error)
	GetUserByEmail(email string) (*model.User, error)
}
type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}
func (s UserService) GetListUser(page int64, limit int64) ([]*model.User, error) {
	return s.userRepo.ListUser(page, limit)
}

func (s UserService) GetDetailUser(userID int64) (*model.User, error) {
	return s.userRepo.DetailUser(userID)

}
func (s UserService) UpdateUserService(user *model.User) (*model.User, error) {
	return s.userRepo.UpdateUser(user)
}
func (s UserService) DeleteUser(userID int64) (int64, error) {
	return s.userRepo.Delete(userID)

}
func (s UserService) CreatUser(user *model.User) (*model.User, error) {
	pass := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	user.Protected = false
	return s.userRepo.Creat(user)
}
