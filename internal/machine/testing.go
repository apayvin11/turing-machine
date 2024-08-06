package machine

import "testing"

func TestCommand(t *testing.T) command {
	return command{
		stateBefore:    "q0",
		symBefore:      '1',
		stateAfter:     "q1",
		symAfter:       '_',
		shiftDirection: 'R',
	}
}

func TestCommandsProto(t *testing.T) []string {
	return []string{
		"q0,1,q1,_,R",
		"q1,1,q1,1,R",
		"q1,*,q2,*,R",
		"q2,1,q3,a,R",
		"q3,1,q3,1,R",
		"q3,=,q3,=,R",
		"q3,-,q3,-,R",
		"q3,_,q4,1,L",
		"q4,1,q4,1,L",
		"q4,a,q2,a,R",
		"q4,=,q4,=,L",
		"q4,-,q5,-,L",
		"q5,1,q4,1,L",
		"q5,a,q5,1,L",
		"q5,*,q6,*,L",
		"q6,1,q7,1,L",
		"q6,_,q8,_,R",
		"q7,1,q7,1,L",
		"q7,_,q0,_,R",
		"q8,1,q8,_,R",
		"q8,*,q8,_,R",
		"q8,-,q9,_,R",
		"q9,1,q10,_,R",
		"q9,=,qz,_,E",
		"q10,1,q10,1,R",
		"q10,=,q10,=,R",
		"q10,_,q11,_,L",
		"q11,1,q12,_,L",
		"q11,=,q12,=,L",
		"q12,1,q12,1,L",
		"q12,=,q12,=,L",
		"q12,_,q9,_,R",
	}
}
