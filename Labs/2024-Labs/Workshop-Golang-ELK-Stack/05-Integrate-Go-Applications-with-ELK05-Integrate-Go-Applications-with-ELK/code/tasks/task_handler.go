package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

// TaskHandler provides HTTP handler functions with added logging.
type TaskHandler struct {
	Repo   ITaskRepository
	Logger *zap.Logger
}

// NewTaskHandler creates a new instance of TaskHandler with logger.
func NewTaskHandler(repo ITaskRepository, logger *zap.Logger) *TaskHandler {
	return &TaskHandler{Repo: repo, Logger: logger}
}

// CreateTaskHandler handles creating a new task and logs the event.
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		h.Logger.Error("Invalid request payload for creating task", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Repo.Save(ctx, &task); err != nil {
		h.Logger.Error("Failed to create task", zap.Error(err))
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("Task created", zap.Int("id", task.ID), zap.String("title", task.Title))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetAllTasksHandler handles retrieving all tasks and logs the event.
func (h *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.Repo.GetAll(ctx)
	if err != nil {
		h.Logger.Error("Failed to retrieve tasks", zap.Error(err))
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	h.Logger.Info("Retrieved all tasks", zap.Int("count", len(tasks)))
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByIDHandler handles retrieving a task by ID with logging.
func (h *TaskHandler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		h.Logger.Error("Invalid task ID", zap.String("id", idStr))
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	task, err := h.Repo.GetByID(ctx, id)
	if err != nil {
		h.Logger.Error("Task not found", zap.Int("id", id), zap.Error(err))
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	h.Logger.Info("Retrieved task by ID", zap.Int("id", id), zap.String("title", task.Title))
	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler handles updating a task by ID and logs the event.
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		h.Logger.Error("Invalid task ID for update", zap.String("id", idStr))
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		h.Logger.Error("Invalid request payload for updating task", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Repo.UpdateByID(ctx, id, &updatedTask); err != nil {
		h.Logger.Error("Failed to update task", zap.Int("id", id), zap.Error(err))
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	h.Logger.Info("Updated task", zap.Int("id", id), zap.String("title", updatedTask.Title))
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTaskHandler handles deleting a task by ID with logging.
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		h.Logger.Error("Invalid task ID for deletion", zap.String("id", idStr))
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Repo.DeleteByID(ctx, id); err != nil {
		h.Logger.Error("Failed to delete task", zap.Int("id", id), zap.Error(err))
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	h.Logger.Info("Deleted task", zap.Int("id", id))
	w.WriteHeader(http.StatusNoContent)
}
