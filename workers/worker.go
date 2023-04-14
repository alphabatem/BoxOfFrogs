package workers

import (
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/openai"
	"github.com/cloakd/common/context"
)

type Worker interface {
	Id() string
	Configure(ctx *context.Context, project *Project, agent *Agent) error
	Execute(task *models.Task, context openai.ChatContext) *models.WorkerResponse
	String() string
}
