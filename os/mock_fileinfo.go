package os

import (
	coreos "os"
	"time"
)

// MockFileInfo provides a mock os.FileInfo structure.
type MockFileInfo struct {
	MockName    string
	MockSize    int64
	MockMode    coreos.FileMode
	MockModTime time.Time
	MockIsDir   bool
	MockSys     interface{}
}

// Verify that ihe os.FileInfo interface is implemented properly.
var _ coreos.FileInfo = new(MockFileInfo)

// Name implements os.FileInfo interface.
func (m MockFileInfo) Name() string { return m.MockName }

// Size implements os.FileInfo interface.
func (m MockFileInfo) Size() int64 { return m.MockSize }

// Mode implements os.FileInfo interface.
func (m MockFileInfo) Mode() coreos.FileMode { return m.MockMode }

// ModTime implements os.FileInfo interface.
func (m MockFileInfo) ModTime() time.Time { return m.MockModTime }

// IsDir implements os.FileInfo interface.
func (m MockFileInfo) IsDir() bool { return m.MockIsDir }

// Sys implements os.FileInfo interface.
func (m MockFileInfo) Sys() interface{} { return m.MockSys }
