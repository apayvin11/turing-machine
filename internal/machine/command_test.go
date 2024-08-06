package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseCommand(t *testing.T) {
	testCases := []struct {
		name     string
		cmd      string
		isValid  bool
		expected func()command
	}{
		{
			name:     "valid R",
			cmd:      "q0,1,q1,_,R",
			expected: func() command {
				return TestCommand(t)
			},
			isValid:  true,
		},
		{
			name:     "valid L",
			cmd:      "q0,1,q1,_,L",
			expected: func() command {
				cmd := TestCommand(t)
				cmd.shiftDirection = 'L'
				return cmd
			},
			isValid:  true,
		},
		{
			name:     "valid E",
			cmd:      "q0,1,q1,_,E",
			expected: func() command {
				cmd := TestCommand(t)
				cmd.shiftDirection = 'E'
				return cmd
			},
			isValid:  true,
		},
		{
			name:     "invalid shift A",
			cmd:      "q0,1,q1,_,A",
			isValid:  false,
		},
		{
			name:     "invalid len",
			cmd:      "q0,q2,1,q1,_,R",
			isValid:  false,
		},
		{
			name:     "invalid src sym",
			cmd:      "q0,21,q1,_,R",
			isValid:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd, err := parseCommand(tc.cmd)
			if tc.isValid {
				assert.NoError(t, err)
				assert. Equal(t, tc.expected(), *cmd)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_validateState(t *testing.T) {
	testCases := []struct {
		name    string
		state   string
		isValid bool
	}{
		{
			name:    "valid",
			state:   "q0",
			isValid: true,
		},
		{
			name:    "invalid len",
			state:   "q",
			isValid: false,
		},
		{
			name:    "invalid format",
			state:   "d0",
			isValid: false,
		},
		{
			name:    "invalid format",
			state:   "q1d",
			isValid: false,
		},
		{
			name:    "valid end state",
			state:   endState,
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, validateState(tc.state))
			} else {
				assert.Error(t, validateState(tc.state))
			}
		})
	}
}
