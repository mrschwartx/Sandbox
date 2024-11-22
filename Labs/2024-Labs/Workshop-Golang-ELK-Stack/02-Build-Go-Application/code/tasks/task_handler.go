package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// TaskHandler provides HTTP handler functions for task operations.
type TaskHandler struct {
	Repo ITaskRepository
}

// NewTaskHandler creates a new instance of TaskHandler.
func NewTaskHandler(repo ITaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}

// CreateTaskHandler handles creating a new task.
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Repo.Save(ctx, &task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetAllTasksHandler handles retrieving all tasks.
func (h *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.Repo.GetAll(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByIDHandler handles retrieving a task by ID.
func (h *TaskHandler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	task, err := h.Repo.GetByID(ctx, id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler handles updating a task by ID.
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Repo.UpdateByID(ctx, id, &updatedTask); err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTaskHandler handles deleting a task by ID.
func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Repo.DeleteByID(ctx, id); err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
