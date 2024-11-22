package tasks

import (
	"context"
	"errors"
	"sync"

	"go.uber.org/zap"
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

// TaskRepository implements the ITaskRepository interface with added logging.
type TaskRepository struct {
	mu     sync.RWMutex
	tasks  map[int]*Task
	nextID int
	logger *zap.Logger
}

// NewTaskRepository creates a new instance of TaskRepository with logger.
func NewTaskRepository(logger *zap.Logger) *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]*Task),
		nextID: 1,
		logger: logger,
	}
}

// Save creates or updates a task with proper lifecycle hooks and logging.
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
		r.logger.Info("Creating new task", zap.String("title", task.Title), zap.Int("id", task.ID))
	} else {
		task.BeforeSave()
		r.logger.Info("Updating existing task", zap.Int("id", task.ID))
	}

	r.tasks[task.ID] = task
	return nil
}

// GetAll retrieves all tasks and logs the action.
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

	r.logger.Info("Retrieving all tasks", zap.Int("count", len(tasks)))
	return tasks, nil
}

// GetByID retrieves a task by its ID with logging.
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
		r.logger.Error("Task not found", zap.Int("id", id))
		return nil, errors.New("task not found")
	}
	r.logger.Info("Retrieved task", zap.Int("id", id), zap.String("title", task.Title))
	return task, nil
}

// UpdateByID updates a task's details with logging.
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
		r.logger.Error("Task not found for update", zap.Int("id", id))
		return errors.New("task not found")
	}

	updatedTask.ID = id
	updatedTask.CreatedAt = task.CreatedAt // Preserve original creation time
	updatedTask.BeforeSave()

	r.tasks[id] = updatedTask
	r.logger.Info("Updated task", zap.Int("id", id), zap.String("title", updatedTask.Title))
	return nil
}

// DeleteByID removes a task by its ID and logs the event.
func (r *TaskRepository) DeleteByID(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		r.logger.Error("Task not found for deletion", zap.Int("id", id))
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	r.logger.Info("Deleted task", zap.Int("id", id))
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
