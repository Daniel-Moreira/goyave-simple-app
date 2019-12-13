package route

import (
	"simple-app/http/controller/user"
	"simple-app/http/request/userrequest"

	"github.com/System-Glitch/goyave/v2"
)

func Register(router *goyave.Router) {
	router.Route("POST", "/user", user.Create, userrequest.User)
}
