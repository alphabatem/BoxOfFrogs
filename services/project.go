package services

import (
	"errors"
	"fmt"
	"github.com/alphabatem/autodevgpt/dto"
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/workers"
	"github.com/cloakd/common/services"
	"io/ioutil"
	"log"
	"os"
)

type ProjectService struct {
	services.DefaultService

	activeProjects map[string]*workers.Project

	taskSvc  *TaskService
	agentSvc *AgentService
}

const PROJECT_SVC = "project_svc"

func (svc ProjectService) Id() string {
	return PROJECT_SVC
}

func (svc *ProjectService) Start() error {
	svc.taskSvc = svc.Service(TASK_SVC).(*TaskService)
	svc.agentSvc = svc.Service(AGENT_SVC).(*AgentService)

	return nil
}

func (svc *ProjectService) Status(projectID string) (*workers.Project, error) {
	v, ok := svc.activeProjects[projectID]
	if !ok {
		return nil, errors.New("not found")
	}

	return v, nil
}

func (svc *ProjectService) Execute(req *dto.PromptRequest) (*workers.Project, error) {
	log.Printf("Executing Prompt: %s", req.Prompt)
	log.Printf("Tech Stack: %s", req.TechStack.String())

	p := workers.NewProject(req.Prompt, req.TechStack)

	//Add our base task
	p.AddTask(svc.taskSvc.PromptToTask(req))

	return svc.Run(p)
}

func (svc *ProjectService) Run(p *workers.Project) (*workers.Project, error) {
	err := svc.createWorkingDir(p)
	if err != nil {
		return nil, err
	}

	err = svc.agentSvc.LoadAgents(p)
	if err != nil {
		return nil, err
	}

	err = svc.processTasks(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (svc *ProjectService) processTasks(p *workers.Project) error {
	taskLen := len(p.Board.Tasks)

	//TODO Change this tasks to a channel that we listen (range over)
	for i := 0; i < taskLen; i++ {
		t := p.Board.Tasks[i]
		resp, err := svc.taskSvc.Execute(p, t)
		if err != nil {
			return err
		}

		for _, st := range resp.Tasks {
			p.AddTask(st)
			taskLen++
		}

		if resp != nil {
			err = svc.saveOutput(p, resp)
			if err != nil {
				log.Printf("#%v - Failed to save: %s", resp.TaskID, err)
			}
		}
	}
	return nil
}

func (svc *ProjectService) createWorkingDir(project *workers.Project) error {
	return os.Mkdir(fmt.Sprintf("static/projects/%s", project.ID), 0755)
}

func (svc *ProjectService) saveOutput(p *workers.Project, resp *models.WorkerResponse) error {
	for ix, r := range resp.Data {
		fp := fmt.Sprintf("%s/%v%v_%s.%s", p.BaseLocation(), resp.TaskID, ix, resp.WorkerID, "md")
		log.Printf("Saving output: %s", fp)
		err := ioutil.WriteFile(fp, []byte(r), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
