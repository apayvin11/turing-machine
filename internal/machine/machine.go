package machine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	TAPE_SIZE              = 64
	DEFAULT_HEAD_INDEX     = 14
	MAX_DEFAULT_HEAD_INDEX = TAPE_SIZE - 1
)

type Machine struct {
	alphabet    []string
	commandBase map[string]map[byte]*Command
	tape        []byte
	headIndex   int
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
	tape := make([]byte, TAPE_SIZE)
	for i := range tape {
		tape[i] = '_'
	}
	copy(tape[DEFAULT_HEAD_INDEX:], []byte(tapeStr))

	// Reading commands
	commandsProto := readFileIntoSliceByStr(commandsPath)
	fmt.Println("command prototypes: ", commandsProto)
	commandBase := make(map[string]map[byte]*Command)
	for _, proto := range commandsProto {
		command, err := parseCommand(proto)
		if err != nil {
			log.Fatal(err)
		}
		_, ok := commandBase[command.stateBefore]
		if !ok {
			commandBase[command.stateBefore] = map[byte]*Command{}
		}
		commandBase[command.stateBefore][command.symBefore] = command
	}
	return &Machine{
		alphabet:    alphabet,
		tape:        tape,
		commandBase: commandBase,
		headIndex:   DEFAULT_HEAD_INDEX,
	}
}

func (m *Machine) Run() {
	state := "q0"
	for {
		fmt.Println(m.GetTapeState())
		if state == END_STATE {
			fmt.Println("Done!")
			break
		}
		symbols, ok := m.commandBase[state]
		if !ok {
			log.Fatal("non-existent state: ", state)
		}
		currentSym := m.tape[m.headIndex]
		cmd := symbols[currentSym]
		fmt.Println(cmd.String())
		m.tape[m.headIndex] = cmd.symAfter
		switch cmd.shiftDirection {
		case 'R':
			m.headIndex++
			if m.headIndex >= TAPE_SIZE {
				log.Fatal("head out of range")
			}
		case 'L':
			m.headIndex--
			if m.headIndex < 0 {
				log.Fatal("head out of range")
			}
		}
		state = cmd.stateAfter
	}

}

func (m *Machine) GetTapeState() string {
	headPos := make([]byte, TAPE_SIZE)
	headPos[m.headIndex] = '^'
	return fmt.Sprintf("%s\n%s", string(m.tape), string(headPos))
	/*
	var firstTapeSymIndex int
	var lastTapeSymIndex int
	for i, sym := range m.tape {
		if sym != '_' {
			firstTapeSymIndex = i
			break
		}
	}
	for i := len(m.tape) - 1; i >= firstTapeSymIndex; i-- {
		if m.tape[i] != '_' {
			lastTapeSymIndex = i
			break
		}
	}
	str1startSpacesLen := firstTapeSymIndex - m.headIndex
	if str1startSpacesLen < 0 {
		str1startSpacesLen = 0
	}
	str1startSpaces := make([]byte, str1startSpacesLen)
	for i := range str1startSpaces {
		str1startSpaces[i] = ' '
	}
	str1endSpacesLen := m.headIndex - lastTapeSymIndex
	if str1endSpacesLen < 0 {
		str1endSpacesLen = 0
	}
	str1endSpaces := make([]byte, str1endSpacesLen)
	for i := range str1endSpaces {
		str1endSpaces[i] = ' '
	}
	str1 := fmt.Sprintf("%s%s%s", string(str1startSpaces),
		string(m.tape[firstTapeSymIndex:lastTapeSymIndex+1]),
		string(str1endSpaces))

	headPos := make([]byte, len(str1))
	headPos[m.headIndex-firstTapeSymIndex+str1startSpacesLen] = '^'
	return fmt.Sprintf("%s\n%s", str1, string(headPos))
	*/
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
