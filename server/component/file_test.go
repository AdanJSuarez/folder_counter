package component

import (
	"testing"
	"time"
)

var fileStat File = File{}

// It should assign and retrieve the right name.
func TestSetandGetName(t *testing.T) {
	fileStat.SetFileName("Rambo")

	actual := fileStat.GetFileName()
	expected := "Rambo"

	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve the right size.
func TestSetandGetSize(t *testing.T) {
	fileStat.SetFileSize(123)
	actual := fileStat.GetFileSize()
	expected := int64(123)
	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve the right last modification time.
func TestSetandGetLastModification(t *testing.T) {
	nowTime := time.Now()
	fileStat.SetFileLastModification(nowTime)
	actual := fileStat.GetFileLastModification()
	expected := nowTime
	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve if it is a directory/folder.
func TestSetandGetIsDirectory(t *testing.T) {
	fileStat.SetIsDirectory(false)
	actual := fileStat.GetIsDirectory()
	expected := false
	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve and modify the name.
func TestModifyField(t *testing.T) {
	actual1 := fileStat.GetFileName()
	expected1 := "Rambo"
	if actual1 != expected1 {
		t.Errorf("Expected: %v, and got: %v", expected1, actual1)
	}
	fileStat.SetFileName("Conan")
	actual2 := fileStat.GetFileName()
	expected2 := "Conan"
	if actual2 != expected2 {
		t.Errorf("Expected: %v, and got: %v", expected2, actual2)
	}
}
