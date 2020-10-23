package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AdanJSuarez/folder_counter/server/folderreader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Init is used to run the server.
func Init() {

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/v1/folder/:folderName/", func(c *fiber.Ctx) error {
		fr := folderreader.FolderReader{}
		folderName := c.Params("folderName")
		if folderName == "xxxoooxxx" {
			folderName = ""
		}
		log.Println("Folder name received:", folderName)
		readErr := fr.Read(folderName)
		if readErr != nil {
			msg := "Folder not found: " + folderName
			return c.Send([]byte(msg))
		}
		jsonFR, err := json.Marshal(fr)
		if err != nil {
			log.Println("Failed marshaling fr:", err)
			c.SendString(fmt.Sprintf("Error: %v", err))
		}
		return c.Send(jsonFR)
	})
	app.Listen(":5000")
}
