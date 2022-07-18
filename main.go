package main

import (
	_ "database/sql"
	"fmt"
	"github.com/Thanh17b4/practice/db"
	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/repo"
	"github.com/Thanh17b4/practice/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	sqlDns := "root:nhatminh21@tcp(165.22.245.167:13306)/backend?parseTime=true"
	db, err := db.NewDB(sqlDns)
	if err != nil {
		fmt.Println("can not connect to database:", err.Error())
	}
	userRepo := repo.NewUser(db)
	userService := service.NewUserService(userRepo)
	userHandle := handler.NewUserHandle(userService)

	otpRepo := repo.NewOtp(db)
	newUserRepo := repo.NewUser(db)
	otpService := service.NewOtpService(newUserRepo, otpRepo)
	otpHandle := handler.NewOtpHandle(otpService)

	loginService := service.NewLogin(otpRepo, userRepo)
	loginHandle := handler.NewLoginHandle(loginService)

	activateService := service.NewActivate(otpRepo, userRepo)
	activateHandle := handler.NewActivateHandle(activateService)

	tkRepo := repo.NewUser(db)
	tkService := service.NewUserService(tkRepo)
	tkHandle := handler.NewToken(tkService)

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

	//r.HandleFunc("/token", tkHandle.LoginToken).Methods("POST")
	//r.HandleFunc("/users/login/home", handler.Home).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
