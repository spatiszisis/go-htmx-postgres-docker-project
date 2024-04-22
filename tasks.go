package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"go-jwt/models"
// 	"io"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// var errNameRequired = errors.New("name is required")
// var errProjectIDRequired = errors.New("project id is required")
// var errUserIDRequired = errors.New("user id is required")

// type TasksService struct {
// 	store Store
// }

// func NewTasksService(store Store) *TasksService {
// 	return &TasksService{store: store}
// }

// func (ts *TasksService) RegisterRoutes(r *mux.Router) {
// 	r.HandleFunc("/tasks", ts.handleCreateTask).Methods("POST")
// 	r.HandleFunc("/tasks/{id}", ts.GetTask).Methods("GET")
// }

// func (ts *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		return
// 	}

// 	defer r.Body.Close()

// 	var task *models.Task
// 	err = json.Unmarshal(body, &task)
// 	if err != nil {
// 		return
// 	}

// 	if err := validateTaskPayload(task); err != nil {
// 		return
// 	}

// 	// t, err := ts.store.CreateTask(task)
// }

// func (ts *TasksService) GetTask(w http.ResponseWriter, r *http.Request) {
// 	// Get a task by ID
// }

// func validateTaskPayload(task *models.Task) error {
// 	if task.Name == "" {
// 		return errNameRequired
// 	}

// 	if task.ProjectId == 0 {
// 		return errProjectIDRequired
// 	}

// 	if task.AssignedId == 0 {
// 		return errUserIDRequired
// 	}

// 	return nil
// }
