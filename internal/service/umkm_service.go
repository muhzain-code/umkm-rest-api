package service

import (
	"fmt"
	"path/filepath"
	"umkm-api/internal/model"
	"umkm-api/internal/repository/filter"
	"umkm-api/internal/repository"
	"umkm-api/internal/request"

	"umkm-api/pkg/response"

	"github.com/google/uuid"
	"os"
)

type UmkmService interface {
	Create(req request.CreateUmkmRequest) (*model.Umkm, error)
	GetAll(page, limit int, filter filter.UmkmFilter) (*PaginatedUmkm, error)
	GetByID(id uuid.UUID) (*model.Umkm, error)
	Update(id uuid.UUID, req request.UpdateUmkmRequest, photoFileName *string) (*model.Umkm, error)
	Delete(id uuid.UUID) error
}

type umkmService struct {
	repo repository.UmkmRepository
}

func NewUmkmService(repo repository.UmkmRepository) UmkmService {
	return &umkmService{repo: repo} 
}

func (s *umkmService) Create(req request.CreateUmkmRequest) (*model.Umkm, error) {
	umkm := model.Umkm{
		ID:           uuid.New(),
		Name:         req.Name,
		OwnerName:    req.OwnerName,
		Nik:          req.Nik,
		Gender:       req.Gender,
		Description:  req.Description,
		PhotoProfile: req.PhotoProfilePath,
		Address:      req.Address,
		Phone:        req.Phone,
		Email:        req.Email,
		WaLink:       req.WaLink,
	}

	if err := s.repo.Create(&umkm); err != nil {
		return nil, err
	}
	return &umkm, nil
}

type PaginatedUmkm struct {
	Data []model.Umkm
	Meta response.Meta
}

func (s *umkmService) GetAll(page, limit int, filter filter.UmkmFilter) (*PaginatedUmkm, error) {
	umkms, total, err := s.repo.FindAll(page, limit, filter)
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

	return &PaginatedUmkm{
		Data: umkms,
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

func (s *umkmService) GetByID(id uuid.UUID) (*model.Umkm, error) {
	return s.repo.FindByID(id)
}

func (s *umkmService) Update(id uuid.UUID, req request.UpdateUmkmRequest, photoFileName *string) (*model.Umkm, error) {
	umkm, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("umkm with id %s not found", id)
	}

	// update field string biasa
	if req.Name != "" {
		umkm.Name = req.Name
	}
	if req.OwnerName != "" {
		umkm.OwnerName = req.OwnerName
	}
	if req.Nik != "" {
		umkm.Nik = req.Nik
	}
	if req.Gender != "" {
		umkm.Gender = req.Gender
	}
	if req.Description != nil {
		umkm.Description = req.Description
	}
	if req.Address != "" {
		umkm.Address = req.Address
	}
	if req.Phone != "" {
		umkm.Phone = req.Phone
	}
	if req.Email != nil {
		umkm.Email = req.Email
	}
	if req.WaLink != "" {
		umkm.WaLink = req.WaLink
	}
	if req.IsActive != nil {
		umkm.IsActive = *req.IsActive
	}

	if photoFileName != nil {
		if umkm.PhotoProfile != nil {
			oldPath := filepath.Join("uploads", *umkm.PhotoProfile)
			_ = os.Remove(oldPath)
		}
		umkm.PhotoProfile = photoFileName
	}

	if err := s.repo.Update(umkm); err != nil {
		return nil, err
	}

	return umkm, nil
}

func (s *umkmService) Delete(id uuid.UUID) error {
	umkm, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("umkm with id %s not found", id)
	}

	if umkm.PhotoProfile != nil {
		filePath := filepath.Join("uploads", *umkm.PhotoProfile)
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove file: %w", err)
		}
	}

	return s.repo.Delete(id)
}
