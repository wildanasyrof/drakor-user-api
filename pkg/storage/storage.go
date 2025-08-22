package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/config"
)

type LocalStorage interface {
	Save(file *multipart.FileHeader) (string, error)
}

type localStorage struct {
	cfg *config.Config
}

func NewLocalStorage(cfg *config.Config) *localStorage {
	return &localStorage{
		cfg: cfg,
	}
}

func (s *localStorage) Save(file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll(s.cfg.Server.UploadDir, 0755); err != nil {
		return "", err
	}
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	path := filepath.Join(s.cfg.Server.UploadDir, filename)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}
	return filename, nil
}

func PublicURL(basePath, filename string) string {
	return fmt.Sprintf("%s/%s", basePath, filename)
}
