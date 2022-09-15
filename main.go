package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Thanh17b4/practice/db"
	"github.com/Thanh17b4/practice/handler"
	"github.com/Thanh17b4/practice/middleware"
	"github.com/Thanh17b4/practice/repo"
	"github.com/Thanh17b4/practice/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	sqlDns := "root:221292@tcp(localhost:13306)/token_service?parseTime=true"
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

	r := chi.NewRouter()
	r.Post("/users/login", loginHandle.Login)

	r.Route("/token", func(r chi.Router) {
		r.With(middleware.RequestToken).Route("/users", func(r chi.Router) {
			r.Put("/{id}", userHandle.UpdateUserHandle)
			r.Get("/", userHandle.GetListUser)
			r.Get("/{id}", userHandle.GetDetailUserHandle)
			r.Delete("/{id}", userHandle.DeleteUserHandle)
			r.Post("/refresh", loginHandle.Refresh)

		})
	})

	r.Post("/users/register", userHandle.CreatUserHandle)
	r.Post("/users_otp/register/otp", otpHandle.CreatUserOTPHandle)
	r.Post("/users/login/active", activateHandle.Active)

	log.Fatal(http.ListenAndServe(":4000", r))
}
