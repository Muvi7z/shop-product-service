package service

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"shop-product-service/internal/product/modal"
	storage2 "shop-product-service/internal/product/storage"
)

type Service struct {
	logger  *slog.Logger
	storage storage2.Storage
}

func New(logger *slog.Logger, storage storage2.Storage) *Service {
	return &Service{logger: logger, storage: storage}
}

func (s *Service) AddProduct(ctx context.Context, dto modal.CreateProductDTO) (string, error) {
	s.logger.Info("attempting to add product")
	pUuid := uuid.New()
	p := modal.Product{
		Uuid:        pUuid.String(),
		Name:        dto.Name,
		Price:       dto.Price,
		Count:       dto.Count,
		Image:       dto.Image,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
	}

	_, err := s.storage.Create(ctx, p)
	if err != nil {
		s.logger.Error("failed to create u: %v", err)
		return "", err
	}

	return pUuid.String(), nil
}

func (s *Service) GetProductById(ctx context.Context, id string) (modal.Product, error) {
	s.logger.Info("attempting to get product")

	p, err := s.storage.FindOne(ctx, id)
	if err != nil {
		s.logger.Error("failed to get product: %v", err)
		return modal.Product{}, err
	}

	return p, nil
}

func (s *Service) GetProductByCategory(ctx context.Context, id int64) ([]modal.Product, error) {
	s.logger.Info("attempting to get product")

	arrP, err := s.storage.FindByCategory(ctx, id)
	if err != nil {
		s.logger.Error("failed to get product: %v", err)
		return nil, err
	}

	return arrP, nil
}
