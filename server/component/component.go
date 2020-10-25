package component

import (
	"time"
)

// Component is the interface for folders and files
type Component interface {
	GetName() string
	GetSize() int64
	GetLastModification() time.Time
	GetIsFolder() bool
}
