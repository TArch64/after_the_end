package styled

import (
	"github.com/mappu/miqt/qt"
)

const Title1 = `
color: #fff;
font-size: 80px;
background: transparent;`

const Title2 = `
color: #fff;
font-size: 60px;
background: transparent;`

const BodyWhite = `
color: #fff;
font-size: 20px;
background: transparent;`

func TitleShadow() *qt.QGraphicsEffect {
	glow := qt.NewQGraphicsDropShadowEffect()
	glow.SetBlurRadius(20)
	glow.SetColor(qt.NewQColor11(0, 0, 0, 120))
	glow.SetOffset2(2, 4)
	return glow.QGraphicsEffect
}
