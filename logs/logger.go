package logs

import (
	"log"
	"os"
)

func InitLogger() (*log.Logger, error) {
	file, err := os.OpenFile("logs/backup.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	logger := log.New(file, "backup: ", log.LstdFlags)
	return logger, nil
}
