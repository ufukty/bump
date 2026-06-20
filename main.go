package main

import (
	"fmt"
	"os"

	"github.com/ufukty/bump/internal/args"
)

func main() {
	if err := args.Dispatch(os.Args); err != nil {
		fmt.Println(err)
		fmt.Println("try: bump help")
		os.Exit(1)
	}
}
