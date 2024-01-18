package main

import (
	"fmt"
	"lab2/internal/machine"
)

const (
	alphabetPath = "alphabet.txt"
	tapePath     = "tape.txt"
	commandsPath = "commands.txt"
)

func main() {
	fmt.Println("Машина Тьюринга")
	fmt.Println("Вариант 39, функция 3x-y")
	machine := machine.New(alphabetPath, tapePath, commandsPath)
	machine.Run()
}
