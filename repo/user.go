package repo

import (
	"database/sql"
	"fmt"
	//"golang.org/x/crypto/bcrypt"
	model "practice/model"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) ListUser(page int64, limit int64) ([]*model.User, error) {
	var users []*model.User
	offset := (page - 1) * limit
	result, err := u.db.Query("SELECT id, name, address FROM users LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		u := &model.User{}
		err := result.Scan(&u.ID, &u.Name, &u.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *User) DetailUser(userID int64) (*model.User, error) {
	user := &model.User{}
	row := u.db.QueryRow(" SELECT id, name, address, username, password FROM users WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.Username, &user.Password)
	if err != nil {
		fmt.Println("id is not available: ", err.Error())
		return nil, err
	}
	//defer u.db.Close()
	//fmt.Println("", user)
	return user, nil
}

func (u *User) UpdateUser(user *model.User) *model.User {
	_, err := u.db.Exec(" UPDATE users SET name = ?, address = ?, activated = ? WHERE id = ?", &user.Name, &user.Address, &user.Activated, user.ID)
	if err != nil {
		fmt.Println("can not connect to SQL: ", err.Error())
	}
	defer u.db.Close()
	return user
}
func (u *User) Delete(userID int64) (int64, error) {
	row, err := u.db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		fmt.Println("can not connect to SQL: ", err.Error())
	}
	result, err := row.RowsAffected()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (u *User) Creat(user *model.User) (*model.User, error) {
	result, err := u.db.Exec("INSERT INTO users ( name, email, address, password) VALUES (? , ?, ?, ?)", user.Name, user.Email, user.Address, user.Password)
	if err != nil {
		fmt.Println("had error: ", err.Error())
		return nil, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	fmt.Println("inserted userID: ", insertID)
	return user, nil
}
func (u *User) GetUserByEmail(email string) *model.User {
	user := &model.User{}
	row := u.db.QueryRow(" SELECT id, name, address, password, activated, email FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Name, &user.Address, &user.Password, &user.Activated, &user.Email)
	if err != nil {
		fmt.Println("could not get user information: ", err.Error())
		return nil
	}
	return user
}
