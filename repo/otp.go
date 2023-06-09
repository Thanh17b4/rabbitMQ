package repo

import (
	"database/sql"

	model "Thanh17b4/practice/model"
	"github.com/pkg/errors"
)

type OtpRepo struct {
	db *sql.DB
	u  *User
}

func NewOtp(db *sql.DB) *OtpRepo {
	return &OtpRepo{db: db}
}

//func (o *OtpRepo) CreatOTP(userOTP *model.UserOTP) (*model.UserOTP, error) {
//
//	_, err := o.database.Exec("INSERT INTO users_otp ( user_id, otp, expired_at) VALUES (? , ?, ?)", &userOTP.UserID, &userOTP.OTP, &userOTP.Expired)
//	if err != nil {
//		return nil, errors.Wrap(err, "had an error with creating OTP")
//	}
//	return userOTP, nil
//}

func (o *OtpRepo) GetUserOTP(userID int) (*model.UserOTP, error) {
	userOTP := &model.UserOTP{}
	row := o.db.QueryRow(" SELECT user_id, otp, expired_at, created_at FROM users_otp WHERE user_id = ? ORDER BY created_at DESC limit 1  ", userID)
	err1 := row.Scan(&userOTP.UserID, &userOTP.OTP, &userOTP.Expired, &userOTP.CreatAt)
	if err1 != nil {
		return nil, errors.Wrap(err1, "could not get OTP")
	}
	return userOTP, nil
}

func (o *OtpRepo) CreatOTP(email string) (*model.UserOTP, error) {
	var userOTP *model.UserOTP
	row, err := o.db.Exec("INSERT INTO users_otp (user_id, otp, expired_at) VALUE (?, ?, ?) ", userOTP.UserID, userOTP.OTP, userOTP.Expired)
	if err != nil {
		return nil, errors.Wrap(err, "had an error with creating OTP")
	}
	affectRow, err := row.LastInsertId()
	if err != nil || affectRow == 0 {
		return nil, errors.Wrap(err, "no row affected")
	}
	return userOTP, nil
}
