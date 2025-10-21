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
func NewUuidShort() string {
	uuid := uuid.New()
	fmt.Println("this is update by main")
	return uuid.String()[:8]
}
