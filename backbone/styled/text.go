package styled

import (
	"github.com/mappu/miqt/qt"
)

const Title1 = `
color: #fff;
font-size: 80px;` + Reset

const Title2 = `
color: #fff;
font-size: 60px;` + Reset

const BodyWhite = `
color: #444;
font-size: 20px;` + Reset

const BodyWhite2 = `
color: #888;
font-size: 18px;` + Reset

func TitleShadow() *qt.QGraphicsEffect {
	glow := qt.NewQGraphicsDropShadowEffect()
	glow.SetBlurRadius(20)
	glow.SetColor(qt.NewQColor11(0, 0, 0, 120))
	glow.SetOffset2(2, 4)
	return glow.QGraphicsEffect
}
