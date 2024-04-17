package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/tolazhewa/todo-backend-go/config"
)

var (
	instance    *sql.DB
	mu          sync.Mutex
	initialized bool
)

func InitDB() *sql.DB {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil || !initialized {
		var err error
		config := config.GetAppConfig()
		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.DbName)
		instance, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("failed to connect to database: %v", err)
			return nil
		}

		if err = instance.Ping(); err != nil {
			log.Printf("failed to ping database: %v\n", err)
			return nil
		}
		log.Println("database connection established")
		initialized = true
	}
	return instance
}

func GetDB() *sql.DB {
	if !initialized {
		return InitDB()
	}
	return instance
}
