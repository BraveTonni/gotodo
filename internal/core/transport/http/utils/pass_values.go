package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/BraveTonni/gotodo/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)

	if pathValue == "" {
		return 0, fmt.Errorf("no key %s in pass values: %w", key, core_errors.InvalidArgument)
	}

	val, err := strconv.Atoi(pathValue)

	if err != nil {
		return 0, fmt.Errorf("cant convert value %s by key %s: %v: %w",
			pathValue,
			key,
			err,
			core_errors.InvalidArgument)
	}

	return val, nil
}
