package services

import (
	"fmt"
	"github.com/alphabatem/AutoDevGPT/models"
	"github.com/cloakd/common/services"
	"sort"
	"sync"
)

type AgentService struct {
	services.DefaultService
}

const AGENT_SVC = "agent_svc"

func (svc AgentService) Id() string {
	return AGENT_SVC
}

func (svc *AgentService) Start() error {

	return nil
}

func (svc *AgentService) CompleteTask(task *models.Task) error {
	tasks := []*models.Task{task}

	// Sort tasks based on priority (higher priority first)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority
	})

	results := make(chan *models.WorkerResponse, len(tasks))

	var wg sync.WaitGroup

	// Create a pool of workers with different tags
	workers := []*models.SampleAIWorker{
		{Tags: []string{"tag1"}},
		{Tags: []string{"tag2"}},
	}

	for _, task := range tasks {
		// Find a suitable worker for the task based on its tags
		var assignedWorker models.Worker
		for _, worker := range workers {
			if worker.CanExecute(task) {
				assignedWorker = worker
				break
			}
		}

		if assignedWorker != nil {
			wg.Add(1)
			go svc.process(task, assignedWorker, results, &wg)
		} else {
			fmt.Printf("No suitable worker found for task %d\n", task.ID)
		}
	}

	wg.Wait()
	close(results)

	for response := range results {
		fmt.Printf("Task %d completed with data: %s\n", response.TaskID, response.Data)
	}

	return nil
}

func (svc *AgentService) process(task *models.Task, worker models.Worker, results chan<- *models.WorkerResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	response := worker.Execute(task)
	results <- response
}
