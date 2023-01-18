package structs

import "time"

type SubscriptionPlan struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Featured   string    `json:"featured"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Delete_at  time.Time `json:"delete_at"`
}
