package service

import (
	"fmt"
	"os"
	"path/filepath"
	"umkm-api/internal/category/model"
	"umkm-api/internal/category/repository"
	"umkm-api/internal/category/request"
	"umkm-api/pkg/response"
)

type CategoryService interface {
	Create(req request.CreateCategoryRequest) (*model.Category, error)
	GetAll(page, limit int, filter repository.CategoryFilter) (*PaginateCategory, error)
	GetByID(id int64) (*model.Category, error)
	Update(id int64, req request.UpdateCategoryRequest, photoProfileName *string) (*model.Category, error)
	Delete(id int64) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(req request.CreateCategoryRequest) (*model.Category, error) {
	category := model.Category{
		Name:  req.Name,
		Photo: req.PhotoPath,
	}

	if err := s.repo.Create(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

type PaginateCategory struct {
	Data []model.Category
	Meta response.Meta
}

func (s *categoryService) GetAll(page, limit int, filter repository.CategoryFilter) (*PaginateCategory, error) {
	categories, total, err := s.repo.FindAll(page, limit, filter)
	if err != nil {
		return nil, err
	}

	lastPage := int((total + int64(limit) - 1) / int64(limit))
	var from, to int
	if total == 0 {
		from = 0
		to = 0
	} else {
		from = (page-1)*limit + 1
		to = page * limit
		if int64(to) > total {
			to = int(total)
		}
	}

	return &PaginateCategory{
		Data: categories,
		Meta: response.Meta{
			CurrentPage: page,
			PerPage:     limit,
			Total:       int(total),
			LastPage:    lastPage,
			From:        from,
			To:          to,
		},
	}, nil
}

func (s *categoryService) GetByID(id int64) (*model.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) Update(id int64, req request.UpdateCategoryRequest, photoFileName *string) (*model.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("category with id %d not found", id)
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if photoFileName != nil {
		if category.Photo != nil {
			oldPath := filepath.Join("uploads", *category.Photo)
			_ = os.Remove(oldPath)
		}

		category.Photo = photoFileName
	}

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Delete(id int64) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if category.Photo != nil {
		filepath := filepath.Join("uploads", *category.Photo)
		if err := os.Remove(filepath); err != nil {
			return err
		}
	}

	return s.repo.Delete(id)
}
