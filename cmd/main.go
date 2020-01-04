package main

import (
	"context"
	"fmt"
	"github.com/gobuz/publicspam/internal/application/http/rest"
	"github.com/gobuz/publicspam/internal/constant"
	"github.com/gobuz/publicspam/internal/domain/service"
	"github.com/gobuz/publicspam/internal/infrastructure/mongodb"
	config "github.com/gobuz/publicspam/pkg/configuration"
	"io"
	"net/http"
	"time"
)

func main() {
	err := config.Init(constant.KeyConfigFileName, []string{"."},true, func() error {
		return nil
	})

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)

	phoneRepository := mongodb.Init(ctx, config.Viper.GetString(constant.KeyDatabaseURI), config.Viper.GetString(constant.KeyDefaultDatabase))
	phoneService := service.InitPhoneService(phoneRepository)
	router, err := rest.InitRouter("",phoneService)

	defer func() {
		closer, ok := phoneRepository.(io.Closer)
		if ok {
			closer.Close()
		}
	}()

	if err != nil {
		fmt.Println(err)
	}
	panic(http.ListenAndServe(":8000", router))
}

