package models

import "time"

type Task struct {
	ID          int64
	Name        string
	Status      *Status
	Project     *Project
	AssignedFor *User
	DateCreated time.Time
}
