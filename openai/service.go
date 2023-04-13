package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloakd/common/services"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type OpenaiService struct {
	services.DefaultService

	BaseURI string

	client *http.Client
	apiKey string
}

const (
	OPENAI_SVC = "openai_svc"
)

func (svc OpenaiService) Id() string {
	return OPENAI_SVC
}

func (svc *OpenaiService) Start() error {
	svc.apiKey = os.Getenv("OPENAI_API_KEY")
	svc.BaseURI = "https://api.com/v1"
	svc.client = &http.Client{}

	return nil
}

func (svc *OpenaiService) CreateChat(history ChatContext, prompt string) (string, ChatContext, error) {
	return svc.CreateChatWithModel("gpt-3.5-turbo", history, prompt)
}

func (svc *OpenaiService) CreateChatWithModel(model string, history ChatContext, prompt string) (string, ChatContext, error) {
	history = append(history, ChatMessage{
		Role:    "user",
		Content: prompt,
	})

	payload := ChatCompletionRequest{
		Model:    model,
		Messages: history,
	}

	var response ChatCompletionResponse
	err := svc.httpRequest("/chat/completions", payload, &response)
	if err != nil {
		return "", history, err
	}

	history = append(history, response.Choices[0].Message)
	return response.Choices[0].Message.Content, history, nil
}

func (svc *OpenaiService) CreateImage(prompt string) (string, error) {

	var response ImageResponse
	err := svc.httpRequest("/images/generations", &ImageRequest{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}, &response)
	if err != nil {
		return "", err
	}

	return response.Data[0].URL, nil
}

func (svc *OpenaiService) httpRequest(endpoint string, payload interface{}, result interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/%s", svc.BaseURI, strings.TrimLeft(endpoint, "/"))
	log.Println(uri)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+svc.apiKey)

	resp, err := svc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s", responseBody)
		return errors.New("request to the OpenAI API failed")
	}

	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return err
	}

	return nil
}
