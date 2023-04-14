package workers

import (
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/openai"
	"github.com/cloakd/common/context"
	"strings"
)

type OpenAiWorker struct {
	SampleAIWorker

	context openai.ChatContext
	openai  *openai.OpenaiService
}

func (w OpenAiWorker) Id() string {
	return "openai"
}

func (w *OpenAiWorker) Configure(ctx *context.Context, project *Project, agent *Agent) error {
	w.openai = ctx.Service(openai.OPENAI_SVC).(*openai.OpenaiService)

	return w.SampleAIWorker.Configure(ctx, project, agent)
}

func (w *OpenAiWorker) UpdateContext(context openai.ChatContext) {
	//TODO Should we just overwrite here?
	if context != nil {
		w.context = append(w.context, context...)
	}
}

func (w *OpenAiWorker) Execute(task *models.Task, context openai.ChatContext) *models.WorkerResponse {
	w.UpdateContext(context)
	prompt := strings.ReplaceAll(w.Agent.Prompt, "{{description}}", task.Input)

	//log.Println(prompt)
	resp, ctx, err := w.openai.CreateChat(w.context, prompt)
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
