package repo

import (
	"database/sql"
	model "github.com/Thanh17b4/practice/model"
	"github.com/pkg/errors"
)

type OtpRepo struct {
	db *sql.DB
}

func NewOtp(db *sql.DB) *OtpRepo {
	return &OtpRepo{db: db}
}
func (o *OtpRepo) CreatOTP(otp *model.UserOTP) (*model.UserOTP, error) {
	_, err := o.db.Exec("INSERT INTO users_otp ( user_id, code, expired_at) VALUES (? , ?, ?)", otp.UserID, otp.OTP, otp.Expired)
	if err != nil {
		return nil, errors.Wrap(err, "had an error with creating OTP")
	}
	return otp, nil
}

func (o *OtpRepo) GetOTP(userID int) (*model.UserOTP, error) {
	userOtp := &model.UserOTP{}
	row := o.db.QueryRow(" SELECT user_id, code, expired_at, created_at FROM users_otp WHERE user_id = ? ORDER BY created_at DESC limit 1  ", userID)
	err := row.Scan(&userOtp.UserID, &userOtp.OTP, &userOtp.Expired, &userOtp.CreatAt)
	if err != nil {
		return nil, errors.Wrap(err, "could not get OTP")
	}
	return userOtp, nil
}
