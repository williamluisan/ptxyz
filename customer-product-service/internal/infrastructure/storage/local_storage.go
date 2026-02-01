package storage

import (
	"context"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	helper "ptxyz/customer-product-service/pkg"
)

type LocalStorage struct {
	BasePath string
	AllowedExt []string
}

func (l *LocalStorage) Save(ctx context.Context, path string, file *multipart.FileHeader) (string, error) {
	l.BasePath = viper.GetString("UPLOAD_BASE_PATH")

	fullPath := filepath.Join(l.BasePath, path)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return "", err
	}

	return fullPath, helper.SaveMultipartFile(file, fullPath, l.AllowedExt)
}