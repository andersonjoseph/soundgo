package shared

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	idHandler, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	id, err := idHandler.MarshalText()
	if err != nil {
		return "", err
	}

	return string(id), nil
}
