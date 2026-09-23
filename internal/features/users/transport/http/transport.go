package users_transport_http

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}
