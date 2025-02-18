package component

import (
	"log"
	"os"
	"sort"
	"time"
)

// Folder represent a folder component.
type Folder struct {
	FolderName             string `json:"name"`
	FolderSize             int64  `json:"size"`
	FolderTotalFiles       int64
	FolderLastModification time.Time `json:"lastModification"`
	IsFolder               bool      `json:"isFolder"`
	listOfComponent        []Component
}

//GetName return the name of the folder
func (f Folder) GetName() string {
	return f.FolderName
}

// GetSize return the size of the folder
func (f Folder) GetSize() int64 {
	var result int64
	for _, component := range f.listOfComponent {
		result += component.GetSize()
	}
	return result
}

// GetTotalFiles return the number of file in this folder, including those in subfolder if any.
func (f *Folder) GetTotalFiles() int64 {
	return f.FolderTotalFiles
}

// GetLastModification returns the last modification time.
func (f Folder) GetLastModification() time.Time {
	return f.FolderLastModification
}

// GetListOfComponent returns the list of all component of this folder.
func (f Folder) GetListOfComponent() []Component {
	return f.listOfComponent
}

// GetIsFolder returns true because folder always is a folder.
func (f Folder) GetIsFolder() bool {
	return f.IsFolder
}

// SetName set fileName as the name of the file.
func (f *Folder) SetName(fileName string) {
	f.FolderName = fileName
}

// SetSize set the fileSize as the name of the file.
func (f *Folder) SetSize(fileSize int64) {
	f.FolderSize = fileSize
}

// SetLastModification set the lastModification as the last modification time of the file.
func (f *Folder) SetLastModification(lastModification time.Time) {
	f.FolderLastModification = lastModification
}

// SetTotalFiles set the number of files.
func (f *Folder) SetTotalFiles(totalNumberOfFiles int64) {
	f.FolderTotalFiles = totalNumberOfFiles
}

//SetIsFolder set true if file is directory/folder, false otherwise
func (f *Folder) SetIsFolder(isFolder bool) {
	f.IsFolder = isFolder
}

// New is used to set files stats and total size and number of files.
func (f *Folder) New(folderName string) error {
	if folderName == "" {
		folderName = "./" + folderName
	}
	stats, err := os.Stat(folderName)
	if err != nil {
		log.Printf("Error#3, Failed to read folder %s: %v", folderName, err)
		return err
	}
	f.FolderName = folderName
	f.FolderLastModification = stats.ModTime()
	f.setListOfComponent(folderName)
	f.sortBySize()
	return nil
}

// setListOfComponent is used to set the list of component of this folder.
// It traverse the tree in a depth-first search (DFS) manner.
func (f *Folder) setListOfComponent(folderName string) (int64, int64) {
	componentReader, err := os.Open(folderName)
	if err != nil {
		log.Printf("Error#4, Failed opening directory: '%v' Please check the name.", folderName)
	}
	defer componentReader.Close()
	componentNames, err := componentReader.Readdirnames(0)
	if err != nil {
		log.Println("Error#5, Failed to read files:", err)
	}
	for _, name := range componentNames {
		file := File{}
		folder := Folder{}
		stats, err := os.Stat(folderName + "/" + name)
		if err != nil {
			log.Printf("Error#6, Failed to read stats of %s: %v", name, err)
		}
		if stats.IsDir() {
			f.setFolderData(folderName, name, &folder, stats)
		} else {
			f.setFileData(folderName, name, &file, stats)
		}
	}
	return f.FolderSize, f.FolderTotalFiles
}

// setFolderData set the data in a folder component.
func (f *Folder) setFolderData(folderName string, name string, folder *Folder, stats os.FileInfo) {
	folder.SetName(folderName + "/" + name)
	folder.SetIsFolder(true)
	folder.SetLastModification(stats.ModTime())
	size, totalFiles := folder.setListOfComponent(folderName + "/" + name)
	f.FolderSize += size
	f.FolderTotalFiles += totalFiles
	f.listOfComponent = append(f.listOfComponent, folder)
}

// setFileData set the data in a file component.
func (f *Folder) setFileData(folderName string, name string, file *File, stats os.FileInfo) {
	file.SetName(folderName + "/" + name)
	file.SetSize(stats.Size())
	file.SetLastModification(stats.ModTime())
	file.SetIsFolder(stats.IsDir())
	f.FolderSize += stats.Size()
	f.FolderTotalFiles++
	f.listOfComponent = append(f.listOfComponent, file)
}

// sortBySize set filesStats order by file size
func (f *Folder) sortBySize() {
	sort.Slice(f.listOfComponent, func(i, j int) bool {
		return f.listOfComponent[i].GetSize() > f.listOfComponent[j].GetSize()
	})
}
