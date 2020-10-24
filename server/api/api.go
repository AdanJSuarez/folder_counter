package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AdanJSuarez/folder_counter/server/folderreader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// API define the API of folder counter
type API struct {
	app *fiber.App
	fr  folderreader.FolderReader
}

// Init is used to run the server.
func (api *API) Init() {

	api.app = fiber.New()
	api.app.Use(cors.New())

	api.app.Get("/api/v1/folder/:folderName/", func(c *fiber.Ctx) error {
		folderName := c.Params("folderName")
		api.fr = folderreader.FolderReader{}
		return api.readFolderName(folderName, c)
	})
	api.app.Get("api/v1/folder/", func(c *fiber.Ctx) error {
		api.fr = folderreader.FolderReader{}
		return api.readFolderName("", c)
	})
	api.app.Listen(":5000")
}

func (api *API) readFolderName(folderName string, c *fiber.Ctx) error {
	log.Println("Folder name received:", folderName)
	readErr := api.fr.Read(folderName)
	if readErr != nil {
		return c.SendStatus(204)
	}
	jsonFR, err := json.Marshal(api.fr)
	if err != nil {
		log.Println("Failed marshaling fr:", err)
		c.SendString(fmt.Sprintf("Error: %v", err))
	}
	return c.Send(jsonFR)
}
