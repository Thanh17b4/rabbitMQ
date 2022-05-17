package main

import (
	_ "database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"practice/DB"
	"practice/Handle"
	"practice/Repo"
	"practice/Service"
)

func main() {
	sqlDns := "root:nhatminh21@tcp(165.22.245.167:13306)/backend"
	db, err := DB.NewDB(sqlDns)
	if err != nil {
		fmt.Println("can not connect to database:", err.Error())
	}
	userRepo := Repo.NewUser(db)
	userService := Service.NewUserService(userRepo)
	userHandle := Handle.NewUserHandle(userService)
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandle.GetListUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandle.GetDetailUserHandle).Methods("GET")
	r.HandleFunc("/users/{id}", userHandle.UpdateUserHandle).Methods("PUT")
	r.HandleFunc("/users", userHandle.CreatUserHandle).Methods("POST")
	r.HandleFunc("/users/{id}", userHandle.DeleteUserHandle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}
