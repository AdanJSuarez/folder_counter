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

	"github.com/AdanJSuarez/folder_counter/src/filestat"
)

// FolderReader is a structure to handle reading folders and files from a specific folder.
// If not folder is specified takes the actual folder where it is running.
type FolderReader struct {
	filesStats         []filestat.FileStat
	totalSize          int64
	totalNumberOfFiles int64
}

// GetFilesStatsString return a slice of all names
func (fr *FolderReader) GetFilesStatsString() string {
	jsonFileStats, err := json.MarshalIndent(fr.filesStats, "  ", "")
	if err != nil {
		log.Printf("Failed to marshal files stats: %v", err)
	}
	return string(jsonFileStats)
}

// GetTotalSize return the total size of all files and folders
func (fr *FolderReader) GetTotalSize() int64 {
	return fr.totalSize
}

// GetTotalNumberOfFiles return the total number of file and folders
func (fr *FolderReader) GetTotalNumberOfFiles() int64 {
	return fr.totalNumberOfFiles
}

// Read is used to set files stats and total size and number of files.
func (fr *FolderReader) Read(folderName string) {
	if folderName == "" {
		folderName = "."
	}
	file, err := os.Open(folderName)
	if err != nil {
		log.Printf("Failed opening directory: '%v' Please check the name.", folderName)
	}
	defer file.Close()

	fileNames, err := file.Readdirnames(0)
	if err != nil {
		log.Println("Failed to read files:", err)
	}
	fr.setListOfFileStats(folderName, fileNames)
}

// setListOfFileStats is used to set listOfFileStats
// A posible performance improvement could be insert files in order using BST (i.e.)
func (fr *FolderReader) setListOfFileStats(folderName string, fileNames []string) {
	for _, name := range fileNames {
		fileStat := filestat.FileStat{}
		fr.readStats(folderName+"/"+name, &fileStat)
		fr.totalSize += fileStat.GetFileSize()
		fr.totalNumberOfFiles++
		fr.filesStats = append(fr.filesStats, fileStat)
	}
	fr.sortBySize()
}

// readStats return the stats of a file of named fileName.
func (fr *FolderReader) readStats(fileName string, fileStat *filestat.FileStat) {
	stats, err := os.Stat(fileName)
	if err != nil {
		log.Printf("Failed to read stats of %s: %v", fileName, err)
	}
	fileStat.SetFileName(fileName)
	fileStat.SetFileSize(stats.Size())
	fileStat.SetFileLastModification(stats.ModTime())
	fileStat.SetIsDirectory(stats.IsDir())
}

// sortBySize set filesStats order by file size
func (fr *FolderReader) sortBySize() {
	sort.Slice(fr.filesStats, func(i, j int) bool {
		return fr.filesStats[i].GetFileSize() > fr.filesStats[j].GetFileSize()
	})
}
