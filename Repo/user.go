package Repo

type User struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Protected bool    `json:"protected"`
	Banned    bool    `json:"banned"`
	Activated bool    `json:"activated"`
	Address   *string `json:"address"`
}
