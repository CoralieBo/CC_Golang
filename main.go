package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := NewConfig()

	repos := getAllRepo(config)
	createCsv(repos)
	clone(repos)
	pullAllRepos(repos)
	compress()

	app := fiber.New()
	app.Get("/downloadZip", downloadRepository)
	log.Fatal(app.Listen(":3000"))
}
