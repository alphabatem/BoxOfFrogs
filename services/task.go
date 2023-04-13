package services

import (
	"github.com/alphabatem/AutoDevGPT/models"
	"github.com/cloakd/common/services"
)

type TaskService struct {
	services.DefaultService

	agentSvc *AgentService

	tasks map[int]*models.Task
}

const TASK_SVC = "task_svc"

func (svc TaskService) Id() string {
	return TASK_SVC
}

func (svc *TaskService) Start() error {

	return nil
}

//Execute a task
func (svc *TaskService) Execute(task *models.Task) error {
	svc.tasks[task.ID] = task

	return svc.agentSvc.CompleteTask(task)
}
