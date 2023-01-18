package controllers

import (
	"database/sql"
	"final-project-sanbercode/database"
	"final-project-sanbercode/repository"
	"final-project-sanbercode/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SubscriptionPlansController is the controller for subscriptionPlans
type SubscriptionPlansController struct {
	SubscriptionPlansRepo repository.SubscriptionPlanRepo
}

// NewSubscriptionPlansController is the constructor for subscriptionPlans controller
func NewSubscriptionPlansController(db *sql.DB) *SubscriptionPlansController {
	return &SubscriptionPlansController{SubscriptionPlansRepo: *repository.NewSubscriptionPlanRepo(database.DbConnection)}
}

// GetAll is the function to get all subscriptionPlans
func (s *SubscriptionPlansController) GetAll(c *gin.Context) {
	subscriptionPlans, err := s.SubscriptionPlansRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriptionPlans})
}

// GetByID is the function to get subscriptionPlans by id
func (s *SubscriptionPlansController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	subscriptionPlans, err := s.SubscriptionPlansRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriptionPlans})
}

// Insert is the function to insert subscriptionPlans
func (s *SubscriptionPlansController) Insert(c *gin.Context) {
	var subscriptionPlans structs.SubscriptionPlan

	err := c.BindJSON(&subscriptionPlans)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	subscriptionPlans.Created_at = time.Now()
	subscriptionPlans.Updated_at = time.Now()

	_, err = s.SubscriptionPlansRepo.Insert(subscriptionPlans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriptionPlans})
}

// Update is the function to update subscriptionPlans
func (s *SubscriptionPlansController) Update(c *gin.Context) {
	var subscriptionPlans structs.SubscriptionPlan

	err := c.BindJSON(&subscriptionPlans)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	subscriptionPlans.Updated_at = time.Now()

	_, err = s.SubscriptionPlansRepo.Update(subscriptionPlans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subscriptionPlans})
}

// Delete is the function to delete subscriptionPlans
func (s *SubscriptionPlansController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err = s.SubscriptionPlansRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
