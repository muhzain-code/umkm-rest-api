package service

import (
	"fmt"

	"umkm-api/internal/model"
	"umkm-api/internal/repository"
	"umkm-api/internal/repository/filter"
	"umkm-api/internal/request"
	"umkm-api/pkg/response"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(req request.CreateProductRequest) (*model.Product, error)
	GetAll(page, limit int, filter filter.ProductFilter) (*PaginateProduct, error)
	GetByID(id uuid.UUID) (*model.Product, error)
	Update(id uuid.UUID, req request.UpdateProductRequest) (*model.Product, error)
	Delete(id uuid.UUID) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(req request.CreateProductRequest) (*model.Product, error) {
	for i, p := range req.Photos {
		fmt.Printf("Photo[%d]: %+v\n", i, p)
	}
	for i, m := range req.Marketplaces {
		fmt.Printf("Marketplace[%d]: %+v\n", i, m)
	}

	parseID, err := uuid.Parse(req.UmkmID)
	if err != nil {
		return nil, fmt.Errorf("invalid umkm_id '%s': %w", req.UmkmID, err)
	}

	product := model.Product{
		ID:          uuid.New(),
		UmkmID:      parseID,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Status:      req.Status,
	}

	for _, p := range req.Photos {
		isActive := true
		if p.IsActive != nil {
			isActive = *p.IsActive
		}

		photo := model.ProductPhoto{
			PhotoType:   p.PhotoType,
			FilePath:    p.FilePath,
			IsActive:    isActive,
			Description: p.Description,
		}
		product.Photos = append(product.Photos, photo)
	}

	for _, m := range req.Marketplaces {
		isActive := true
		if m.IsActive != nil {
			isActive = *m.IsActive
		}

		marketplace := model.Marketplace{
			Name:            m.Name,
			Price:           m.Price,
			MarketplaceLink: m.MarketplaceLink,
			IsActive:        isActive,
		}
		product.Marketplaces = append(product.Marketplaces, marketplace)
	}

	if err := s.repo.Create(&product); err != nil {
		return nil, fmt.Errorf("cannot create product: %w", err)
	}

	return &product, nil
}

type PaginateProduct struct {
	Data []model.Product
	Meta response.Meta
}

func (s *productService) GetAll(page, limit int, filter filter.ProductFilter) (*PaginateProduct, error) {
	products, total, err := s.repo.FindAll(page, limit, filter)
	if err != nil {
		return nil, err
	}

	lastPage := int((total + int64(limit) - 1) / int64(limit))
	var from, to int
	if total == 0 {
		from, to = 0, 0
	} else {
		from = (page-1)*limit + 1
		to = page * limit
		if int64(to) > total {
			to = int(total)
		}
	}

	return &PaginateProduct{
		Data: products,
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

func (s *productService) GetByID(id uuid.UUID) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Update(id uuid.UUID, req request.UpdateProductRequest) (*model.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("product with id %s not found", id.String())
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Status = req.Status
	product.CategoryID = req.CategoryID

	var photos []*model.ProductPhoto
	for _, p := range req.Photos {
		isActive := true
		if p.IsActive != nil {
			isActive = *p.IsActive
		}

		desc := ""
		if p.Description != nil {
			desc = *p.Description
		}

		photos = append(photos, &model.ProductPhoto{
			ProductID:   product.ID,
			PhotoType:   p.PhotoType,
			FilePath:    p.FilePath,
			IsActive:    isActive,
			Description: &desc,
		})
	}

	var marketplaces []*model.Marketplace
	for _, m := range req.Marketplaces {
		isActive := true
		if m.IsActive != nil {
			isActive = *m.IsActive
		}

		marketplaces = append(marketplaces, &model.Marketplace{
			ProductID:       product.ID,
			Name:            m.Name,
			Price:           m.Price,
			MarketplaceLink: m.MarketplaceLink,
			IsActive:        isActive,
		})
	}

	if err := s.repo.Update(product, photos, marketplaces); err != nil {
		return nil, fmt.Errorf("cannot update product: %w", err)
	}

	updatedProduct, err := s.repo.FindByID(product.ID)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (s *productService) Delete(id uuid.UUID) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("product with id %s not found", id.String())
	}
	return s.repo.Delete(id)
}
