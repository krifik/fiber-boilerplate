package app

import (
	"fmt"
	"log"
	"mangojek-backend/config"
	"mangojek-backend/module"
	"mangojek-backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/urfave/cli"
)

func InitializedApp() *fiber.App {
	configuration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configuration)
	userController := module.NewUserModule(database)
	productController := module.NewProductModule(database)
	// Setup Fiber
	// app := fiber.New(config.NewFiberConfig())
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New(), cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	// Setup Routing
	routes.Route(app, userController, productController)
	return app

}

func InitializeDB() {
	configration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configration)

	cmdApp := cli.NewApp()

	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(cli *cli.Context) error {
				// migration function
				config.NewRunMigration(database)
				fmt.Println("================ migrated successfully ================")
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(cli *cli.Context) error {
				// seeding function
				config.NewRunSeed(database)
				fmt.Println("================ seeded successfully ================")
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
