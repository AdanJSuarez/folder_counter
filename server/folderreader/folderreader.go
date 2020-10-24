/**
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package folderreader

import (
	"encoding/json"
	"log"
	"os"
	"sort"

	"github.com/AdanJSuarez/folder_counter/server/component"
)

// FolderReader is a structure to handle reading folders and files from a specific folder.
// If not folder is specified takes the actual folder where it is running.
type FolderReader struct {
	ListOfComponent    []component.Component `json:"listOfComponents"`
	TotalSize          int64                 `json:"totalSize"`
	TotalNumberOfFiles int64                 `json:"totalNumberOfFiles"`
}

// GetFilesStatsString return a slice of all names
func (fr *FolderReader) GetFilesStatsString() string {
	jsonFileStats, err := json.MarshalIndent(fr.ListOfComponent, "  ", "")
	if err != nil {
		log.Printf("Failed to marshal files stats: %v", err)
	}
	return string(jsonFileStats)
}

// GetTotalSize return the total size of all files and folders
func (fr *FolderReader) GetTotalSize() int64 {
	return fr.TotalSize
}

// GetTotalNumberOfFiles return the total number of file and folders
func (fr *FolderReader) GetTotalNumberOfFiles() int64 {
	return fr.TotalNumberOfFiles
}

// Read is used to set files stats and total size and number of files.
func (fr *FolderReader) Read(folderName string) error {
	if folderName == "" {
		folderName = "."
	}
	component, err := os.Open(folderName)
	if err != nil {
		log.Printf("Failed opening directory: '%v' Please check the name.", folderName)
		return err
	}
	defer component.Close()

	// componentNames, err := component.Readdirnames(0)
	if err != nil {
		log.Println("Failed to read files:", err)
	}
	// fr.setListOfComponent(folderName, componentNames)
	fr.sortBySize()
	return nil
}

// setListOfComponent is used to set listOfFileStats
// A posible performance improvement could be insert files in order using BST (i.e.)
// func (fr *FolderReader) setListOfComponent(folderName string, componentNames []string) {
// 	for _, name := range componentNames {
// 		file := component.File{}
// 		folder := component.Folder{}
// 		stats, err := os.Stat(name)
// 		if err != nil {
// 			log.Printf("Failed to read stats of %s: %v", name, err)
// 		}
// 		if stats.IsDir() {
// 			fr.readStats(folderName+"/"+name, &folder)
// 		} else {
// 			file.SetName(name)
// 			file.SetSize(stats.Size())
// 			file.SetLastModification(stats.ModTime())
// 			file.SetIsFolder(stats.IsDir())
// 			fr.TotalSize += stats.Size()
// 			fr.TotalNumberOfFiles++
// 			fr.ListOfComponent = append(fr.ListOfComponent, &file)
// 		}

// 		// fr.TotalSize += file.GetSize()
// 	}
// }

// readStats return the stats of a file of named fileName.
func (fr *FolderReader) readSize(fileName string, component component.Component) {
	stats, err := os.Stat(fileName)
	if err != nil {
		log.Printf("Failed to read stats of %s: %v", fileName, err)
	}
	if stats.IsDir() {
		//component.SetSize(stats.Size())

	} else {
		fr.TotalSize += stats.Size()
		fr.TotalNumberOfFiles++
	}

}

// sortBySize set filesStats order by file size
func (fr *FolderReader) sortBySize() {
	sort.Slice(fr.ListOfComponent, func(i, j int) bool {
		return fr.ListOfComponent[i].GetSize() > fr.ListOfComponent[j].GetSize()
	})
}
