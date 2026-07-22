package useCase

import (
	"nimbo-api/model"
	"nimbo-api/repository"
)

type CaptureUsecase struct {
	repository repository.CaptureRepository
}

func NewCaptureUseCase(repo repository.CaptureRepository) CaptureUsecase {
	return CaptureUsecase{
		repository: repo,
	}
}

func (cu *CaptureUsecase) GetCapture() ([]model.Capture, error) {
	return cu.repository.GetCapture()
}
