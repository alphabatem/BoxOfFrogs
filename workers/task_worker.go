package workers

import (
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/openai"
	"log"
	"strings"
)

type TaskWorker struct {
	SampleAIWorker
}

func (w TaskWorker) Id() string {
	return "task"
}

func (w *TaskWorker) Execute(task *models.Task, context openai.ChatContext) *models.WorkerResponse {

	taskText := strings.Split(task.Input, "\n## ")
	tasks := make([]*models.Task, len(taskText))

	log.Printf("Split: %v", len(taskText))

	output := []string{}
	for i, t := range taskText {
		if strings.TrimSpace(t) == "" {
			continue
		}

		output = append(output, t)
		tasks[i] = models.NewTask(t, w.Agent.Output, task.ID+uint(i))
	}

	return &models.WorkerResponse{
		TaskID:   task.ID,
		WorkerID: w.Id(),
		Data:     output,
		Tasks:    tasks,
	}
}
