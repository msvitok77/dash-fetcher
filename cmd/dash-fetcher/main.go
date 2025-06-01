package main

import (
	"fmt"
	"os"

	"github.com/msvitok77/dash-fetcher/internal/commands"
)

func main() {
	err := commands.NewRoot().Execute()
	if err != nil {
		fmt.Printf("💀 fetcher execution err: %v", err)
		os.Exit(1)
	}

	fmt.Println("\n🏁 done")
}
