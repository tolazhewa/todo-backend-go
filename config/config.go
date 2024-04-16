package config

import (
	"log"
	"os"
	"sync"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type AppConfig struct {
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"pass"`
	DbName   string `koanf:"dbname"`
}

var (
	appConfig   *AppConfig
	mu          sync.Mutex
	initialized bool
)

func InitAppConfig() {
	mu.Lock()
	defer mu.Unlock()

	var configFilePath string
	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "local":
		configFilePath = "config/local.json"
	case "dev":
		configFilePath = "config/dev.json"
	case "prod":
		configFilePath = "config/prod.json"
	default:
		log.Fatal("environment type is not declared correctly in the PROJECT_BASE/.env file. Acceptable values are local/dev/prod")
	}
	k := koanf.New(".")
	if err := k.Load(file.Provider(configFilePath), json.Parser()); err != nil {
		log.Fatalf("error loading config for env=%s: %v", env, err)
	}
	k.Unmarshal("", &appConfig)
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Fatal("database password is not declared correctly in the PROJECT_BASE/.env file")
	}
	appConfig.Password = dbPass
	initialized = true

	log.Println("gathered app configs")
}

func GetAppConfig() *AppConfig {
	if !initialized || appConfig == nil {
		InitAppConfig()
	}
	return appConfig
}
