package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func downloadRepository(c *fiber.Ctx) error {
	zipFilePath := "./archive.zip"

	_, err := os.Stat(zipFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("Fichier ZIP introuvable")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur serveur")
	}

	c.Set("Content-Disposition", "attachment; filename=archive.zip")
	c.Set("Content-Type", "application/zip")

	return c.SendFile(zipFilePath)
}
