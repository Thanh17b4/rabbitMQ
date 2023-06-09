package repo

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"Thanh17b4/practice/model"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) Create(user *model.User) (*model.User, error) {
	_, err := u.db.Exec("INSERT INTO users ( name, email, address, password, username) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Address, user.Password, user.Username)
	if err != nil {
		return nil, errors.Wrap(err, "could not creat new user")
	}
	return user, nil
}

func (u *User) ListUser(page int64, limit int64) ([]*model.User, error) {
	var users []*model.User
	offset := (page - 1) * limit
	result, err := u.db.Query("SELECT id, name, address, username, email FROM users LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "could not get list user from database")
	}
	for result.Next() {
		u := &model.User{}
		err := result.Scan(&u.ID, &u.Name, &u.Address, &u.Username, &u.Email)
		if err != nil {
			return nil, errors.Wrap(err, "could not scan user information")
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *User) DetailUser(userID int64) (*model.User, error) {
	user := &model.User{}
	row := u.db.QueryRow(" SELECT id, name, address, username, email FROM users WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.Username, &user.Email)
	if err != nil {
		return nil, errors.Wrap(err, "userID is not correct")
	}
	return user, nil
}

func (u *User) UpdateUser(user *model.User) (*model.User, error) {
	row, err := u.db.Exec(" UPDATE users SET name = ?, address = ?, activated = ?, username = ?, email = ?, banned = ?, protected = ?, password = ? WHERE id = ?", &user.Name, &user.Address, &user.Activated, &user.Username, &user.Email, &user.Banned, &user.Protected, &user.Password, &user.ID)
	rowAffected, err := row.RowsAffected()
	if err != nil || rowAffected == 0 {
		return nil, errors.Wrap(err, "Could not update user")
	}
	//defer u.database.Close()
	return user, nil
}
func (u *User) Delete(userID int64) (int64, error) {
	result, err := u.db.Exec("DELETE FROM users WHERE id = ?", userID)
	count, err := result.RowsAffected()
	if err != nil || count == 0 {
		return 0, errors.Wrap(err, "userID is not correct")
	}
	return userID, nil
}

func (u *User) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	row := u.db.QueryRow(" SELECT id, name, address, activated, email, password FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.Activated, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "email is not correct, please try again")
	}
	return user, nil
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	row := u.db.QueryRow(" SELECT id, name, address, activated, email, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.Activated, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "username is not correct, please try again")
	}
	return user, nil
}

func (u *User) CountUsers() (int64, error) {
	var totalUsers int64
	err := u.db.QueryRow("SELECT COUNT(id) FROM users").Scan(&totalUsers)
	if err != nil {
		return 0, err
	}
	fmt.Sprintln("a", totalUsers)
	return totalUsers, nil
}
