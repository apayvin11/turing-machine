package machine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
const (
	TAPE_SIZE = 256
	DEFAULT_HEAD_INDEX	= 14
	MAX_DEFAULT_HEAD_INDEX = TAPE_SIZE - 1
)

type Machine struct {
	alphabet []string
	commands []*Command
	tape     []byte
	headIndex int
}

func New(alphabetPath, tapePath, commandsPath string) *Machine {
	// Reading alphabet
	alphabetSlice := readFileIntoSliceByStr(alphabetPath)
	if len(alphabetSlice) != 1 {
		log.Fatal("invalid alphabet file")
	}
	alphabet := strings.Fields(alphabetSlice[0])
	for _, sym := range alphabet {
		if len(sym) != 1 {
			log.Fatal("invalid alphabet file")
		}
	}
	fmt.Printf("alphabet: %s\n", alphabet)

	// Reading tape
	tapeSlice := readFileIntoSliceByStr(tapePath)
	if len(tapeSlice) != 1 {
		log.Fatal("invalid tape file")
	}
	tapeStr := tapeSlice[0]
	fmt.Printf("tape: %s\n", tapeStr)
	tape := make([]byte, 256)
	copy(tape[14:], []byte(tapeStr))

	// Reading commands
	commandsProto := readFileIntoSliceByStr(commandsPath)
	fmt.Println("command prototypes: ", commandsProto)
	commands := []*Command{}
	for _, proto := range commandsProto {
		command, err := parseCommand(proto)
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, command)
	}
	return &Machine{
		alphabet: alphabet,
		tape:     tape,
		commands: commands,
		headIndex: DEFAULT_HEAD_INDEX,
	}
}

func (m *Machine) Run() {
	for _, command := range m.commands {
		fmt.Println(command.String())
	}
}

// readFileIntoSliceByStr читает файл построчно в слайс
func readFileIntoSliceByStr(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
}