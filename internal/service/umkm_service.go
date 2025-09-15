package service

import (
	"fmt"
	"umkm-api/internal/model"
	"umkm-api/internal/repository"
	"umkm-api/internal/request"

	"github.com/google/uuid"
)

type UmkmService interface {
	Create(req request.CreateUmkmRequest) (*model.Umkm, error)
	GetAll() ([]model.Umkm, error)
	GetByID(id uuid.UUID) (*model.Umkm, error)
	Update(id uuid.UUID, req request.UpdateUmkmRequest) (*model.Umkm, error)
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
		ID: uuid.New(),
		Name: req.Name,	
		OwnerName: req.OwnerName,
		Nik: req.Nik,
		Gender: req.Gender,
		Description: req.Description,
		PhotoProfile: req.PhotoProfile,
		Address: req.Address,
		Phone: req.Phone,
		Email: req.Email,
		WaLink: req.WaLink,
	}

	if err := s.repo.Create(&umkm); err != nil {
		return nil, err
	}
	return &umkm, nil
}

func (s *umkmService) GetAll() ([]model.Umkm, error) {
	return s.repo.FindAll()
}

func (s *umkmService) GetByID(id uuid.UUID) (*model.Umkm, error) {
	return s.repo.FindByID(id)
}

func (s *umkmService) Update(id uuid.UUID, req request.UpdateUmkmRequest) (*model.Umkm, error) {
	umkm, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("umkm with id %s not found", id)
	}

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
	if req.PhotoProfile != nil {
		umkm.PhotoProfile = req.PhotoProfile
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
	// Kenapa pakai *req.IsAcive? karena untuk cek apakah uses benar" mengirim true/false, karena bisa saja user 
	if req.IsActive != nil {
		umkm.IsActive = *req.IsActive
	}

	if err := s.repo.Update(umkm); err != nil {
		return nil, err
	}

	return umkm, nil
}

func (s *umkmService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
