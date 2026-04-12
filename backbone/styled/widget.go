package styled

import (
	"fmt"
)

const Transparent = "background: transparent;"

func S(selector, styles string) string {
	return fmt.Sprintf("%s {%s\n}", selector, styles)
}
