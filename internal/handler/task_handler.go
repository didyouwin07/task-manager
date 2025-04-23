package handler

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"

    "github.com/gorilla/mux"
    "github.com/google/uuid"
    "github.com/didyouwin07/task-manager/internal/model"
    "github.com/didyouwin07/task-manager/internal/service"
)

type TaskHandler struct {
    service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
    return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        log.Printf("Failed to decode request body: %v", err)
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    task.ID = uuid.NewString()
    h.service.Create(task)

    log.Printf("Task created with ID: %s", task.ID)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")

    page, _ := strconv.Atoi(pageStr)
    if page <= 0 {
        page = 1
    }

    limit, _ := strconv.Atoi(limitStr)
    if limit <= 0 {
        limit = 10
    }

    tasks := h.service.GetAll()

    // Filter by status
    if status != "" {
        filtered := []model.Task{}
        for _, t := range tasks {
            if strings.EqualFold(t.Status, status) {
                filtered = append(filtered, t)
            }
        }
        tasks = filtered
        log.Printf("Filtered tasks by status: %s", status)
    }

    // Pagination
    start := (page - 1) * limit
    end := start + limit
    if start >= len(tasks) {
        tasks = []model.Task{}
    } else if end > len(tasks) {
        tasks = tasks[start:]
    } else {
        tasks = tasks[start:end]
    }

    log.Printf("Returning tasks page %d with limit %d", page, limit)
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    task, found := h.service.GetByID(id)
    if !found {
        log.Printf("Task not found with ID: %s", id)
        http.NotFound(w, r)
        return
    }

    log.Printf("Retrieved task with ID: %s", id)
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        log.Printf("Failed to decode request body: %v", err)
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    task.ID = id
    updated := h.service.Update(id, task)
    if !updated {
        log.Printf("Failed to update, task not found with ID: %s", id)
        http.NotFound(w, r)
        return
    }

    log.Printf("Updated task with ID: %s", id)
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    deleted := h.service.Delete(id)
    if !deleted {
        log.Printf("Failed to delete, task not found with ID: %s", id)
        http.NotFound(w, r)
        return
    }

    log.Printf("Deleted task with ID: %s", id)
    w.WriteHeader(http.StatusNoContent)
}
