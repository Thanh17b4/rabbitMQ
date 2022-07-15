package service

import (
	"crypto/tls"
	"fmt"
	goMail "gopkg.in/mail.v2"
	"math/rand"
	model "practice/model"
	"time"
)

type OtpRepo interface {
	CreatOTP(otp *model.UserOTP) (*model.UserOTP, error)
	GetOTP(userID int) *model.UserOTP
}
type OtpService struct {
	otpRepo OtpRepo
}

func NewOtpService(otpRepo OtpRepo) *OtpService {
	return &OtpService{otpRepo: otpRepo}
}
func sendEmail(code int64, subject string, receive string) {
	m := goMail.NewMessage()
	// Set E-Mail sender
	m.SetHeader("From", "thanhpv@vmodev.com")
	// Set E-Mail receivers
	m.SetHeader("To", "gdpphamthanh@gmail.com")
	// Set E-Mail subject
	m.SetHeader("Subject", subject)
	// Set E-Mail body. You can set plain text or html with text/html
	msg := fmt.Sprintf("Hello: %s, here is your code: %d", receive, code)
	m.SetBody("text/plain", msg)
	// Settings for SMTP server
	d := goMail.NewDialer("smtp.gmail.com", 587, "thanhpv@vmodev.com", "qxpkehhtatnwzzok")
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
func (s OtpService) CreatOTPs(otp *model.UserOTP) (*model.UserOTP, error) {
	//var user *model.User
	otp.CreatAt = time.Now()
	otp.Expired = otp.CreatAt.Add(time.Second * 300)
	min, max := 100000, 999999
	code := min + rand.Intn(max-min)

	fmt.Println("code ", code)

	otp.OTP = code

	// send email
	sendEmail(int64(code), "test send email", "gdpphamthanh@gmail.com")

	return s.otpRepo.CreatOTP(otp)
}
func (s OtpService) GetOTPs(userID int) *model.UserOTP {
	return s.otpRepo.GetOTP(userID)
}
