package main

import (
	"fmt"
	"go-jwt/cmd/go-jwt/handlers"
	"go-jwt/interval/repositories"
	"go-jwt/interval/services"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection parameters
	username := "gojwt_user"
	password := "gojwt_password"
	dbname := "gojwt_db"
	port := "5433"

	db, err := repositories.NewDB(username, password, dbname, port)
	if err != nil {
		fmt.Println("Error creating DB connection:", err)
		return
	}

	router := gin.Default()
	jsonContentTypeMiddleware(router)

	//Start APIS -------

	//For Task
	// Create a new task repository with the DB connection
	taskRepo := repositories.NewTaskRepository(db)
	// Create a new task service with the task repository
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)
	router.GET("/task/:id", taskHandler.GetTaskById)
	router.POST("/task", taskHandler.CreateTask)
	router.PUT("/task/:id", taskHandler.UpdateTask)
	router.DELETE("/task/:id", taskHandler.DeleteTask)
	router.GET("/task/user/:id", taskHandler.GetTasksByUserID)

	//For User
	// Create a new user repository with the DB connection
	userRepo := repositories.NewUserRepository(db)
	// Create a new user service with the user repository
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	router.GET("/user/:id", userHandler.GetUserById)
	router.POST("/user", userHandler.CreateUser)
	router.PUT("/user/:id", userHandler.UpdateUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)

	//End APIS -------

	if err := router.Run(":9000"); err != nil {
		fmt.Println("Error starting server:", err)
	}

	//Create a task
	// task := models.Task{
	// 	ID:          1,
	// 	Name:        "Task 1",
	// 	Status:      &models.Status{ID: 1},
	// 	Project:     &models.Project{ID: 1},
	// 	AssignedFor: &models.User{ID: 1},
	// 	DateCreated: time.Now(),
	// }

	// err = taskService.CreateTask(context.Background(), &task)
	// if err != nil {
	// 	fmt.Println("Error creating task:", err)
	// 	return
	// }
}

func jsonContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
