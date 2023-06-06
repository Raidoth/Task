package config

import (
	"encoding/json"
	"log"
	"os"
)

type ServiceCfg struct {
	UserMysql     string `json:"usermysql" validate:"required"`
	PasswordMysql string `json:"passwordmysql" validate:"required"`
	AddressMysql  string `json:"addressmysql" validate:"required"`
	DbMysql       string `json:"dbmysql" validate:"required"`
	PortGrpc      string `json:"portgrpc" validate:"required"`
}

// Initializing configuration
func InitService(path string) ServiceCfg {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error read config: %s", err)
	}
	service := ServiceCfg{}
	//can use easyjson
	if err = json.Unmarshal(data, &service); err != nil {
		log.Fatalf("Error load config: %s", err)
	}
	log.Printf("Config loaded\n")
	return service
}
