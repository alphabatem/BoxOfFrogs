package services

import (
	"errors"
	"fmt"
	"github.com/alphabatem/autodevgpt/dto"
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/workers"
	"github.com/cloakd/common/services"
	"log"
)

type TaskService struct {
	services.DefaultService

	agentSvc *AgentService

	tasks map[uint]*models.Task
}

const TASK_SVC = "task_svc"

func (svc TaskService) Id() string {
	return TASK_SVC
}

func (svc *TaskService) Start() error {
	svc.tasks = map[uint]*models.Task{}

	svc.agentSvc = svc.Service(AGENT_SVC).(*AgentService)

	return nil
}

func (svc *TaskService) PromptToTask(req *dto.PromptRequest) *models.Task {
	prompt := fmt.Sprintf("%s, using the following languages and frameworks %s.", req.Prompt, req.TechStack.String())

	return models.NewTask(prompt, []string{"base"}, 1)
}

//Execute a task
func (svc *TaskService) Execute(project *workers.Project, task *models.Task) (*models.WorkerResponse, error) {
	svc.tasks[task.ID] = task
	tasks := append([]*models.Task{task}, task.Children...)

	for _, t := range tasks {
		agent := svc.taskAgent(project, t)
		if agent == nil {
			log.Printf("No suitable worker found for task (#%v) %v", t.ID, t.Tags)
			return nil, errors.New("no worker available")
		}

		return svc.process(t, agent.Worker())
	}

	return nil, errors.New("no worker available")
}

func (svc *TaskService) taskAgent(project *workers.Project, task *models.Task) *workers.Agent {
	for _, a := range project.Agents {
		if a.CanExecute(task) {
			return a
		}
	}
	return nil
}

func (svc *TaskService) process(task *models.Task, worker workers.Worker) (*models.WorkerResponse, error) {
	log.Printf("#%v - Executing on %s", task.ID, worker.Id())
	response := worker.Execute(task, nil) //TODO Wire context

	return response, nil
}
