package router

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/didyouwin07/task-manager/internal/handler"
    "github.com/didyouwin07/task-manager/internal/repository"
    "github.com/didyouwin07/task-manager/internal/service"
)

func SetupRouter() http.Handler {
    repo := repository.NewTaskRepository()
    svc := service.NewTaskService(repo)
    h := handler.NewTaskHandler(svc)

    r := mux.NewRouter()
    r.HandleFunc("/tasks", h.CreateTask).Methods("POST")
    r.HandleFunc("/tasks", h.GetAllTasks).Methods("GET")
    r.HandleFunc("/tasks/{id}", h.GetTaskByID).Methods("GET")
    r.HandleFunc("/tasks/{id}", h.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", h.DeleteTask).Methods("DELETE")

    return r
}
