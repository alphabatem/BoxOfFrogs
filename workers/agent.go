package workers

import (
	"github.com/alphabatem/autodevgpt/models"
	"github.com/cloakd/common/context"
)

type Agent struct {
	Name     string   `json:"name"`
	Prompt   string   `json:"prompt"`
	Tags     []string `json:"tags"`
	Output   []string `json:"output"` //Output tag
	WorkerID string   `json:"worker"`

	worker Worker
}

func (w *Agent) CanExecute(task *models.Task) bool {
	for _, taskTag := range task.Tags {
		for _, workerTag := range w.Tags {
			if taskTag == workerTag {
				return true
			}
		}
	}
	return false
}

func (w *Agent) SetWorker(ctx *context.Context, project *Project, worker Worker) error {
	w.worker = worker
	return w.worker.Configure(ctx, project, w)
}

func (w *Agent) Worker() Worker {
	return w.worker
}
