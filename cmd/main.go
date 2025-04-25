package main

import (
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/hamza-sharif/home-assessment-kai-cyber"
	"github.com/hamza-sharif/home-assessment-kai-cyber/apis"
	"github.com/hamza-sharif/home-assessment-kai-cyber/config"
	"github.com/hamza-sharif/home-assessment-kai-cyber/gen/restapi"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	runT, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := apis.NewApis(runT, swaggerSpec)

	server := restapi.NewServer(api)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	server.Host = viper.GetString(config.ServerHost)

	if err != nil {
		panic(err)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
