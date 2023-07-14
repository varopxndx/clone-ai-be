package controller

import (
	"context"
	"net/http"

	"github.com/varopxndx/clone-ai-be/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Usecase contains the usecase methods
type Usecase interface {
	GetSample() (*model.SampleResponse, error)
	GetAnswer(ctx context.Context, message string) (*model.Answer, error)
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

// Ping godoc
// @Summary Checks if the application is up and running
// @Accept json
// @Success 200
// @Failure 400
// @Router /ping [get]
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
// @Summary Sends a question to Open AI and receives a response
// @Accept json
// @Success 200 {object} model.Answer
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /v1/get-answer [post]
func (c *Controller) GetAnswer(g *gin.Context) {
	// call usecase layer
	answerRequest := model.AnswerRequest{}
	err := g.ShouldBind(&answerRequest)
	if err != nil {
		c.logger.Error().Msg(err.Error())
		g.AbortWithStatus(http.StatusBadRequest)
		return
	}
	response, err := c.usecase.GetAnswer(g.Request.Context(), answerRequest.Message)
	if err != nil {
		c.logger.Error().Msgf("error getting sample: %s", err.Error())
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	g.JSON(http.StatusOK, response)
}
