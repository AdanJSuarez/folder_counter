/**
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package component

import "time"

// File represent a file or a folder containing all relevant info.
type File struct {
	FileName             string    `json:"fileName"`
	FileSize             int64     `json:"fileSize"`
	FileLastModification time.Time `json:"fileLastModification"`
	IsFolder             bool      `json:"fileIsFolder"`
}

//GetName return the name of the file
func (f *File) GetName() string {
	return f.FileName
}

// GetSize return the size of the file.
func (f *File) GetSize() int64 {
	return f.FileSize
}

// GetLastModification returns the last modification time.
func (f *File) GetLastModification() time.Time {
	return f.FileLastModification
}

// GetIsFolder returns true if file is a directory/folder, false otherwise.
func (f *File) GetIsFolder() bool {
	return f.IsFolder
}

// SetName set fileName as the name of the file.
func (f *File) SetName(fileName string) {
	f.FileName = fileName
}

// SetSize set the fileSize as the name of the file.
func (f *File) SetSize(fileSize int64) {
	f.FileSize = fileSize
}

// SetLastModification set the lastModification as the last modification time of the file.
func (f *File) SetLastModification(lastModification time.Time) {
	f.FileLastModification = lastModification
}

//SetIsFolder set true if file is directory/folder, false otherwise
func (f *File) SetIsFolder(isFolder bool) {
	f.IsFolder = isFolder
}
