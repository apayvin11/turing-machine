package main

import (
	"fmt"
	"lab2/internal/machine"
	"log"
	"os"
)

const (
	alphabetPath = "alphabet.txt"
	tapePath     = "tape.txt"
	commandsPath = "commands.txt"
	resultPath   = "result.txt"
)

func main() {
	fmt.Println("Машина Тьюринга")
	fmt.Println("Вариант 39, функция 3x-y")
	f, err := os.Create(resultPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	machine := machine.New(alphabetPath, tapePath, commandsPath, f)
	machine.Run()
}
