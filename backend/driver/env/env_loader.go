package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ValueNotFoundError = "%s is not found. this environment value must be set."

func GetStrEnv(key string, required bool) (string, error) {
	v, exist := os.LookupEnv(key)
	if required && !exist {
		return "", errors.New(fmt.Sprintf(ValueNotFoundError, key))
	}

	return v, nil
}

func GetIntEnv(key string, required bool) (int, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if required {
			return 0, errors.New(fmt.Sprintf(ValueNotFoundError, key))
		}
		return 0, nil
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return n, nil
}
