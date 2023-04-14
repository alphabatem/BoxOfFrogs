package workers

import (
	"fmt"
	"github.com/alphabatem/autodevgpt/models"
	"github.com/google/uuid"
)

type Project struct {
	ID        string           `json:"ID"`
	Prompt    string           `json:"prompt"`
	TechStack models.TechStack `json:"techStack"`
	//
	Board *models.Board `json:"board"`

	Agents []*Agent `json:"-"`
}

func (p *Project) BaseLocation() string {
	return fmt.Sprintf("static/projects/%s", p.ID)
}

func (p *Project) AddTask(t *models.Task) {
	p.Board.Add(t)
}

func NewProject(prompt string, stack models.TechStack) *Project {
	return &Project{
		ID:        uuid.New().String(),
		Prompt:    prompt,
		TechStack: stack,
		Board: &models.Board{
			Tasks: nil,
		},
		Agents: []*Agent{},
	}
}
