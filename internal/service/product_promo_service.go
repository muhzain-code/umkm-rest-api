package service

import (
	"umkm-api/internal/model"
	"umkm-api/internal/repository"
	"umkm-api/internal/repository/filter"
	"umkm-api/internal/request"
	"umkm-api/pkg/response"
)

type ProductPromoService interface {
	Create(req request.CreateProductPromoRequest) (*model.ProductPromo, error)
	GetAll(page, limit int, filter filter.ProductPromoFilter) (*PaginateProductPromo, error)
	GetByID(id int64) (*model.ProductPromo, error)
	Update(id int64, req request.UpdateProductPromoRequest) (*model.ProductPromo, error)
	Delete(id int64) error
}

type productPromoService struct {
	repo repository.ProductPromoRepository
}

func NewProductPromoService(repo repository.ProductPromoRepository) ProductPromoService {
	return &productPromoService{repo: repo}
}

func (s *productPromoService) Create(req request.CreateProductPromoRequest) (*model.ProductPromo, error) {
	productPromo := model.ProductPromo{
		EventID:         req.EventID,
		ProductID:       req.ProductID,
		PromoPrice:      req.PromoPrice,
		DiscountAmount:  req.DiscountAmount,
		DiscountPercent: req.DiscountPercent,
		IsActive:        *req.IsActive,
		Terms:           req.Terms,
	}

	if err := s.repo.Create(&productPromo); err != nil {
		return nil, err
	}

	return &productPromo, nil

}

type PaginateProductPromo struct {
	Data []model.ProductPromo `json:"data"`
	Meta response.Meta        `json:"meta"`
}

func (s *productPromoService) GetAll(page, limit int, filter filter.ProductPromoFilter) (*PaginateProductPromo, error) {
	productPromo, total, err := s.repo.FindAll(page, limit, filter)

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

	return &PaginateProductPromo{
		Data: productPromo,
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

func (s *productPromoService) GetByID(id int64) (*model.ProductPromo, error) {
	return s.repo.FindByID(id)
}

func (s *productPromoService) Update(id int64, req request.UpdateProductPromoRequest) (*model.ProductPromo, error) {
	productPromo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	productPromo.EventID = req.EventID
	productPromo.ProductID = req.ProductID
	productPromo.PromoPrice = req.PromoPrice
	productPromo.DiscountAmount = req.DiscountAmount
	productPromo.DiscountPercent = req.DiscountPercent
	productPromo.IsActive = *req.IsActive

	if err := s.repo.Update(productPromo); err != nil {
		return nil, err
	}

	return productPromo, nil
}

func (s *productPromoService) Delete(id int64) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
