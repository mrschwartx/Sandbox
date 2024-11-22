package tasks

import (
	"context"
	"errors"
	"sync"
)

// ITaskRepository defines the interface for task operations.
type ITaskRepository interface {
	Save(ctx context.Context, task *Task) error                      // Create a new task
	GetAll(ctx context.Context) ([]*Task, error)                     // Get all tasks
	GetByID(ctx context.Context, id int) (*Task, error)              // Get a task by its ID
	UpdateByID(ctx context.Context, id int, updatedTask *Task) error // Update an existing task
	DeleteByID(ctx context.Context, id int) error                    // Delete a task by its ID
	InitData(ctx context.Context) error                              // Initialize the repository with predefined data
}

// TaskRepository implements the ITaskRepository interface.
type TaskRepository struct {
	mu     sync.RWMutex
	tasks  map[int]*Task
	nextID int
}

// NewTaskRepository creates a new instance of TaskRepository.
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// Save creates or updates a task with proper lifecycle hooks.
func (r *TaskRepository) Save(ctx context.Context, task *Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if task.ID == 0 {
		task.BeforeCreate()
		task.ID = r.nextID
		r.nextID++
	} else {
		task.BeforeSave()
	}

	r.tasks[task.ID] = task
	return nil
}

// GetAll retrieves all tasks.
func (r *TaskRepository) GetAll(ctx context.Context) ([]*Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetByID retrieves a task by its ID.
func (r *TaskRepository) GetByID(ctx context.Context, id int) (*Task, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// UpdateByID updates a task's details using lifecycle hooks.
func (r *TaskRepository) UpdateByID(ctx context.Context, id int, updatedTask *Task) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	updatedTask.ID = id
	updatedTask.CreatedAt = task.CreatedAt // Preserve original creation time
	updatedTask.BeforeSave()

	r.tasks[id] = updatedTask
	return nil
}

// DeleteByID removes a task by its ID.
func (r *TaskRepository) DeleteByID(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}

// InitData initializes the repository with predefined tasks.
func (r *TaskRepository) InitData(ctx context.Context) error {
	defaultTasks := []*Task{
		{Title: "Learn Go", Status: "Pending"},
		{Title: "Build a web app", Status: "In Progress"},
		{Title: "Write documentation", Status: "Pending"},
		{Title: "Refactor codebase", Status: "Completed"},
	}

	for _, task := range defaultTasks {
		if err := r.Save(ctx, task); err != nil {
			return err
		}
	}
	return nil
}
