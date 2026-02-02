package pkg

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrInvalidFileType = errors.New("invalid file type")
	ErrFileTooLarge    = errors.New("file size exceeds limit")
)

const MaxUploadSize = 2 * 1024 * 1024 // 2MB

func SaveMultipartFile(file *multipart.FileHeader, dst string, allowedExt []string) error {
	// check filesize
	if file.Size > MaxUploadSize {
		return ErrFileTooLarge
	}

	// check extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ! extContains(allowedExt, ext) {
		return ErrInvalidFileType
	}
	
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func extContains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
