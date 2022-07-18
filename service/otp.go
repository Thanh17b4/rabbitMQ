package service

import (
	"crypto/tls"
	"fmt"
	model "github.com/Thanh17b4/practice/model"
	goMail "gopkg.in/mail.v2"
	"math/rand"
	"time"
)

type OtpRepo interface {
	CreatOTP(otp *model.UserOTP) (*model.UserOTP, error)
	GetOTP(userID int) *model.UserOTP
}
type OtpService struct {
	userRepo UserRepo
	otpRepo  OtpRepo
}

func NewOtpService(userRepo UserRepo, otpRepo OtpRepo) *OtpService {
	return &OtpService{userRepo: userRepo, otpRepo: otpRepo}
}
func (s OtpService) sendEmail(code int64, subject string, receive string, name string) {
	m := goMail.NewMessage()
	m.SetHeader("From", "thanhpv@vmodev.com")
	m.SetHeader("To", receive)
	m.SetHeader("Subject", subject)
	msg := fmt.Sprintf("Hello: %s 様, here is your code: %d", name, code)
	m.SetBody("text/plain", msg)
	// Settings for SMTP server
	d := goMail.NewDialer("smtp.gmail.com", 587, "thanhpv@vmodev.com", "qxpkehhtatnwzzok")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
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
	code := rand.Intn(max-min) + min

	fmt.Println("code ", code)
	otp.OTP = code
	user, err := s.userRepo.DetailUser(int64(otp.UserID))
	if err != nil {
		fmt.Println("userID is not valid: ")
		return nil, err
	}
	receive := user.Email
	name := user.Name
	// send email
	s.sendEmail(int64(code), "test send email", receive, name)
	return s.otpRepo.CreatOTP(otp)
}
func (s OtpService) GetOTPs(userID int) *model.UserOTP {
	return s.otpRepo.GetOTP(userID)
}
