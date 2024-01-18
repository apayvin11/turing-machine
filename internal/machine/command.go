package machine

import (
	"fmt"
	"strconv"
	"strings"
)

const END_STATE = "qz"

type Command struct {
	stateBefore    string
	symBefore      byte
	stateAfter     string
	symAfter       byte
	shiftDirection byte
}

func parseCommand(prototype string) (*Command, error) {
	params := strings.Split(prototype, ",")
	if len(params) != 5 {
		return nil, fmt.Errorf("invalid command: %s", prototype)
	}
	if err := validateState(params[0]); err != nil {
		return nil, fmt.Errorf("command parsing error: %s", err.Error())
	}
	if len(params[1]) != 1 {
		return nil, fmt.Errorf("command parsing error, invalid src sym %s", params[1])
	}
	if err := validateState(params[2]); err != nil {
		return nil, fmt.Errorf("command parsing error: %s", err.Error())
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
	return &Command{
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
		if state == END_STATE {
			return nil
		}
		return fmt.Errorf("invalid command state %s", state)
	}
	return nil
}

func(c *Command) String()string {
	return fmt.Sprintf("%s %c  -->  %s %c %c", 
	c.stateBefore,
	c.symBefore,
	c.stateAfter,
	c.symAfter,
	c.shiftDirection)
}