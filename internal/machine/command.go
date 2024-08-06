package machine

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	startState = "q0"
	endState   = "qz"
)

type command struct {
	stateBefore    string
	symBefore      byte
	stateAfter     string
	symAfter       byte
	shiftDirection byte
}

// parseCommand parses the passed command prototype
// command prototype format, parameters separated by commas: <qi>,<aj>,<qi*>,<aj*>,<dk>
// qi и qi* - machine state before and after command execution
// aj и aj* - observed character in a cell before and after execution
// dk - symbol indicating the direction of head shift (R, L, E)
// command prototype example: "q0,1,q1,_,R"
func parseCommand(prototype string) (*command, error) {
	params := strings.Split(prototype, ",")
	if len(params) != 5 {
		return nil, fmt.Errorf("invalid command: %s", prototype)
	}
	if err := validateState(params[0]); err != nil {
		return nil, fmt.Errorf("command parsing error: %v", err)
	}
	if len(params[1]) != 1 {
		return nil, fmt.Errorf("command parsing error, invalid src sym %s", params[1])
	}
	if err := validateState(params[2]); err != nil {
		return nil, fmt.Errorf("command parsing error: %v", err)
	}
	if len(params[3]) != 1 {
		return nil, fmt.Errorf("command parsing error, invalid rec sym %s", params[3])
	}
	if len(params[4]) != 1 {
		return nil, fmt.Errorf("command parsing error, invalid shift direction sym %s", params[4])
	}
	switch params[4] {
	case "R", "L", "E":
	default:
		return nil, fmt.Errorf("command parsing error, invalid shift direction sym %s", params[4])
	}
	return &command{
		stateBefore:    params[0],
		symBefore:      params[1][0],
		stateAfter:     params[2],
		symAfter:       params[3][0],
		shiftDirection: params[4][0],
	}, nil
}

func validateState(state string) error {
	if len(state) < 2 {
		return fmt.Errorf("invalid command state %s", state)
	}
	if state[0] != 'q' {
		return fmt.Errorf("invalid command state %s", state)
	}
	if _, err := strconv.Atoi(state[1:]); err != nil {
		if state == endState {
			return nil
		}
		return fmt.Errorf("invalid command state %s", state)
	}
	return nil
}

// String returns a string with command info for debug
func (c *command) String() string {
	return fmt.Sprintf("%s %c  -->  %s %c %c\n",
		c.stateBefore,
		c.symBefore,
		c.stateAfter,
		c.symAfter,
		c.shiftDirection)
}
