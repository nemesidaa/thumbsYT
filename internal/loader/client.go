package loader

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type LoaderDone struct{}

func (l *Loader) Load(id, resolution string, ctx context.Context) (context.Context, error) {
	// Forming URL
	url := fmt.Sprintf("https://img.youtube.com/vi/%s/%s.jpg", id, resolution)

	// Sending Get request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileName := fmt.Sprintf("%s.jpg", id)

	// Создаем путь к файлу
	filePath := filepath.Join(l.SubDir, fileName)

	// Проверяем и создаем директорию, если её нет
	if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Создаем файл
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Copying response body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Thumbnail saved as %s\n", fileName)
	return context.WithValue(ctx, LoaderDone{}, id), nil
}
