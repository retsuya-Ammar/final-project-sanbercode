package structs

import "time"

type SubscriptionPlan struct {
	ID                      int       `json:"id"`
	Name                    string    `json:"name"`
	Price                   int       `json:"price"`
	Active_period_in_months int       `json:"active_period_in_months"`
	Featured                string    `json:"featured"`
	Created_at              time.Time `json:"created_at"`
	Updated_at              time.Time `json:"updated_at"`
}
