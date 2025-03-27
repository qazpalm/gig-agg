package models

import "time"

type User struct {
	ID       		int  		`json:"id"`
	Username 		string 		`json:"username"`
	Email    		string	 	`json:"email"`
	PasswordHash 	string 		`json:"password_hash"`
	LastLogin 		time.Time 	`json:"last_login"`
	CreatedAt 		time.Time 	`json:"created_at"`

	RememberedToken string 		`json:"remembered_token"`

	IsAdmin 		bool 		`json:"is_admin"`
}



