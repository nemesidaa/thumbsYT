package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ClientConfig struct {
	ServerPort int    `json:"port"`
	ServerHost string `json:"host"`
	Resolution string `json:"resolution"`
	LogLevel   string `json:"log_level"`
	Timeout    int    `json:"timeout"`
}

func NewConfig() *ClientConfig {
	return &ClientConfig{}
}

func (s *ClientConfig) ParseFlags(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error handling file:", err)
		return err
	}
	defer file.Close()

	// Создаём декодер JSON
	decoder := json.NewDecoder(file)

	// Декодируем JSON из файла
	err = decoder.Decode(s)
	if err != nil {
		fmt.Println("error decoding:", err)
		return err
	}
	return nil
}
