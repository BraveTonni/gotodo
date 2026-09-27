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
		return nil, fmt.Errorf("limit must be not negative: %w", core_errors.InvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("limit must be not negative: %w", core_errors.InvalidArgument)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("get users from repo: %w", err)
	}

	return users, nil
}
