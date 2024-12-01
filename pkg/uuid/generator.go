package uuid

import (
	"github.com/google/uuid"
    "strings"
)

func GenerateUUID() string {
    return uuid.New().String()
}

func GenerateShortUUID() string {
    fullUUID := uuid.New().String()
    return strings.ReplaceAll(fullUUID, "-", "")[:8]
}