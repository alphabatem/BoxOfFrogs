package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alphabatem/autodevgpt/workers"
	"github.com/cloakd/common/context"
	"github.com/cloakd/common/services"
	"io/ioutil"
)

type AgentService struct {
	services.DefaultService

	cx *context.Context
}

const AGENT_SVC = "agent_svc"

func (svc AgentService) Id() string {
	return AGENT_SVC
}

func (svc *AgentService) Configure(ctx *context.Context) error {
	svc.cx = ctx
	return svc.DefaultService.Configure(ctx)
}

func (svc *AgentService) Start() error {

	return nil
}

func (svc *AgentService) LoadAgents(project *workers.Project) error {
	fi, err := ioutil.ReadDir("./agents")
	if err != nil {
		return err
	}

	for _, f := range fi {
		fn := fmt.Sprintf("./agents/%s", f.Name())
		data, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}

		var agent workers.Agent
		err = json.Unmarshal(data, &agent)
		if err != nil {
			return err
		}

		err = svc.RegisterAgent(project, &agent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (svc *AgentService) RegisterAgent(project *workers.Project, agent *workers.Agent) error {
	worker := workers.GetWorker(agent.WorkerID)
	if worker == nil {
		return errors.New(fmt.Sprintf("no worker for %s", agent.WorkerID))
	}

	err := agent.SetWorker(svc.cx, project, worker)
	if err != nil {
		return err
	}

	project.Agents = append(project.Agents, agent)
	return nil
}
