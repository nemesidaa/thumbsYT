package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	ServerPort       int    `json:"port"`
	ServerHost       string `json:"host"`
	Resolution       string `json:"resolution"`
	DBName           string `json:"dbname"`
	DefaultDBTimeout int    `json:"default_db_timeout"`
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
		return err
	}
	defer file.Close()

	// Создаём декодер JSON
	decoder := json.NewDecoder(file)

	// Декодируем JSON из файла
	err = decoder.Decode(s)
	if err != nil {

		return err
	}
	return nil
}
