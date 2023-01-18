package structs

import "time"

type Users struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Email_verified_at time.Time `json:"email_verified_at"`
	Password          string    `json:"password"`
	Remember_token    string    `json:"remember_token"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}
