package listener

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type Header struct {
	Version    string `mapstructure:"ver"`
	Server     string
	Serial     int
	Pool       string
	PoolSerial int
	EventName  string
	Len        int
}

type ProcessState struct {
	ProcessName string
	GroupName   string
	FromState   string `mapstructure:"from_state"`
	Pid         int
	Expected    bool
	Tries       int
}

type ProcessLog struct {
	ProcessName string
	GroupName   string
	Pid         int
	Data        string
}

type RemoteCommunication struct {
}

type ProcessCommunication struct {
}

type Tick struct {
	When int
}

type Result string

const (
	RESULT_FAIL Result = "FAIL"
	RESULT_OK   Result = "OK"
)

func (r Result) String() string {
	return fmt.Sprintf("RESULT %d\n%s", len(string(r)), string(r))
}

func parseTokenData(data string, into interface{}) (err error) {
	m := make(map[string]string)
	tuples := strings.Split(strings.TrimSpace(data), " ")
	for _, tuple := range tuples {
		keyValue := strings.SplitN(tuple, ":", 2)
		if len(keyValue) != 2 {
			err = errors.New("Malformed token data")
			return
		}

		m[keyValue[0]] = keyValue[1]
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           into,
		WeaklyTypedInput: true,
	})

	if err != nil {
		return
	}

	return decoder.Decode(m)

}

func parseProcessState(data string) (ps ProcessState, err error) {
	err = parseTokenData(data, &ps)
	return
}

func parseTick(data string) (t Tick, err error) {
	err = parseTokenData(data, &t)
	return
}

func parseHeader(line string) (hdr Header, err error) {
	err = parseTokenData(line, &hdr)
	return
}
