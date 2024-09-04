package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	ServerPort int    `json:"port"`
	ServerHost string `json:"host"`
	Resolution string `json:"resolution"`
	DBName     string `json:"dbname"`
	// // resolutionArray???
	// BrokerCapacity          int
	// MaxBrokerRetriesCounter int8
	// IdealCaching            bool
}

func NewConfig() *ServerConfig {
	return &ServerConfig{}
}

func (s *ServerConfig) ParseFlags(fileName string) error {
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
