package listener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHeader(t *testing.T) {
	testInput := "ver:3.0 server:supervisor serial:21 pool:listener poolserial:10 eventname:PROCESS_COMMUNICATION_STDOUT len:54"
	hdr, err := parseHeader(testInput)
	assert.NoError(t, err)

	assert.Equal(t, "3.0", hdr.Version)
	assert.Equal(t, "supervisor", hdr.Server)
	assert.Equal(t, 21, hdr.Serial)
	assert.Equal(t, "listener", hdr.Pool)
	assert.Equal(t, 10, hdr.PoolSerial)
	assert.Equal(t, "PROCESS_COMMUNICATION_STDOUT", hdr.EventName)
	assert.Equal(t, 54, hdr.Len)
}

func TestParseProcessState(t *testing.T) {
	testInput := "processname:cat groupname:cat from_state:STOPPED tries:0 expected:1 pid:2456"
	ps, err := parseProcessState(testInput)
	assert.NoError(t, err)

	assert.Equal(t, "cat", ps.ProcessName)
	assert.Equal(t, "cat", ps.GroupName)
	assert.Equal(t, "STOPPED", ps.FromState)
	assert.Equal(t, 0, ps.Tries)
	assert.Equal(t, true, ps.Expected)
	assert.Equal(t, 2456, ps.Pid)
}
