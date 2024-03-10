package uuid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

/*
New is a function that returns a UUID with a prefix.
Eg. abc_1234567890abcdef1234567890abcdef
*/
func New(prefix string) (string, error) {
	if len(prefix) != 3 {
		return "", fmt.Errorf("prefix must be 3 characters long")
	}

	str := uuid.New().String()
	clean := strings.ReplaceAll(str, "-", "")

	uuid := fmt.Sprintf("%s_%s", prefix, clean)

	return uuid, nil
}
