package models

import (
	"time"

	"github.com/google/uuid"
)

func GenerateId() string {
	return uuid.New().String()
}

func GenerateTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
