package structs

import "time"

type UserSubscription struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	SubscriptionPlanID int       `json:"subscription_plan_id"`
	Price              int       `json:"price"`
	ExpiryDate         time.Time `json:"expiry_date"`
	PaymentStatus      string    `json:"payment_status"`
	SnapToken          string    `json:"snap_token"`
	Created_at         time.Time `json:"created_at"`
	Updated_at         time.Time `json:"updated_at"`
}
