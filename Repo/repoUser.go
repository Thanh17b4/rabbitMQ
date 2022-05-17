package Repo

import (
	"database/sql"
	"fmt"
	model "practice/Model"
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

func (u *User) DetailUser(userID int64) *model.User {
	user := &model.User{}
	_ = u.db.QueryRow(" SELECT id, name, address FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.Address)
	defer u.db.Close()
	return user
}

func (u *User) UpdateUser(user *model.User) *model.User {
	_, err := u.db.Exec(" UPDATE users SET name = ?, address = ? WHERE id = ?", &user.Name, &user.Address, user.ID)
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

func (u *User) Creat(user *model.User) (int64, error) {
	result, err := u.db.Exec("INSERT INTO users ( name, email, address) VALUES (? , ?, ?)", user.Name, user.Email, user.Address)
	if err != nil {
		fmt.Println("had error: ", err.Error())
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertID, nil
}
