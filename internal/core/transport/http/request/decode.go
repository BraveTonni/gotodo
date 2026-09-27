package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/BraveTonni/gotodo/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf(
			"failed to decode request body: %v: %w",
			err,
			core_errors.InvalidArgument,
		)
	}

	if err := requestValidator.Struct(dest); err != nil {
		return fmt.Errorf("failed to validate request body: %v: %w", err, core_errors.InvalidArgument)
	}

	return nil
}
