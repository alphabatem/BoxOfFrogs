package main

import (
	"github.com/alphabatem/AutoDevGPT/openai"
	"github.com/alphabatem/AutoDevGPT/services"
	"github.com/cloakd/common/context"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, err := context.NewCtx(
		&openai.OpenaiService{},
		&services.AgentService{},
		&services.TaskService{},
		&services.HttpService{},
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = ctx.Run()
	log.Fatal(err)
}
