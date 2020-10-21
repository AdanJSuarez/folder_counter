/**
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package folderreader

import (
	"log"
	"os"

	"github.com/AdanJSuarez/folder_counter/src/filestat"
)

// FolderReader is a structure to handle reading folders and files from a specific folder.
// If not folder is specified takes the actual folder where it is running.
type FolderReader struct {
	filesStats []filestat.FileStat
}

// GetFileNames return a slice of all names
func (fr *FolderReader) GetFileNames() []filestat.FileStat {
	return fr.filesStats
}

// Read set the name of folders and files to fileNames
func (fr *FolderReader) Read(folderName string) []filestat.FileStat {
	var result []filestat.FileStat
	if folderName == "" {
		folderName = "."
	}
	file, err := os.Open(folderName)
	if err != nil {
		log.Printf("Failed opening directory: %v. Please check the name.", err)
	}
	defer file.Close()

	fileNames, err := file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		log.Println("Failed to read files:", err)
	}
	for _, name := range fileNames {
		fileStat := filestat.FileStat{}
		fr.ReadStats(name, fileStat)
		result = append(result, fileStat)
	}
	return result
}

// ReadStats return the stats of a file of named fileName.
func (fr *FolderReader) ReadStats(fileName string, fileStat filestat.FileStat) {
	//Gets stats of the file
	stats, err := os.Stat(fileName)
	if err != nil {
		log.Printf("Failed to read stats of %s: %v", fileName, err)
	}
	fileStat.SetFileSize(stats.Size())
	fileStat.SetFileLastModification(stats.ModTime())
	fileStat.SetIsDirectory(stats.IsDir())
	//Prints stats of the file
	// fmt.Printf("Permission: %s\n", stats.Mode())
	// fmt.Printf("Size: %d\n", size)
	// fmt.Printf("Modification Time: %s\n", lastModification)
	// fmt.Printf("Is directory: %v\n", isFolder)
}
