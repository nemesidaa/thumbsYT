package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"
)

var errWTF error = errors.New("something went wrong") // 0_0 костылёчек)))

func (c *Client) HandleLoad(id, resolution string) error {

	if _, ok := AllAvailableResolutions[Resolution(resolution)]; !ok {
		resolution = c.Resolution
	}

	log.Println("Creating deadline...")
	// Определяем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	defer cancel()

	// Отправляем запрос
	r, err := c.Load(ctx, &pb.LoadRequest{ServiceID: c.ServiceID, VideoID: id, Resolution: resolution})
	if err != nil {
		return err
	}

	log.Printf("Received RawData of length: %d\n", len(r.RawData))

	fileName := fmt.Sprintf("%s.jpg", id)

	// Создаем путь к файлу
	filePath := filepath.Join("thumbs/", fileName)

	// Проверяем и создаем директорию, если её нет
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Создаем файл
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копируем тело ответа в файл
	if _, err := io.Copy(file, bytes.NewReader([]byte(r.RawData))); err != nil {
		return err
	}

	if fileName != "nil" {
		log.Printf("Thumbnail saved as %s\n", fileName)
	} else {
		return errWTF
	}

	return nil
}
