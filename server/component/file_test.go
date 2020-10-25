package component

import (
	"testing"
	"time"
)

var file File = File{}

// It should assign and retrieve the right name.
func TestSetandGetName(t *testing.T) {
	file.SetName("Rambo")

	actual := file.GetName()
	expected := "Rambo"

	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve the right size.
func TestSetandGetSize(t *testing.T) {
	file.SetSize(123)
	actual := file.GetSize()
	expected := int64(123)
	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve the right last modification time.
func TestSetandGetLastModification(t *testing.T) {
	nowTime := time.Now()
	file.SetLastModification(nowTime)
	actual := file.GetLastModification()
	expected := nowTime
	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve if it is a directory/folder.
func TestSetandGetIsDirectory(t *testing.T) {
	file.SetIsFolder(false)
	actual := file.GetIsFolder()
	expected := false
	if actual != expected {
		t.Errorf("Expected: %v, and got: %v", expected, actual)
	}
}

// It should assign and retrieve and modify the name.
func TestModifyField(t *testing.T) {
	actual1 := file.GetName()
	expected1 := "Rambo"
	if actual1 != expected1 {
		t.Errorf("Expected: %v, and got: %v", expected1, actual1)
	}
	file.SetName("Conan")
	actual2 := file.GetName()
	expected2 := "Conan"
	if actual2 != expected2 {
		t.Errorf("Expected: %v, and got: %v", expected2, actual2)
	}
}
