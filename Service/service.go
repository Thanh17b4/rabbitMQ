package Service

import model "practice/Model"

type UserRepo interface {
	ListUser(page int64, limit int64) ([]*model.User, error)
	DetailUser(userID int64) *model.User
	UpdateUser(user *model.User) *model.User
	Delete(userID int64) (int64, error)
	Creat(user *model.User) (int64, error)
}
type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *userService {
	return &userService{userRepo: userRepo}
}
func (s userService) GetListUser(page int64, limit int64) ([]*model.User, error) {
	return s.userRepo.ListUser(page, limit)
}

func (s userService) GetDetailUser(userID int64) *model.User {
	return s.userRepo.DetailUser(userID)

}
func (s userService) UpdateUserService(user *model.User) *model.User {
	return s.userRepo.UpdateUser(user)
}
func (s userService) DeleteUser(userID int64) (int64, error) {
	return s.userRepo.Delete(userID)

}
func (s userService) CreatUser(user *model.User) (int64, error) {
	return s.userRepo.Creat(user)
}
