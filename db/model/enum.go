package model

import (
	"errors"
)

func scanEnum[V ~string](value any, variants []V) (V, error) {
	if value == nil {
		return "", nil
	}
	var text string
	if str, ok := value.(string); ok {
		text = str
	} else if bytes, ok := value.([]byte); ok {
		text = string(bytes)
	} else {
		return "", errors.New("invalid enum data type")
	}
	if validateEnumValue(text, variants) {
		return V(text), nil
	}
	return "", errors.New("invalid enum value")
}

func valueEnum[V ~string](value V, variants []V) (string, error) {
	if text := string(value); validateEnumValue(text, variants) {
		return string(value), nil
	}
	return "", errors.New("invalid enum value")
}

func validateEnumValue[V ~string](value string, variants []V) bool {
	for _, variant := range variants {
		if value == string(variant) {
			return true
		}
	}
	return false
}
