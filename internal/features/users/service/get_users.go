package users_service

import (
	"context"
	"fmt"

	"github.com/BraveTonni/gotodo/internal/core/domain"
	core_errors "github.com/BraveTonni/gotodo/internal/core/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be not negative", core_errors.InvalidArgument)
	}

	return nil, nil
}
