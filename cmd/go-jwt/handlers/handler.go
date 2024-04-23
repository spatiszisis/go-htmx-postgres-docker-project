package handlers

import "go-jwt/interval/repositories"

type Handler struct {
	task *repositories.TaskRepository
}

func New(task *repositories.TaskRepository) *Handler {
	return &Handler{task: task}
}
