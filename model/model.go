package model

import (
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Protected int    `json:"protected"`
	Banned    int    `json:"banned"`
	Activated int    `json:"activated"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}
type UserOTP struct {
	UserID  int       `json:"userID"`
	OTP     int       `json:"otp"`
	Expired time.Time `json:"expired"`
	CreatAt time.Time `json:"creatAt"`
}

var Token string
