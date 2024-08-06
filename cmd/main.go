package main

import (
	"fmt"
	"github.com/apayvin11/turing-machine/internal/machine"
	"log"
	"os"
)

const (
	alphabetPath = "settings/alphabet.txt"
	tapePath     = "settings/tape.txt"
	commandsPath = "settings/commands.txt"
	resultPath   = "output/result.txt"
)

func main() {
	fmt.Println("Turing machine")
	fmt.Println("function 2x-y")

	resultFile, err := os.Create(resultPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resultFile.Close()

	machine, err := machine.New(alphabetPath, tapePath, commandsPath, resultFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := machine.Run(); err != nil {
		log.Fatal(err)
	}
}
