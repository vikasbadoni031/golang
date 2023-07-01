package main

import (
	cmd "kubectl-deploystat/cmd"
	"log"
)

func main() {
	if err := cmd.Deploystat().Execute(); err != nil {
		log.Fatal(err)
	}
}
