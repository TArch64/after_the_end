package resources

import (
	"fmt"
	"strings"

	"github.com/mappu/miqt/qt"
)

func Resource(root string, parts ...string) string {
	return fmt.Sprintf(":/%s/%s", root, strings.Join(parts, "/"))
}

func Image(parts ...string) string {
	return Resource("images", parts...)
}

func Icon(name string) string {
	return Resource("icons", name+".svg")
}

func QIcon(name string) *qt.QIcon {
	return qt.NewQIcon4(Icon(name))
}

func Font(name string) string {
	return Resource("fonts", name+".ttf")
}
