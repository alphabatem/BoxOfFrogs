package workers

import (
	"fmt"
	"github.com/alphabatem/autodevgpt/models"
	"github.com/cloakd/common/context"
)

type SampleAIWorker struct {
	//
	Agent *Agent

	Project *Project
}

func (w SampleAIWorker) Id() string {
	return "simple"
}

func (w SampleAIWorker) Name() string {
	return fmt.Sprintf("%s Worker", w.Id())
}

func (w *SampleAIWorker) Configure(ctx *context.Context, project *Project, agent *Agent) error {
	w.Agent = agent
	w.Project = project
	return nil
}

func (w *SampleAIWorker) Execute(task *models.Task) *models.WorkerResponse {
	// Placeholder function to simulate AI agent task execution
	return &models.WorkerResponse{
		TaskID:   task.ID,
		WorkerID: w.Id(),
		Tasks:    w.ToTaskOutput(task, task.Input),
		Data:     nil,
	}
}

func (w *SampleAIWorker) String() string {
	return fmt.Sprintf("Worker: %s", w.Id())
}

func (w *SampleAIWorker) ToTaskOutput(task *models.Task, resp string) []*models.Task {
	tasks := make([]*models.Task, len(w.Agent.Output))
	for i, o := range w.Agent.Output {
		tasks[i] = models.NewTask(resp, []string{o}, (task.ID*10)+1+uint(i))
	}
	return tasks
}
