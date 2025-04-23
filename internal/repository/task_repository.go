package repository

import (
    "sync"
    "github.com/didyouwin07/task-manager/internal/model"
)

type TaskRepository struct {
    mu    sync.Mutex
    tasks map[string]model.Task
}

func NewTaskRepository() *TaskRepository {
    return &TaskRepository{
        tasks: make(map[string]model.Task),
    }
}

func (r *TaskRepository) Create(task model.Task) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.tasks[task.ID] = task
}

func (r *TaskRepository) GetAll() []model.Task {
    r.mu.Lock()
    defer r.mu.Unlock()
    tasks := []model.Task{}
    for _, t := range r.tasks {
        tasks = append(tasks, t)
    }
    return tasks
}

func (r *TaskRepository) GetByID(id string) (model.Task, bool) {
    r.mu.Lock()
    defer r.mu.Unlock()
    task, ok := r.tasks[id]
    return task, ok
}

func (r *TaskRepository) Update(id string, task model.Task) bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.tasks[id]; !exists {
        return false
    }
    r.tasks[id] = task
    return true
}

func (r *TaskRepository) Delete(id string) bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.tasks[id]; !exists {
        return false
    }
    delete(r.tasks, id)
    return true
}
