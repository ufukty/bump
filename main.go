package main

import (
	"fmt"
	"os"

	"github.com/ufukty/bump/internal/commands"
)

func main() {
	if err := commands.Dispatch(os.Args[1:]); err != nil {
		fmt.Println(err)
		fmt.Println("try: bump help")
		os.Exit(1)
	}
}
