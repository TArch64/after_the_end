package styled

import (
	"fmt"
)

const Reset = Transparent + BorderNone

const Transparent = "background: transparent;"

const BorderNone = "border: none;"

func S(selector, styles string) string {
	return fmt.Sprintf("%s { %s }", selector, styles)
}
