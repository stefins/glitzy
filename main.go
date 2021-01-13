package main

import (
	"fmt"

	"github.com/iamstefin/glitzy/src/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Println(err)
	}
}
