package dto

import (
	"github.com/alphabatem/autodevgpt/models"
)

type PromptRequest struct {
	Prompt    string           `json:"prompt"`
	TechStack models.TechStack `json:"techStack"`
}
