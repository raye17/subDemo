package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func Say() {
	fmt.Println("this is utils")
}
func NewUuid() string {
	uuid := uuid.New()
	return uuid.String()
}
