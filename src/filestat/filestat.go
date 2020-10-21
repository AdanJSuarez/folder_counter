/**
Adan J. Suarez
e: adanjsuarez@gmail.com
*/

package filestat

import "time"

// FileStat represent a file or a folder containing all relevant info.
type FileStat struct {
	fileName             string    `json:"fileName"`
	fileSize             int64     `json:"fileSize"`
	fileLastModification time.Time `json:"fileLastModification"`
	fileIsDirectory      bool      `json:"fileIsDirectory"`
}

//GetFileName return the name of the file
func (fs *FileStat) GetFileName() string {
	return fs.fileName
}

// GetFileSize return the size of the file.
func (fs *FileStat) GetFileSize() int64 {
	return fs.fileSize
}

// GetFileLastModification returns the last modification time.
func (fs *FileStat) GetFileLastModification() time.Time {
	return fs.fileLastModification
}

// GetIsDirectory returns true if file is a directory/folder, false otherwise.
func (fs *FileStat) GetIsDirectory() bool {
	return fs.fileIsDirectory
}

// SetFileName set fileName as the name of the file.
func (fs *FileStat) SetFileName(fileName string) {
	fs.fileName = fileName
}

// SetFileSize set the fileSize as the name of the file.
func (fs *FileStat) SetFileSize(fileSize int64) {
	fs.fileSize = fileSize
}

// SetFileLastModification set the lastModification as the last modification time of the file.
func (fs *FileStat) SetFileLastModification(lastModification time.Time) {
	fs.fileLastModification = lastModification
}

//SetIsDirectory set true if file is directory/folder, false otherwise
func (fs *FileStat) SetIsDirectory(isDirectory bool) {
	fs.fileIsDirectory = isDirectory
}
