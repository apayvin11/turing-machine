package machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Machine(t *testing.T) {
	alphabetSlice := []string{"1 a * = - _"}
	tapeSlice := []string{"11*11-1="}
	commandsProto := TestCommandsProto(t)
	mockPrintFunc := func(string) error { return nil }
	machine, err := newMachine(alphabetSlice, tapeSlice, commandsProto, mockPrintFunc)
	require.NoError(t, err)
	require.NoError(t, machine.Run())
	expectedTape := make([]byte, tapeSize)
	for i := range expectedTape {
		expectedTape[i] = '_'
	}
	copy(expectedTape[18:], "111")
	assert.Equal(t, string(expectedTape), string(machine.tape[:]))
}
