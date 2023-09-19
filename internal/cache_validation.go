package internal

import (
	"errors"
	"fmt"
)

func validateKey(key string) error {
	if len(key) > 0 {
		return nil
	}
	return errors.New(keyIsEmptyErrText)
}

func validateValue(value any) error {
	if value != nil {
		return nil
	}
	return errors.New(valueIsEmptyErrText)
}

func getKeyNotFoundError(key string) error {
	return fmt.Errorf(keyNotFoundErrTemplate, key)
}
