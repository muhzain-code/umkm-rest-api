package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"umkm-api/internal/event/model"
	"umkm-api/internal/event/repository"
	"umkm-api/internal/event/request"
	"umkm-api/pkg/response"

	"github.com/google/uuid"
)

type EventService interface {
	GetAll(page, limit int, filter repository.EventFilter) (*PaginateEvent, error)
	GetByID(id int) (*model.Event, error)
	CreateEvent(req request.CreateEventRequest) (*model.Event, error)
	UpdateEvent(id int, req request.UpdateEventRequest) (*model.Event, error)
	DeleteEvent(id int) error
}

type eventService struct {
	repo repository.EventRepository
}
type PaginateEvent struct {
	Data []model.EventResponse `json:"data"`
	Meta response.Meta         `json:"meta"`
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) GetAll(page, limit int, filter repository.EventFilter) (*PaginateEvent, error) {
	events, total, err := s.repo.FindAll(page, limit, filter)
	if err != nil {
		return nil, err
	}

	// mapping ke response struct
	var eventResponses []model.EventResponse
	for _, e := range events {
		var umkms []model.EventUmkmResponse
		for _, eu := range e.EventUmkms {
			umkms = append(umkms, model.EventUmkmResponse{
				UmkmID:   eu.UmkmID,
				IsActive: eu.IsActive,
			})
		}
		eventResponses = append(eventResponses, model.EventResponse{
			ID:          e.ID,
			Name:        e.Name,
			Description: e.Description,
			Photo:       e.Photo,
			StartDate:   e.StartDate,
			EndDate:     e.EndDate,
			IsActive:    e.IsActive,
			EventUmkms:  umkms,
		})
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

	return &PaginateEvent{
		Data: eventResponses,
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

func (e *eventService) GetByID(id int) (*model.Event, error) {
	return e.repo.FindByID(id)
}

func (s *eventService) CreateEvent(req request.CreateEventRequest) (*model.Event, error) {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %w", err)
	}

	// Convert umkm_id dari string ke UUID
	umkmUUID, err := uuid.Parse(req.UmkmID)
	if err != nil {
		return nil, fmt.Errorf("invalid umkm_id: %w", err)
	}

	// Buat struct Event
	event := model.Event{
		Name:        req.Name,
		Description: req.Description,
		Photo:       req.PhotoPath,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := s.repo.Create(&event, []uuid.UUID{umkmUUID}); err != nil {
		return nil, err
	}
	return &event, nil
}

func (i *eventService) UpdateEvent(id int, req request.UpdateEventRequest) (*model.Event, error) {
	event, err := i.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("event with id %d not found", id)
	}

	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date: %w", err)
		}
		event.StartDate = startDate
	}

	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end date: %w", err)
		}
		event.EndDate = endDate
	}

	if req.Name != "" {
		event.Name = req.Name
	}

	if req.Description != "" {
		event.Description = req.Description
	}

	if req.PhotoPath != nil {
		event.Photo = req.PhotoPath
	}

	if req.IsActive != nil {
		event.IsActive = *req.IsActive
	}

	if err := i.repo.Update(event); err != nil {
		return nil, err
	}

	return event, nil
}

func (x *eventService) DeleteEvent(id int) error {
	event, err := x.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("event with id %d not found", id)
	}

	if event.Photo != nil {
		filePath := filepath.Join("uploads", *event.Photo)
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove file: %w", err)
		}
	}

	return x.repo.Delete(id)
}
