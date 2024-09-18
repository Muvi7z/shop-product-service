package service

import (
	"context"
	"log/slog"
	"shop-product-service/internal/category/modal"
	storage2 "shop-product-service/internal/category/storage"
)

type Service struct {
	logger  *slog.Logger
	storage storage2.Storage
}

func New(logger *slog.Logger, storage storage2.Storage) *Service {
	return &Service{logger: logger, storage: storage}
}

func (s *Service) AddCategory(ctx context.Context, dto modal.CreateCategoryDto) (int64, error) {
	s.logger.Info("attempting to add category")
	c := modal.Category{
		Name:           dto.Name,
		ParentCategory: dto.ParentCategory,
	}

	_, err := s.storage.Create(ctx, c)
	if err != nil {
		s.logger.Error("failed to create u: %v", err)
		return 0, err
	}

	return c.Id, nil
}

func (s *Service) GetCategoryById(ctx context.Context, id int64) (modal.Category, error) {
	s.logger.Info("attempting to get product")

	p, err := s.storage.FindOne(ctx, id)
	if err != nil {
		s.logger.Error("failed to get product: %v", err)
		return modal.Category{}, err
	}

	return p, nil
}

func (s *Service) GetCategoryByParent(ctx context.Context, id int64) ([]modal.Category, error) {
	s.logger.Info("attempting to get product")

	arrP, err := s.storage.FindByCategory(ctx, id)
	if err != nil {
		s.logger.Error("failed to get product: %v", err)
		return nil, err
	}

	return arrP, nil
}

func (s *Service) GetRootCategories(ctx context.Context) ([]modal.Category, error) {
	s.logger.Info("attempting to get product")

	arrP, err := s.storage.FindRootCategories(ctx)
	if err != nil {
		s.logger.Error("failed to get product: %v", err)
		return nil, err
	}

	return arrP, nil
}
