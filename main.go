package main

import (
	"Evermos_Rakamin_BE_Project/app/config"
	"Evermos_Rakamin_BE_Project/app/router"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func init() {
	env := ".env"
	if len(os.Args) > 1 {
		env = os.Getenv(os.Args[1])
	}
	if err := gotenv.Load(env); err != nil {
		if err := gotenv.Load(); err != nil {
			log.Fatalln("Failed To Load .env :", err)
		}
	}
	config.InitConfig()
}

// @title Evermost Final Project
// @version 1.0.0
// @description Evermos Rakamin Virtual Internship Project
// @description Design DB: https://drive.google.com/file/d/1L7pzFNjMNfUU-f3tsrZPKTjMpbIGaodW/view
// @description API Contract: shorturl.at/AFH09
// @host :8000
func main() {
	if err := router.NewRouter(); err != nil {
		log.Println(err)
	}
}
