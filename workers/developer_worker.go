package workers

import (
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/openai"
	"github.com/cloakd/common/context"
	"strings"
)

type DeveloperWorker struct {
	OpenAiWorker
}

func (w DeveloperWorker) Id() string {
	return "developer"
}

func (w *DeveloperWorker) Configure(ctx *context.Context, project *Project, agent *Agent) error {
	return w.OpenAiWorker.Configure(ctx, project, agent)
}

func (w *DeveloperWorker) Execute(task *models.Task, context openai.ChatContext) *models.WorkerResponse {
	w.UpdateContext(context)
	prompt := strings.ReplaceAll(w.Agent.Prompt, "{{description}}", task.Input)

	//log.Println(prompt)
	resp, ctx, err := w.openai.CreateChat(w.context, w.ContextualisePrompt(prompt))
	if err != nil {
		return &models.WorkerResponse{Error: err}
	}

	w.context = ctx
	return &models.WorkerResponse{
		TaskID:   task.ID,
		WorkerID: w.Id(),
		Data:     []string{resp},
		Tasks:    w.ToTaskOutput(task, resp),
	}
}

func (w *DeveloperWorker) ContextualisePrompt(prompt string) string {
	if w.Project.TechStack.Backend != nil {
		prompt = strings.ReplaceAll(prompt, "{{backend.language}}", w.Project.TechStack.Backend.Language)
	}
	if w.Project.TechStack.Frontend != nil {
		prompt = strings.ReplaceAll(prompt, "{{frontend.framework}}", w.Project.TechStack.Frontend.Framework)
	}
	if w.Project.TechStack.Database != nil {
		prompt = strings.ReplaceAll(prompt, "{{database.driver}}", w.Project.TechStack.Database.Driver)
	}

	return prompt
}
