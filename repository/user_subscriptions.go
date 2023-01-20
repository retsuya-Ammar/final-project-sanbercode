package repository

import (
	"database/sql"
	"final-project-sanbercode/structs"
)

// UserSubscriptionRepo is the repository for user subscription
type UserSubscriptionRepo struct {
	DB *sql.DB
}

// NewUserSubscriptionRepo is the constructor for user subscription repository
func NewUserSubscriptionRepo(db *sql.DB) *UserSubscriptionRepo {
	return &UserSubscriptionRepo{DB: db}
}

// GetAll is the function to get all user subscriptions
func (s *UserSubscriptionRepo) GetAll() ([]structs.UserSubscription, error) {
	var userSubscriptions []structs.UserSubscription

	rows, err := s.DB.Query("SELECT * FROM user_subscriptions")
	if err != nil {
		return nil, err
	}

	var snapToken sql.NullString

	for rows.Next() {
		var userSubscription structs.UserSubscription
		err = rows.Scan(&userSubscription.ID, &userSubscription.UserID, &userSubscription.SubscriptionPlanID, &userSubscription.Price, &userSubscription.ExpiryDate, &userSubscription.PaymentStatus, &snapToken, &userSubscription.Created_at, &userSubscription.Updated_at)
		if err != nil {
			return nil, err
		}
		userSubscriptions = append(userSubscriptions, userSubscription)
	}

	return userSubscriptions, nil
}

// GetByID is the function to get user subscription by id
func (s *UserSubscriptionRepo) GetByID(id int) (structs.UserSubscription, error) {
	var userSubscription structs.UserSubscription

	var snapToken sql.NullString

	err := s.DB.QueryRow("SELECT * FROM user_subscriptions WHERE id = $1", id).Scan(&userSubscription.ID, &userSubscription.UserID, &userSubscription.SubscriptionPlanID, &userSubscription.Price, &userSubscription.ExpiryDate, &userSubscription.PaymentStatus, &snapToken, &userSubscription.Created_at, &userSubscription.Updated_at)
	if err != nil {
		return userSubscription, err
	}

	return userSubscription, nil
}

// Insert is the function to insert user subscription
func (s *UserSubscriptionRepo) Insert(userSubscription structs.UserSubscription) (structs.UserSubscription, error) {
	err := s.DB.QueryRow("INSERT INTO user_subscriptions (user_id, subscription_plan_id, price, expired_date, payment_status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", userSubscription.UserID, userSubscription.SubscriptionPlanID, userSubscription.Price, userSubscription.ExpiryDate, userSubscription.PaymentStatus, userSubscription.Created_at, userSubscription.Updated_at).Scan(&userSubscription.ID)
	if err != nil {
		return userSubscription, err
	}

	return userSubscription, nil
}

// Update is the function to update user subscription parameter id
func (s *UserSubscriptionRepo) Update(id int, userSubscription structs.UserSubscription) (structs.UserSubscription, error) {
	_, err := s.DB.Exec("UPDATE user_subscriptions SET user_id = $1, subscription_plan_id = $2, price = $3, expired_date = $4, payment_status = $5, updated_at = $6 WHERE id = $7", userSubscription.UserID, userSubscription.SubscriptionPlanID, userSubscription.Price, userSubscription.ExpiryDate, userSubscription.PaymentStatus, userSubscription.Updated_at, id)

	if err != nil {
		return userSubscription, err
	}

	return userSubscription, nil
}

// Delete is the function to delete user subscription
func (s *UserSubscriptionRepo) Delete(id int) error {
	_, err := s.DB.Exec("DELETE FROM user_subscriptions WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
