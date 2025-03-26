package main

import (
	"edugo/config"
	"edugo/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.InitDB()

	routes.NormalRoutes(app)  //Global use
	routes.AuthRoutes(app)    //Authentication
	routes.ProfileRoutes(app) //Profile
	routes.TutorRoutes(app)   //Tutor
	routes.AdminRoutes(app)   //Admin

	log.Fatal(app.Listen(":3000"))
}
