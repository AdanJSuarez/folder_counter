package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AdanJSuarez/folder_counter/server/folderreader"
	"github.com/gofiber/fiber/v2"
)

// Init is used to run the server.
func Init() {

	app := fiber.New()

	app.Get("/:folderName", func(c *fiber.Ctx) error {
		fr := folderreader.FolderReader{}
		folderName := c.Params("folderName")
		log.Println("Folder name received:", folderName)
		fr.Read(folderName)
		jsonFR, err := json.Marshal(fr)
		if err != nil {
			log.Println("Failed marshaling fr:", err)
			c.SendString(fmt.Sprintf("Error: %v", err))
		}
		return c.Send(jsonFR)
	})
	app.Listen(":5000")
}
