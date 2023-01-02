package config

import (
	"os"
	"strconv"
)

type ENV struct {
	SWAGGER  bool
	APP_PORT string
}

var Data ENV

func InitConfig() {
	Data.APP_PORT = os.Getenv("APP_PORT")
	Data.SWAGGER, _ = strconv.ParseBool(os.Getenv("SWAGGER"))
}

func AppPort() string {
	return Data.APP_PORT
}

func SwaggerMode() bool {
	return Data.SWAGGER
}
