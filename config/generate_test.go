package config

import (
	"bytes"
	"testing"
)

func TestGenerateProgramConfig(t *testing.T) {
	expected := `[program:test]
command = echo "Hello world"
priority = 5
`
	buffer := new(bytes.Buffer)
	err := GenerateProgramConfig("test", map[string]string{
		Command:  `echo "Hello world"`,
		Priority: "5",
	}, buffer)

	if err != nil {
		t.Log("Unexpected error:", err)
		t.Fatal(err)
	}

	if buffer.String() != expected {
		t.Log("Return value does not match expected value")
		t.Fail()
	}
}
