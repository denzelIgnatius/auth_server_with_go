package validation

import "github.com/denzelIgnatius/auth_server_with_go/models"

func IsValidRequest(request models.Request) bool {

	if len(request.Username) == 0 || len(request.Password) == 0 {
		return false
	}
	return true
}
