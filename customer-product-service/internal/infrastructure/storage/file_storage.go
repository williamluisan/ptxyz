package storage

import (
	"context"
	"mime/multipart"
)

type FileStorage interface {
	Save(ctx context.Context, path string, file *multipart.FileHeader) (string, error)
}