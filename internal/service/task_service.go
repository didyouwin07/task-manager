package service

import (
    "github.com/didyouwin07/task-manager/internal/model"
    "github.com/didyouwin07/task-manager/internal/repository"
)

type TaskService struct {
    repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) Create(task model.Task) {
    s.repo.Create(task)
}

func (s *TaskService) GetAll() []model.Task {
    return s.repo.GetAll()
}

func (s *TaskService) GetByID(id string) (model.Task, bool) {
    return s.repo.GetByID(id)
}

func (s *TaskService) Update(id string, task model.Task) bool {
    return s.repo.Update(id, task)
}

func (s *TaskService) Delete(id string) bool {
    return s.repo.Delete(id)
}
