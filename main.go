package main

import (
	_ "database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"practice/db"
	"practice/handle"
	"practice/repo"
	"practice/service"
	"practice/token"
)

func main() {
	sqlDns := "root:nhatminh21@tcp(165.22.245.167:13306)/backend?parseTime=true"
	db, err := db.NewDB(sqlDns)
	if err != nil {
		fmt.Println("can not connect to database:", err.Error())
	}
	userRepo := repo.NewUser(db)
	userService := service.NewUserService(userRepo)
	userHandle := handle.NewUserHandle(userService)

	otpRepo := repo.NewOtp(db)
	otpService := service.NewOtpService(otpRepo)
	otpHandle := handle.NewOtpHandle(otpService)

	loginService := service.NewLogin(otpRepo, userRepo)
	loginHandle := handle.NewLoginHandle(loginService)

	activateService := service.NewActivate(otpRepo, userRepo)
	activateHandle := handle.NewActivateHandle(activateService)

	tkRepo := repo.NewUser(db)
	tkService := service.NewUserService(tkRepo)
	tkHandle := handle.NewToken(tkService)

	//var test *testing.T
	//token := token2.TestJWTMaker{test}
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandle.GetListUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandle.GetDetailUserHandle).Methods("GET")
	r.HandleFunc("/users/{id}", userHandle.UpdateUserHandle).Methods("PUT")
	r.HandleFunc("/users/register", userHandle.CreatUserHandle).Methods("POST")
	r.HandleFunc("/users/{id}", userHandle.DeleteUserHandle).Methods("DELETE")
	r.HandleFunc("/users/register/otp", otpHandle.CreatUserOTPHandle).Methods("POST")
	r.HandleFunc("/users/login", loginHandle.Login).Methods("POST")
	r.HandleFunc("/users/login/active", activateHandle.Active).Methods("POST")

	r.HandleFunc("/token/{id}", tkHandle.CreatToken).Methods("POST")
	r.HandleFunc("/verifyToken", tkHandle.VerifyToken).Methods("GET")
	r.HandleFunc("/refresh", tkHandle.Refresh).Methods("POST")

	r.HandleFunc("/create", token.Creat).Methods("POST")
	r.HandleFunc("/verify", token.Verify).Methods("GET")
	r.HandleFunc("/refresh1", token.Refresh).Methods("POST")
	//r.HandleFunc("/token", tkHandle.LoginToken).Methods("POST")
	//r.HandleFunc("/users/login/home", handle.Home).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
