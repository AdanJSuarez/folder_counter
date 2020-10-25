package api

import "github.com/AdanJSuarez/folder_counter/server/component"

// Response is structure use to set the api response.
type Response struct {
	ListOfComponent    []component.Component `json:"listOfComponent"`
	TotalSize          int64                 `json:"totalSize"`
	TotalNumberOfFiles int64                 `json:"totalNumberOfFiles"`
}
