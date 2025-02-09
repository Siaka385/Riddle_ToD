package utils

import "github.com/google/uuid"

func GenerateRandomKey() string {
	key := uuid.New()
	return key.String()
}

