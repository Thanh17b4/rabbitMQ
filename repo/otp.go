package repo

import (
	"database/sql"
	"fmt"
	model "practice/model"
)

type OtpRepo struct {
	db *sql.DB
}

func NewOtp(db *sql.DB) *OtpRepo {
	return &OtpRepo{db: db}
}
func (o *OtpRepo) CreatOTP(otp *model.UserOTP) (*model.UserOTP, error) {
	result, err := o.db.Exec("INSERT INTO users_otp ( user_id, code, expired_at) VALUES (? , ?, ?)", otp.UserID, otp.OTP, otp.Expired)
	fmt.Println("expirationTime:", otp.Expired)
	if err != nil {
		fmt.Println("had an error with creating OTP: ", err.Error())
		return nil, err
	}
	effectRow, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("RowsAffected: ", effectRow)
	return otp, nil
}

func (o *OtpRepo) GetOTP(userID int) *model.UserOTP {
	userOtp := &model.UserOTP{}
	row := o.db.QueryRow(" SELECT user_id, code, expired_at, created_at FROM users_otp WHERE user_id = ? ORDER BY created_at DESC limit 1  ", userID)
	//fmt.Println("userID:", userID)
	err := row.Scan(&userOtp.UserID, &userOtp.OTP, &userOtp.Expired, &userOtp.CreatAt)
	if err != nil {
		fmt.Println("could not get OTP: ", err.Error())
		return nil
	}
	//fmt.Println("ss:", userOtp)
	return userOtp
}
