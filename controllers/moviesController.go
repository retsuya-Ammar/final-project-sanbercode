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

// MoviesController is the controller for movies
type MoviesController struct {
	MoviesRepo repository.MoviesRepo
}

// NewMoviesController is the constructor for movies controller
func NewMoviesController(db *sql.DB) *MoviesController {
	return &MoviesController{MoviesRepo: *repository.NewMoviesRepo(database.DbConnection)}
}

// GetAll is the function to get all movies
func (m *MoviesController) GetAll(c *gin.Context) {
	movies, err := m.MoviesRepo.GetAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

// GetByID is the function to get movies by id
func (m *MoviesController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	movies, err := m.MoviesRepo.GetByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

// Insert is the function to insert movies with create_at and update_at
func (m *MoviesController) Insert(c *gin.Context) {
	var movies structs.Movies

	err := c.BindJSON(&movies)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	movies.Created_at = time.Now()
	movies.Updated_at = time.Now()

	_, err = m.MoviesRepo.Insert(movies, database.DbConnection)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Insert movies success"})
}

// Update is the function to update movies
func (m *MoviesController) Update(c *gin.Context) {
	var movies structs.Movies

	err := c.BindJSON(&movies)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	movies.Updated_at = time.Now()

	_, err = m.MoviesRepo.Update(id, movies)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update movies success"})
}

// Delete is the function to delete movies
func (m *MoviesController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	err = m.MoviesRepo.Delete(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete movies success"})
}
