package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/varopxndx/clone-ai-be/model"
)

// Service contains service client
type Service struct {
	client      *resty.Client
	openAIToken string
}

// New returns a service struct
func New(openAIToken string) *Service {
	return &Service{
		client:      resty.New(),
		openAIToken: openAIToken,
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
func (s *Service) requestAnswerFromGPT(message string) (model.RequestAnswerGPT, error) {
	url := "https://api.openai.com/v1/chat/completions"
	userToken := fmt.Sprintf("Bearer %s", s.openAIToken)

	// Open our jsonFile
	jsonFile, err := os.Open("/assets/chats.json")
	if err != nil {
		return model.RequestAnswerGPT{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var messages []model.GPTMessage
	json.Unmarshal(byteValue, &messages)
	messages = append(messages, model.GPTMessage{
		Role:    "user",
		Content: message,
	})

	resp, err := s.client.
		R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", userToken).
		SetBody(fmt.Sprintf(`{"model":"gpt-4", "messages":"%s"}`, messages)).
		Get(url)

	request := model.RequestAnswerGPT{}
	if json.Unmarshal(resp.Body(), &request); err != nil {
		return model.RequestAnswerGPT{}, err
	}
	return request, nil
}

// GetAnswer call OpenAI API
func (s *Service) GetAnswer(message string) (*model.RequestAnswerGPT, error) {
	// mocked data
	response, err := s.requestAnswerFromGPT(message)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
