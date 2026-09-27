package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/BraveTonni/gotodo/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)

	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)

	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key '%s' not a valid int: %v: %w",
			param, key,
			err, core_errors.InvalidArgument,
		)
	}

	return &val, nil
}