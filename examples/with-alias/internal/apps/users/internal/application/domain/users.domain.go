package domain

import (
	"strings"
	"time"

	"with-alias/internal/lib"
)

type User struct {
	DueDate     lib.ITime
	Task        string
	Description string
}

func NewUser(task string, description string, dueDate lib.ITime) (*User, error) {

	v := lib.NewValidator()

	// check task
	task = strings.TrimSpace(task)
	v.Check("task", len(task) < 3, "task should be minimum 3 characters")

	// check description
	if description != "" {
		v.Check("description", len(description) < 5, "description should be minimum 5 characters")
	}

	// check DueDate
	// need more tolerance
	if dueDate.IsZero() {
		v.AddError("due_date", "due cannot be empty")
	} else {
		due := time.Now().Compare(dueDate.Time())
		v.Check("due_date", due >= 0, "due should be in future")
	}

	if v.Valid() {
		return &User{
			Task:        task,
			Description: description,
			DueDate:     dueDate,
		}, nil
	} else {
		return nil, v.Errors()
	}
}
