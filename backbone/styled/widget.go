package styled

import (
	"fmt"

	qt "github.com/mappu/miqt/qt6"
)

func TitleShadow() *qt.QGraphicsEffect {
	glow := qt.NewQGraphicsDropShadowEffect()
	glow.SetBlurRadius(20)
	glow.SetColor(qt.NewQColor11(0, 0, 0, 120))
	glow.SetOffset2(2, 4)
	return glow.QGraphicsEffect
}

func S(selector, styles string) string {
	return fmt.Sprintf("%s {%s\n}", selector, styles)
}
