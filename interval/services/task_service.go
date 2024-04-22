package services

import (
	"context"
	"fmt"
	"go-jwt/interval/models"
	"go-jwt/interval/repositories"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetTaskById(ctx context.Context, id int) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id int) error
	GetTasksByUserID(ctx context.Context, userID int) ([]*models.Task, error)
}

type taskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) TaskService {
	return &taskService{taskRepo}
}

func (s *taskService) CreateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Create(ctx, task)
}

func (s *taskService) GetTaskById(ctx context.Context, taskID int) (*models.Task, error) {
	task, err := s.taskRepo.GetById(ctx, taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, fmt.Errorf("task with ID %d not found", taskID)
	}
	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Update(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, taskID int) error {
	return s.taskRepo.Delete(ctx, taskID)
}

func (s *taskService) GetTasksByUserID(ctx context.Context, userID int) ([]*models.Task, error) {
	return s.taskRepo.GetByUserID(ctx, userID)
}
