package users_transport_http

import (
	"net/http"

	core_logger "github.com/BraveTonni/gotodo/internal/core/logger"
	core_http_response "github.com/BraveTonni/gotodo/internal/core/transport/http/response"
	core_http_utils "github.com/BraveTonni/gotodo/internal/core/transport/http/utils"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse("failed to get user with id", err)

		return
	}

	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse("failed to delete user", err)

		return
	}

	responseHandler.NoContentResponse()
}
