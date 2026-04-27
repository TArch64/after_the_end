package resources

//go:generate miqt-rcc -Input "./app/resources/resources.qrc" -OutputGo "resources.go" -OutputRcc "resources.rcc" -Package "resources" -Qt6 -RccBinary "/Users/personal/Qt/6.10.3/macos/libexec/rcc"

import (
	"embed"

	qt "github.com/mappu/miqt/qt6"
)

//go:embed resources.rcc
var _resourceRcc []byte

func init() {
	_ = embed.FS{}
	qt.QResource_RegisterResourceWithRccData(&_resourceRcc[0])
}
