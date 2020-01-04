package rest

import (
	"github.com/gobuz/publicspam/internal/domain/port/service"
	"github.com/gin-gonic/gin"
	"os"
)

// API will be a factory for all implementation of the Dependencies
type API struct {
	*gin.Engine
	Identifier   string
	PhoneService service.Phone
}

// Init init the api that wrap gin gonic with project
func Init(identifier string, phone service.Phone) (*API, error) {

	api := &API{
		Identifier:identifier,
		Engine: gin.New(),
		PhoneService: phone,
	}

	if os.Getenv("GO_DEBUG") == "true"{
		api.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	return api, nil
}

