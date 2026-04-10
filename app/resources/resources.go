package resources

//go:generate miqt-rcc -Input "./app/resources/resources.qrc" -OutputGo "resources.go" -OutputRcc "resources.rcc" -Package "resources"

import (
	"embed"

	"github.com/mappu/miqt/qt"
)

//go:embed resources.rcc
var _resourceRcc []byte

func init() {
	_ = embed.FS{}
	qt.QResource_RegisterResourceWithRccData(&_resourceRcc[0])
}
