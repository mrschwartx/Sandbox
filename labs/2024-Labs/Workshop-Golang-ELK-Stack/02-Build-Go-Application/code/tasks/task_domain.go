package tasks

import (
	"time"
)

// Task defines the structure for a task in the task management.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BeforeCreate sets the creation and update timestamps for a new task.
func (t *Task) BeforeCreate() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

// BeforeSave updates the timestamp for an existing task.
func (t *Task) BeforeSave() {
	t.UpdatedAt = time.Now()
}
