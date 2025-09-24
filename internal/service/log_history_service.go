package service

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"umkm-api/internal/model"
	"umkm-api/internal/repository"
	"umkm-api/internal/request"
	"github.com/google/uuid"
)

type LogHistoryService interface {
	CreateAsync(req request.LogHistoryRequest)
	ValidateProduct(productID uuid.UUID) (uuid.UUID, error)
}

type logHistoryService struct {
	repo repository.LogHistoryRepository
}

func NewLogHistoryService(repo repository.LogHistoryRepository) LogHistoryService {
	return &logHistoryService{repo: repo}
}

func (s *logHistoryService) ValidateProduct(productID uuid.UUID) (uuid.UUID, error) {
    return s.repo.FindUmkmByProduct(productID)
}

func (s *logHistoryService) CreateAsync(req request.LogHistoryRequest) {
	go func() {
		UmkmID, err := s.repo.FindUmkmByProduct(req.ProductID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("[LogHistoryService] Product not found: %v\n", req.ProductID)
			} else {
				fmt.Printf("[LogHistoryService] Error finding UMKM by product: %v\n", err)
			}
			return
		}

		fmt.Println(req.ProductID)
		fmt.Println(UmkmID)

		const prefix = "PTR"

		applicationStr := fmt.Sprintf("%02d", req.ApplicationID)
		umkmStr := UmkmID.String()[0:2]
		tanggal := time.Now().Format("02")

		fixedLength := len(prefix) + len(applicationStr) + len(umkmStr) + len(tanggal)
		randomLen := 17 - fixedLength
		if randomLen < 1 {
			randomLen = 1
		}

		var resi string
		for {
			var sb strings.Builder
			for i := 0; i < randomLen; i++ {
				sb.WriteString(strconv.Itoa(rand.Intn(10)))
			}
			resi = prefix + applicationStr + umkmStr + tanggal + sb.String()

			count, err := s.repo.CountByResi(resi)
			if err != nil {
				fmt.Printf("[LogHistoryService] Error checking resi: %v\n", err)
				return
			}
			if count == 0 {
				break
			}
		}

		log := &model.LogHistory{
			ApplicationID: req.ApplicationID,
			UmkmID:        UmkmID,
			Resi:          resi,
			BuyerName:     req.BuyerName,
			BuyerPhone:    req.BuyerPhone,
			BuyerAddress:  req.BuyerAddress,
			IPAddress:     req.IPAddress,
			UserAgent:     req.UserAgent,
		}

		if err := s.repo.Create(log); err != nil {
			fmt.Printf("[LogHistoryService] Failed to save log: %v\n", err)
		}
	}()
}
