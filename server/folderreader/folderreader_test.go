package folderreader

import (
	"strings"
	"testing"

	"github.com/AdanJSuarez/folder_counter/server/filestat"
)

var fileStat1 filestat.FileStat = filestat.FileStat{FileName: "turbo1", FileSize: 100}
var fileStat2 filestat.FileStat = filestat.FileStat{FileName: "turbo2", FileSize: 200}
var fileStat3 filestat.FileStat = filestat.FileStat{FileName: "turbo3", FileSize: 300}
var folderReader FolderReader

// setup reset folderReader to an empty object.
func setup() {
	folderReader = FolderReader{}
}

// It should return the right string representing the list of fileStats.
func TestGetFileStatsString(t *testing.T) {
	setup()
	folderReader.filesStats = append(folderReader.filesStats, fileStat1)
	actual := folderReader.GetFilesStatsString()
	expected := `
	[
    	{
			"fileName": "turbo1",
			"fileSize": 100,
			"fileLastModification": "0001-01-01T00:00:00Z",
			"fileIsDirectory": false
		}
	]`
	actual = strings.Replace(actual, "\t", "", -1)
	actual = strings.Replace(actual, "\n", "", -1)
	actual = strings.Replace(actual, " ", "", -1)

	expected = strings.Replace(expected, "\t", "", -1)
	expected = strings.Replace(expected, "\n", "", -1)
	expected = strings.Replace(expected, " ", "", -1)

	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should return the right number of files and size
func TestReadExistingFolder(t *testing.T) {
	setup()
	folderReader.Read("testFolder")
	actual1 := folderReader.GetTotalNumberOfFiles()
	expected1 := int64(1)
	if actual1 != expected1 {
		t.Errorf("Expected: %v, and got: %v", expected1, actual1)
	}
	actual2 := folderReader.GetTotalSize()
	expected2 := int64(31)
	if actual2 != expected2 {
		t.Errorf("Expected: %v, and got: %v", expected2, actual2)
	}
}

// It should return 0 files and 0 size for a none existing folder.
func TestReadNoneExistingFolder(t *testing.T) {
	setup()
	folderReader.Read("xxx")
	actual1 := folderReader.GetTotalNumberOfFiles()
	expected1 := int64(0)
	if actual1 != expected1 {
		t.Errorf("Expected: %v, and got: %v", expected1, actual1)
	}
	actual2 := folderReader.GetTotalSize()
	expected2 := int64(0)
	if actual2 != expected2 {
		t.Errorf("Expected: %v, and got: %v", expected2, actual2)
	}
}

// It should return a sorted by size list of fileStat.
func TestSortBySize(t *testing.T) {
	setup()
	folderReader.filesStats = append(folderReader.filesStats, fileStat2)
	folderReader.filesStats = append(folderReader.filesStats, fileStat3)
	folderReader.filesStats = append(folderReader.filesStats, fileStat1)

	if folderReader.filesStats[0].FileSize != 200 || folderReader.filesStats[1].FileSize != 300 || folderReader.filesStats[2].FileSize != 100 {
		t.Errorf("filesStats are not append in the right order: %v", folderReader.filesStats)
	}

	folderReader.sortBySize()
	actual := folderReader.filesStats
	folderReader.filesStats = append(folderReader.filesStats, fileStat3)
	folderReader.filesStats = append(folderReader.filesStats, fileStat2)
	folderReader.filesStats = append(folderReader.filesStats, fileStat1)
	expected := folderReader.filesStats

	if actual[0] != expected[0] || actual[1] != expected[1] || actual[2] != expected[2] {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}
