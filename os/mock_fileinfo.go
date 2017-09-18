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

func (m MockFileInfo) Name() string          { return m.MockName }
func (m MockFileInfo) Size() int64           { return m.MockSize }
func (m MockFileInfo) Mode() coreos.FileMode { return m.MockMode }
func (m MockFileInfo) ModTime() time.Time    { return m.MockModTime }
func (m MockFileInfo) IsDir() bool           { return m.MockIsDir }
func (m MockFileInfo) Sys() interface{}      { return m.MockSys }
