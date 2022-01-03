package main

import (
	"Donation/Configration"
	"Donation/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main(){
	log.Println("Start Listening ...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	app:=fiber.New()
	Routes.SetUpUSer(app)
	_=app.Listen(":8080")
	Configration.SetUpDB()

}
