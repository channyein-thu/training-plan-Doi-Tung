package helper

import (
	"io"
	"os"
	"path/filepath"
)

type Storage interface {
	Upload(
		objectPath string,
		file io.Reader,
		contentType string,
	) (string, error)

	Delete(objectPath string) error
}


type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{BasePath: basePath}
}

func (s *LocalStorage) Upload(
	objectPath string,
	file io.Reader,
	_ string,
) (string, error) {

	fullPath := filepath.Join(s.BasePath, objectPath)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return "", err
	}

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	// Return URL-like path
	return "/" + fullPath, nil
}

func (s *LocalStorage) Delete(objectPath string) error {
	fullPath := filepath.Join(s.BasePath, objectPath)
	return os.Remove(fullPath)
}
