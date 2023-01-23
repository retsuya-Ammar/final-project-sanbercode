package routers

import (
	"database/sql"
	"final-project-sanbercode/controllers"
	"final-project-sanbercode/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter is the function to setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Movies
	moviesController := controllers.NewMoviesController(&sql.DB{})
	r.GET("/movies", moviesController.GetAll)
	r.GET("/movies/:id", moviesController.GetByID)
	r.POST("/movies", middleware.BasicAuth(), moviesController.Insert)
	r.PUT("/movies/:id", middleware.BasicAuth(), moviesController.Update)
	r.DELETE("/movies/:id", middleware.BasicAuth(), moviesController.Delete)

	// Users
	usersController := controllers.NewUserController(&sql.DB{})
	r.GET("/users", usersController.GetAll)
	r.GET("/users/:id", usersController.GetByID)
	r.POST("/register", usersController.Insert)
	r.POST("/login", usersController.Login)
	r.PUT("/users/:id", middleware.BasicAuth(), usersController.Update)
	r.DELETE("/users/:id", middleware.BasicAuth(), usersController.Delete)

	// subscription
	subscriptionController := controllers.NewSubscriptionPlansController(&sql.DB{})
	r.GET("/subscription", subscriptionController.GetAll)
	r.GET("/subscription/:id", subscriptionController.GetByID)
	r.POST("/subscription", middleware.BasicAuth(), subscriptionController.Insert)
	r.PUT("/subscription/:id", middleware.BasicAuth(), subscriptionController.Update)
	r.DELETE("/subscription/:id", middleware.BasicAuth(), subscriptionController.Delete)

	// user subscription
	userSubscriptionController := controllers.NewUserSubscriptionsController(&sql.DB{})
	r.GET("/user-subscription", userSubscriptionController.GetAll)
	r.GET("/user-subscription/:id", userSubscriptionController.GetByID)
	r.POST("/user-subscription", userSubscriptionController.Insert)
	r.PUT("/user-subscription/:id", userSubscriptionController.Update)
	r.DELETE("/user-subscription/:id", userSubscriptionController.Delete)

	return r
}
