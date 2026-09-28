package tasks_transport

import core_http_server "github.com/BraveTonni/gotodo/internal/core/transport/http/server"

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
}

func NewTasksHTTPHandler(
	tasksService TasksService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{}
}
