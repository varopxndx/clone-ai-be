package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/sashabaranov/go-openai"
	"github.com/varopxndx/clone-ai-be/model"
)

// Service contains service client
type Service struct {
	client       *resty.Client
	openAIclient *openai.Client
}

// New returns a service struct
func New(openAIToken string) *Service {
	return &Service{
		client:       resty.New(),
		openAIclient: openai.NewClient(openAIToken),
	}
}

// GetSample gets sample data
func (s *Service) GetSample() (*model.SampleResponse, error) {
	// mocked data
	response := &model.SampleResponse{
		ID:   1,
		Name: "SomeName",
		Age:  30,
	}

	return response, nil
}

// requestAnswerFromGPT request a response from OpenAI getting the requests
func (s *Service) requestAnswerFromGPT(ctx context.Context, message string) (openai.ChatCompletionResponse, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("./chats.json")
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var messages []openai.ChatCompletionMessage
	json.Unmarshal(byteValue, &messages)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: message,
	})

	resp, err := s.openAIclient.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4,
			Messages: messages,
		},
	)
	return resp, nil
}

// GetAnswer call OpenAI API
func (s *Service) GetAnswer(ctx context.Context, message string) (*openai.ChatCompletionResponse, error) {
	// mocked data
	response, err := s.requestAnswerFromGPT(ctx, message)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
