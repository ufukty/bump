package help

import (
	_ "embed"
	"fmt"
)

//go:embed synopsis.txt
var synopsis string

func Run() error {
	fmt.Println(synopsis)
	return nil
}
