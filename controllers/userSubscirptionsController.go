package controllers

import (
	"database/sql"
	"final-project-sanbercode/database"
	"final-project-sanbercode/repository"
	"final-project-sanbercode/structs"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UserSubscriptionsController is the controller for userSubscriptions
type UserSubscriptionsController struct {
	UserSubscriptionsRepo repository.UserSubscriptionRepo
}

// NewUserSubscriptionsController is the constructor for userSubscriptions controller
func NewUserSubscriptionsController(db *sql.DB) *UserSubscriptionsController {
	return &UserSubscriptionsController{UserSubscriptionsRepo: *repository.NewUserSubscriptionRepo(database.DbConnection)}
}

// GetAll is the function to get all userSubscriptions
func (u *UserSubscriptionsController) GetAll(c *gin.Context) {
	userSubscriptions, err := u.UserSubscriptionsRepo.GetAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userSubscriptions})
}

// GetByID is the function to get userSubscriptions by id
func (u *UserSubscriptionsController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	userSubscriptions, err := u.UserSubscriptionsRepo.GetByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userSubscriptions})
}

// Insert is the function to insert userSubscriptions
func (u *UserSubscriptionsController) Insert(c *gin.Context) {
	var userSubscriptions structs.UserSubscription

	err := c.BindJSON(&userSubscriptions)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	// create expriry_date time now + 3 months
	userSubscriptions.ExpiryDate = time.Now().AddDate(0, 3, 0)

	userSubscriptions.Created_at = time.Now()
	userSubscriptions.Updated_at = time.Now()

	_, err = u.UserSubscriptionsRepo.Insert(userSubscriptions)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Subscriptions Created"})
}

// Update is the function to update userSubscriptions
func (u *UserSubscriptionsController) Update(c *gin.Context) {
	var userSubscriptions structs.UserSubscription

	err := c.BindJSON(&userSubscriptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	userSubscriptions.Updated_at = time.Now()

	_, err = u.UserSubscriptionsRepo.Update(id, userSubscriptions)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Subscriptions Updated"})
}

// Delete is the function to delete userSubscriptions
func (u *UserSubscriptionsController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err = u.UserSubscriptionsRepo.Delete(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Subscriptions Deleted"})
}
