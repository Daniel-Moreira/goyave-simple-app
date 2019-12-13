package user

import (
	"net/http"

	"github.com/System-Glitch/goyave/v2"
	"github.com/System-Glitch/goyave/v2/lang"
)

func Create(response *goyave.Response, request *goyave.Request) {
	response.String(http.StatusOK, lang.Get(request.Lang, lang.Get(request.Lang, "create_response")))
}
