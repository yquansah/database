package main

import (
	"fmt"
	"os"

	"github.com/yquansah/database/cmd/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
