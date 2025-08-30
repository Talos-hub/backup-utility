package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Talos-hub/backup-utilit/pkg/read"
)

func main() {
	data, err := read.Read[string](os.Stdin, read.StrategyDirect)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
