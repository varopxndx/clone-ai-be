package usecase

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/varopxndx/clone-ai-be/model"

	"github.com/rs/zerolog"
)

// Service has the service layer methods
type Service interface {
	GetSample() (*model.SampleResponse, error)
	GetAnswer(ctx context.Context, message string) (*openai.ChatCompletionResponse, error)
}

// Usecase structure
type Usecase struct {
	service Service
	logger  zerolog.Logger
}

// New creates a usecase
func New(service Service, logger zerolog.Logger) *Usecase {
	return &Usecase{
		service: service,
		logger:  logger,
	}
}

// GetSample gets sample data
func (u *Usecase) GetSample() (*model.SampleResponse, error) {
	// bussiness logic
	response, err := u.service.GetSample()
	if err != nil {
		u.logger.Error().Msg("GetSample: getting sample data")
		return nil, err
	}

	return response, nil
}

// GetAnswer Call the service layout to get the message
func (u *Usecase) GetAnswer(ctx context.Context, message string) (*model.Answer, error) {
	// bussiness logic
	response, err := u.service.GetAnswer(ctx, message)
	if err != nil {
		u.logger.Error().Msg("GetAnswer: getting open AI request")
		return nil, err
	}

	answer := model.Answer{
		Text: response.Choices[0].Message.Content,
	}

	return &answer, nil
}
