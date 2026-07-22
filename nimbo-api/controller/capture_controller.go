package controller

import (
	"net/http"
	"nimbo-api/useCase"

	"github.com/gin-gonic/gin"
)

type CaptureController struct {
	captureUseCase useCase.CaptureUsecase
}

func NewCaptureController(usecase useCase.CaptureUsecase) CaptureController {
	return CaptureController{
		captureUseCase: usecase,
	}
}

func (p *CaptureController) GetCapture(c *gin.Context) {
	captures, err := p.captureUseCase.GetCapture()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, captures)
}
