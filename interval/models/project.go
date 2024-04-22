package models

import "time"

type Project struct {
	ID          int64
	Name        string
	DateCreated time.Time
}
