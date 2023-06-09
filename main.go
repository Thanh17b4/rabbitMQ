package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"Thanh17b4/practice/database"
	"Thanh17b4/practice/handler"
	"Thanh17b4/practice/middleware"
	"Thanh17b4/practice/repo"
	"Thanh17b4/practice/service"
)

func main() {
	sqlDns := "root:admin@(127.0.0.1:3306)/token?parseTime=true"
	db, err := database.NewDB(sqlDns)
	if err != nil {
		fmt.Println("can not connect to database:", err.Error())
		return
	}

	if err1 != nil {
		return
	}

	//migrate database
	//driver, _ := mysql.WithInstance(db, &mysql.Config{})
	//migrateOps, err := migrate.NewWithDatabaseInstance(
	//	"file://./database/migrations",
	//	"token",
	//	driver,
	//)
	//
	//if err != nil {
	//	fmt.Println("could not migrateDB: ", err.Error())
	//	return
	//}
	//
	//err = migrateOps.Steps(6)
	//if err != nil {
	//	log.Fatal("could not migrateDB: ", err.Error())
	//	return
	//}
	//logrus.Infof("End init migrate")

	userRepo := repo.NewUser(db)
	userService := service.NewUserService(userRepo)
	userHandle := handler.NewUserHandle(userService)

	otpRepo := repo.NewOtp(db)
	newUserRepo := repo.NewUser(db)
	otpService := service.NewOtpService(newUserRepo, otpRepo)
	otpHandle := handler.NewOtpHandle(otpService)

	loginService := service.NewLogin(otpRepo, userRepo)
	loginHandle := handler.NewLoginHandle(loginService)
	//
	activateRepo := repo.NewActivate(db)
	activateService := service.NewActivate(activateRepo, userRepo, otpRepo)
	activateHandle := handler.NewActivateHandler(activateService)

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
	//r.Put("/users/login/active", activateHandle.Active)
	r.Put("/users/login/active", activateHandle.Active)
	log.Fatal(http.ListenAndServe(":4000", r))
}
