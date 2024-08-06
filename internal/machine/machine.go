package machine

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/apayvin11/turing-machine/internal/util"
)

const (
	tapeSize            = 48
	defaultHeadIndex    = 10
	maxDefaultHeadIndex = tapeSize - 1
)

type Machine struct {
	alphabet    []string
	commandBase map[string]map[byte]*command // command base, nested map, key1 - state, key2 - source symbol
	tape        []byte
	headIndex   int
	resultFile  *os.File
}

// New creates and returns a new Turing Machine object
// input parrams:
// alphabetPath - path to the alphabet (it must contain characters separated by spaces, in one line)
// tapePath - path to the input tape file
// commandsPath - path to the file with commands (each command on a separate line)
// resultFile - path to the file to write the result
func New(alphabetPath, tapePath, commandsPath string, resultFile *os.File) (*Machine, error) {
	// Reading alphabet
	alphabetSlice, err := util.ReadFileIntoSliceByLines(alphabetPath)
	if err != nil {
		return nil, err
	}
	if l := len(alphabetSlice); l != 1 {
		return nil, fmt.Errorf("invalid alphabet file, expected 1 line, received: %d", l)
	}
	alphabet := strings.Fields(alphabetSlice[0])
	for _, sym := range alphabet {
		if len(sym) != 1 {
			return nil, fmt.Errorf("invalid alphabet file, invalid character found: %s", sym)
		}
	}
	fmt.Printf("alphabet: %s\n", alphabet)

	// Reading tape
	tapeSlice, err := util.ReadFileIntoSliceByLines(tapePath)
	if err != nil {
		return nil, err
	}
	if l := len(tapeSlice); l != 1 {
		return nil, fmt.Errorf("invalid tape file, expected 1 line, received: %d", l)
	}
	tapeStr := tapeSlice[0]
	fmt.Printf("tape: %s\n", tapeStr)
	tape := make([]byte, tapeSize)
	for i := range tape {
		tape[i] = '_'
	}
	copy(tape[defaultHeadIndex:], tapeStr)

	// Reading commands
	commandsProto, err := util.ReadFileIntoSliceByLines(commandsPath)
	if err != nil {
		return nil, err
	}
	commandBase := make(map[string]map[byte]*command)
	for _, proto := range commandsProto {
		cmd, err := parseCommand(proto)
		if err != nil {
			return nil, fmt.Errorf("command %s parsing error: %v", proto, err)
		}
		_, ok := commandBase[cmd.stateBefore]
		if !ok {
			commandBase[cmd.stateBefore] = map[byte]*command{}
		}
		commandBase[cmd.stateBefore][cmd.symBefore] = cmd
	}
	return &Machine{
		alphabet:    alphabet,
		tape:        tape,
		commandBase: commandBase,
		headIndex:   defaultHeadIndex,
		resultFile:  resultFile,
	}, nil
}

// Run launches a Turing machine
func (m *Machine) Run() error {
	state := startState
	for {
		if err := m.print(m.getTapeStateStr()); err != nil {
			return err
		}
		if state == endState {
			m.print("Done!")
			break
		}
		symbols, ok := m.commandBase[state]
		if !ok {
			return fmt.Errorf("non-existent state: %s", state)
		}
		currentSym := m.tape[m.headIndex]
		cmd, ok := symbols[currentSym]
		if !ok {
			return fmt.Errorf("non-existent source symbol %c for state %s", currentSym, state)
		}
		if err := m.print(cmd.String()); err != nil {
			return err
		}
		m.tape[m.headIndex] = cmd.symAfter
		switch cmd.shiftDirection {
		case 'R':
			m.headIndex++
			if m.headIndex >= tapeSize {
				return errors.New("head out of range")
			}
		case 'L':
			m.headIndex--
			if m.headIndex < 0 {
				return errors.New("head out of range")
			}
		}
		state = cmd.stateAfter
	}
	return nil
}

// getTapeStateStr returns the current state of the tape indicating the position of the head
func (m *Machine) getTapeStateStr() string {
	headPos := make([]byte, tapeSize)
	for i := range headPos {
		headPos[i] = '_'
	}
	headPos[m.headIndex] = '^'
	return fmt.Sprintf("%s\n%s\n", string(m.tape), string(headPos))
}

// print prints the passed string to stdout and to resultFile
func (m *Machine) print(str string) error {
	fmt.Print(str)
	_, err := m.resultFile.WriteString(str)
	return err
}
