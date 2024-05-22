package main

import (
	"fmt"

	"server/config"
	"server/pkg"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	pkg.HandleEnv()

	config.DatabaseConnection(pkg.DB_USER, pkg.DB_PASSWORD, pkg.DB_URI, pkg.DB_NAME)

	app := fiber.New(fiber.Config{DisablePreParseMultipartForm: true, StreamRequestBody: true})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     pkg.URI_HOST,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	app.Static("/", pkg.DIR_PUBLIC)

	routes.MidtransRoute(app)
	routes.AuthRoute(app)
	routes.SantunanRoute(app)
	routes.PengurusRoute(app)
	routes.DonaturRoute(app)
	routes.ProfileRoute(app)
	routes.AnakYatimRoute(app)
	routes.DanaSantunanRoute(app)
	routes.LaporanDanaSantunanRoute(app)
	routes.DanaSantunanAnakAsuhRoute(app)
	routes.DataCountRoute(app)

	app.Listen(fmt.Sprintf(":%s", pkg.URI_PORT))
}
