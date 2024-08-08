package main

import (
	"fmt"
	"os"

	"github.com/datacharmer/cliptool/cmd"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		fmt.Println(cmd.CliptoolVersion)
	}

}
