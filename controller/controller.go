package controller

import (
	"net/http"

	"github.com/varopxndx/clone-ai-be/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Usecase contains the usecase methods
type Usecase interface {
	GetSample() (*model.SampleResponse, error)
	GetAnswer(message string) (*model.Answer, error)
}

// Controller structure
type Controller struct {
	usecase Usecase
	logger  zerolog.Logger
}

// New creates a controller
func New(usecase Usecase, logger zerolog.Logger) *Controller {
	return &Controller{
		usecase: usecase,
		logger:  logger,
	}
}

// Ping checks if the application is up and running
func (c *Controller) Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "service healthy")
}

// GetSample gets sample data controller
func (c *Controller) GetSample(g *gin.Context) {
	// call usecase layer

	response, err := c.usecase.GetSample()
	if err != nil {
		c.logger.Error().Msgf("error getting sample: %s", err.Error())
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK, response)
}

// GetAnswer Request a response from Open AI
func (c *Controller) GetAnswer(g *gin.Context) {
	// call usecase layer
	answerRequest := model.AnswerRequest{}
	err := g.ShouldBind(&answerRequest)
	if err != nil {
		c.logger.Error().Msg(err.Error())
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	response, err := c.usecase.GetAnswer(answerRequest.Message)
	if err != nil {
		c.logger.Error().Msgf("error getting sample: %s", err.Error())
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK, response)
}
