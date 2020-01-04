package rest

import (
	"github.com/gobuz/publicspam/internal/domain/port/service"
)

// InitRouter init route for rest api
func InitRouter(apiIdentity string, phoneService service.Phone) (*API, error) {
	rest, err := Init(apiIdentity, phoneService)

	rest.GET("phone/spams",rest.SpamPhones)
	rest.GET("phone/spams/:phone_number", rest.SpamPhone)

	return rest, err
}
