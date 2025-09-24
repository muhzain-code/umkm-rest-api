package service

import (
	"umkm-api/internal/model"
	"umkm-api/internal/repository"
	"umkm-api/internal/request"
	"umkm-api/pkg/response"
)

type ApplicationService interface {
	Create(request request.CreateApplicationRequest) (*model.Application, error)
	GetAll(page, limit int, name string) (*PaginateApplication, error)
	GetByID(id int64) (*model.Application, error)
	Update(id int64, req request.UpdateApplicationRequest) (*model.Application, error)
	Delete(id int64) error
}

type ApplicationServiceImpl struct {
	repo repository.ApplicationRepository
}

func NewApplicationService(repo repository.ApplicationRepository) ApplicationService {
	return &ApplicationServiceImpl{repo: repo}
}

func (s *ApplicationServiceImpl) Create(req request.CreateApplicationRequest) (*model.Application, error) {
	app := model.Application{
		Name: req.Name,
	}

	if err := s.repo.Create(&app); err != nil {
		return nil, err
	}

	return &app, nil
}

type PaginateApplication struct {
	Data []model.Application `json:"data"`
	Meta response.Meta       `json:"meta"`
}

func (s *ApplicationServiceImpl) GetAll(page, limit int, name string) (*PaginateApplication, error) {
	apps, total, err := s.repo.FindAll(page, limit, name)
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

	return &PaginateApplication{
		Data: apps,
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

func (s *ApplicationServiceImpl) GetByID(id int64) (*model.Application, error) {
	return s.repo.FindByID(id)
}

func (s *ApplicationServiceImpl) Update(id int64, req request.UpdateApplicationRequest) (*model.Application, error) {

	app, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	app.Name = req.Name

	if err := s.repo.Update(app); err != nil {
		return nil, err
	}

	return app, nil
}

func (s *ApplicationServiceImpl) Delete(id int64) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}