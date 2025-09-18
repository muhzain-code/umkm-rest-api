package utils

import (
	"path/filepath"
	"github.com/go-playground/validator/v10"
	"strings"
	"mime/multipart"
	"errors"
	"slices"
)

var Validate = validator.New()

func ValidatePhoto(photo *multipart.FileHeader, required bool) error {
	if photo == nil {
		if required {
			return errors.New("photo profile required")
		}
		return nil
	}

	ext := strings.ToLower(filepath.Ext(photo.Filename))
	allowed := []string{".jpg", ".png", ".jpeg"}
	if !slices.Contains(allowed, ext) {
		return errors.New("invalid photo format")
	}
	return nil
}