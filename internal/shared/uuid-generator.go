package shared

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	idHandler, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("eror creating uuid: %w", err)
	}

	id, err := idHandler.MarshalText()
	if err != nil {
		return "", fmt.Errorf("eror marshalling uuid: %w", err)
	}

	return string(id), nil
}
