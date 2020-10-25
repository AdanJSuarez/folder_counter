package component

import (
	"testing"
)

var file1 File = File{FileName: "turbo1", FileSize: 100}
var file2 File = File{FileName: "turbo2", FileSize: 200}
var file3 File = File{FileName: "turbo3", FileSize: 300}
var folder Folder

// setup reset folderReader to an empty object.
func setup() {
	folder = Folder{}
}

// It should return the right number of files and size
func TestReadExistingFolder(t *testing.T) {
	setup()
	folder.New("testFolder")
	actual1 := folder.GetTotalFiles()
	expected1 := int64(2)
	if actual1 != expected1 {
		t.Errorf("Expected: %v, and got: %v", expected1, actual1)
	}
	actual2 := folder.GetSize()
	expected2 := int64(67)
	if actual2 != expected2 {
		t.Errorf("Expected: %v, and got: %v", expected2, actual2)
	}
}

// It should return 0 files and 0 size for a none existing folder.
func TestReadNoneExistingFolder(t *testing.T) {
	setup()
	folder.New("xxx")
	actual1 := folder.GetTotalFiles()
	expected1 := int64(0)
	if actual1 != expected1 {
		t.Errorf("Expected: %v, and got: %v", expected1, actual1)
	}
	actual2 := folder.GetSize()
	expected2 := int64(0)
	if actual2 != expected2 {
		t.Errorf("Expected: %v, and got: %v", expected2, actual2)
	}
}

// It should return a sorted by size list of fileStat.
func TestSortBySize(t *testing.T) {
	setup()
	folder.listOfComponent = append(folder.listOfComponent, &file2)
	folder.listOfComponent = append(folder.listOfComponent, &file3)
	folder.listOfComponent = append(folder.listOfComponent, &file1)

	if folder.listOfComponent[0].GetSize() != 200 || folder.listOfComponent[1].GetSize() != 300 || folder.listOfComponent[2].GetSize() != 100 {
		t.Errorf("files are not append in the right order: %v", folder.listOfComponent)
	}
	folder.sortBySize()
	actual0 := folder.listOfComponent[0].GetSize()
	actual1 := folder.listOfComponent[1].GetSize()
	actual2 := folder.listOfComponent[2].GetSize()

	if actual0 != 300 || actual1 != 200 || actual2 != 100 {
		t.Errorf("Expected list of component sorted, and got: %v", folder.listOfComponent)
	}
}
