package maintesting

import (
	"ridge/common/global"
	"testing"
)

// Run test logic - init
func Run(m *testing.M, rootPath string) int {
	// global root path
	global.GConfig.Root = rootPath

	return m.Run()
}
