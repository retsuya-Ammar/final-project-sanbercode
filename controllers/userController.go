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

// UserController is the controller for user
type UserController struct {
	UserRepo repository.UserRepo
}

// NewUserController is the constructor for user controller
func NewUserController(db *sql.DB) *UserController {
	return &UserController{UserRepo: *repository.NewUserRepo(database.DbConnection)}
}

// GetAll is the function to get all users
func (u *UserController) GetAll(c *gin.Context) {
	users, err := u.UserRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetByID is the function to get user by id
func (u *UserController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	user, err := u.UserRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// login
func (u *UserController) Login(c *gin.Context) {
	var user structs.Users

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	user, err = u.UserRepo.Login(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Insert is the function to insert user
func (u *UserController) Insert(c *gin.Context) {
	var user structs.Users

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	user, err = u.UserRepo.Insert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update is the function to update user
func (u *UserController) Update(c *gin.Context) {
	var user structs.Users

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	user.Updated_at = time.Now()

	user, err = u.UserRepo.Update(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Delete is the function to delete user
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err = u.UserRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
