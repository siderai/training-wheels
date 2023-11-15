package main

import (
	"database/sql"
	"fitness-app/internal/app/controller"
	"fitness-app/internal/app/repository"
	"fitness-app/internal/app/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize the database connection
	db, err := sql.Open("postgres", "your-database-connection-string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository and service instances
	userRepository := repository.UserRepository{DB: db}
	userService := service.UserService{UserRepository: userRepository}
	userController := controller.UserController{UserService: userService}

	// Initialize the web server
	router := gin.Default()

	// Define API routes
	router.GET("/users/:id", userController.GetUserByID)
	router.POST("/users", userController.AddUser)

	// Start the web server
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server is running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
