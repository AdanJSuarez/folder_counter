/**
Entry point for "folder_counter"
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package main

import (
	"log"

	"github.com/AdanJSuarez/folder_counter/src/folderreader"
)

func main() {
	fr := folderreader.FolderReader{}
	fr.Read("")
	log.Println(fr)
	return
}
