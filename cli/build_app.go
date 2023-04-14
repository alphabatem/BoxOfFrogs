package main

import (
	"github.com/alphabatem/autodevgpt/dto"
	"github.com/alphabatem/autodevgpt/models"
	"github.com/alphabatem/autodevgpt/openai"
	"github.com/alphabatem/autodevgpt/services"
	"github.com/cloakd/common/context"
	"github.com/joho/godotenv"
	"log"
)

var devGpt *services.ProjectService

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, err := context.NewContext(
		&openai.OpenaiService{},
		&services.AgentService{},
		&services.TaskService{},
		&services.ProjectService{},
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = ctx.Run()
	if err != nil {
		panic(err)
	}

	devGpt = ctx.Service(services.PROJECT_SVC).(*services.ProjectService)
}

func main() {
	p, err := devGpt.Execute(&dto.PromptRequest{
		Prompt: "Platform that allows users to signup for a whitelist with their Solana wallet",
		TechStack: models.TechStack{
			Frontend: &models.Frontend{
				Framework: "Vue",
			},
			Backend: &models.Backend{
				Language: "Golang",
			},
			Database: &models.Database{
				Driver: "sqlite",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	log.Println("Output: ")
	log.Printf("%+v\n", p)
}
