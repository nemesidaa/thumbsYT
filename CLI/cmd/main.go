package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/nemesidaa/thumbsYT/proto/gen/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := grpc.NewClient("localhost:5252", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMainstreamClient(conn)

	// Определяем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Отправляем запрос
	r, err := c.Load(ctx, &pb.LoadRequest{RequestID: "test", VideoID: "dQw4w9WgXcQ", Resolution: "hqdefault"})
	if err != nil {
		log.Fatalf("could not handle: %v", err)
	}

	// Отправляем запрос

	if err != nil {
		log.Fatalf("could not handle: %v", err)
	}
	fileName := fmt.Sprintf("%s.jpg", "test")

	// Создаем путь к файлу
	filePath := filepath.Join("tests/", fileName)

	// Проверяем и создаем директорию, если её нет
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Создаем файл
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Копируем тело ответа в файл
	if _, err := io.Copy(file, bytes.NewReader([]byte(r.RawData))); err != nil {
		log.Fatal(err)
	}

	log.Printf("Thumbnail saved as %s\n", fileName)

	log.Print(r.String())
}
