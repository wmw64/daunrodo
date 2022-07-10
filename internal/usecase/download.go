package usecase

import (
	"context"
	"daunrodo/internal/entity"
	"daunrodo/pkg/utils"
	"fmt"
)

type DownloadService interface {
	Download(ctx context.Context, link string) ([]entity.File, error)
}

type DownloadUseCase struct {
	service DownloadService
}

func New(d DownloadService) *DownloadUseCase {
	return &DownloadUseCase{
		service: d,
	}
}

func (u *DownloadUseCase) Download(ctx context.Context, link string) ([]entity.File, error) {

	if !utils.IsURLValid(link) {
		link = utils.FixURL(link)
	}

	files, err := u.service.Download(ctx, link)
	if err != nil {
		return []entity.File{}, fmt.Errorf("DownloadUseCase - Get - u.repo.Get: %w", err)
	}

	return files, nil
}
