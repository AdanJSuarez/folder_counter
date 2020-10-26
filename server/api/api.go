package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AdanJSuarez/folder_counter/server/component"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// API define the API of folder counter
type API struct {
	app *fiber.App
	f   component.Folder
}

// Init is used to run the server.
func (api *API) Init() {

	api.app = fiber.New()
	api.app.Use(cors.New())

	api.app.Post("/api/v1/folder/", func(c *fiber.Ctx) error {
		folderName := string(c.Body())
		api.f = component.Folder{}
		return api.readFolderName(folderName, c)
	})
	api.app.Listen(":5000")
}

// readFolderName set a new folder and populate the whole tree of components.
func (api *API) readFolderName(folderName string, c *fiber.Ctx) error {
	log.Println("Folder name received:", folderName)
	readErr := api.f.New(folderName)
	if readErr != nil {
		return c.SendStatus(204)
	}
	response := Response{TotalSize: api.f.GetSize(), TotalNumberOfFiles: api.f.GetTotalFiles(),
		ListOfComponent: api.f.GetListOfComponent()}
	jsonFR, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed marshaling fr:", err)
		c.SendString(fmt.Sprintf("Error: %v", err))
	}
	return c.Send(jsonFR)
}
