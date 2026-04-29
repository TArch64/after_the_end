package styled

import (
	"embed"
)

//go:embed *.css
var fs embed.FS

var Global string

func init() {
	entries, err := fs.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var sheet []byte
	for _, entry := range entries {
		sheet, err = fs.ReadFile(entry.Name())
		if err != nil {
			panic(err)
		}

		Global += string(sheet)
		Global += "\n"
	}
}
