package service

import (
	"context"
	batikRepo "kenalbatik-be/internal/batik/repository"
	"kenalbatik-be/internal/domain"
	"time"
)

type BatikService interface {
	GetAllBatik(ctx context.Context, batikParam domain.BatikParams) ([]domain.BatikResponse, error)
	GetBatikByID(ctx context.Context, batikID int) (domain.BatikResponse, error)
}

type batikService struct {
	batikRepository batikRepo.BatikRepository
}

func NewBatikService(batikRepo batikRepo.BatikRepository) BatikService {
	return &batikService{
		batikRepository: batikRepo,
	}
}

func (s *batikService) GetAllBatik(ctx context.Context, batikParam domain.BatikParams) ([]domain.BatikResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var batiks []domain.Batik

	err := s.batikRepository.FindAll(ctx, &batiks, batikParam)

	var batikResponses []domain.BatikResponse
	for _, batik := range batiks {
		batikResponse := domain.BatikResponse{
			ID:          batik.ID,
			Name:        batik.Name,
			JavaName:    batik.JavaName,
			City:        batik.City,
			Province:    batik.Province.Name,
			Island:      batik.Island.Name,
			Description: batik.Description,
			Link_Image:  batik.Link_Image,
		}
		batikResponses = append(batikResponses, batikResponse)
	}

	select {
	case <-ctx.Done():
		return batikResponses, domain.ErrTimeout
	default:
		return batikResponses, err
	}
}

func (s *batikService) GetBatikByID(ctx context.Context, batikId int) (domain.BatikResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var batik domain.Batik
	err := s.batikRepository.FindByID(ctx, &batik, batikId)
	if err != nil {
		return domain.BatikResponse{}, err
	}

	batikResponse := domain.BatikResponse{
		ID:          batik.ID,
		Name:        batik.Name,
		JavaName:    batik.JavaName,
		Province:    batik.Province.Name,
		Island:      batik.Island.Name,
		City:        batik.City,
		Description: batik.Description,
		Link_Image:  batik.Link_Image,
	}

	select {
	case <-ctx.Done():
		return batikResponse, domain.ErrTimeout
	default:
		return batikResponse, err
	}
}
