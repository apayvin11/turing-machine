package machine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	TAPE_SIZE              = 48
	DEFAULT_HEAD_INDEX     = 10
	MAX_DEFAULT_HEAD_INDEX = TAPE_SIZE - 1
)

type Machine struct {
	alphabet    []string
	commandBase map[string]map[byte]*Command // база команд, вложенный map, key1 - state, key2 - source symbol
	tape        []byte
	headIndex   int
	resultFile  *os.File
}

// New создает и возвращает новый объект Машины Тьюринга
// input parrams:
// alphabetPath - путь к алфавиту (в нем должны быть в символы через пробел, в одну строку)
// tapePath - путь файлу с входной лентой
// commandsPath - путь к файлу с командами (каждая команда на отдельной строке)
// resultFile - путь к файлу для записи результата
func New(alphabetPath, tapePath, commandsPath string, resultFile *os.File) *Machine {
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
		resultFile:  resultFile,
	}
}

// Run запускаем машину Тьюринга
func (m *Machine) Run() {
	state := "q0"
	for {
		m.print(m.GetTapeState())
		if state == END_STATE {
			m.print("Done!")
			break
		}
		symbols, ok := m.commandBase[state]
		if !ok {
			log.Fatal("non-existent state: ", state)
		}
		currentSym := m.tape[m.headIndex]
		cmd, ok := symbols[currentSym]
		if !ok {
			log.Fatalf("non-existent source symbol %c for state %s", currentSym, state)
		}
		m.print(cmd.String())
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

// GetTapeState возвращает текущее состояние ленты с указанием положения головки
func (m *Machine) GetTapeState() string {
	headPos := make([]byte, TAPE_SIZE)
	headPos[m.headIndex] = '^'
	return fmt.Sprintf("%s\n%s\n", string(m.tape), string(headPos))
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

// print выводит переданную строку в stdout и в resultFile
func (m *Machine) print(str string) {
	fmt.Print(str)
	if _, err := m.resultFile.WriteString(str); err != nil {
		log.Fatal(err)
	}
}
