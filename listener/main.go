package listener

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ProcessStatePrefix          = "PROCESS_STATE"
	RemoteCommunicationPrefix   = "REMOTE_COMMUNICATION"
	ProcessLogPrefix            = "PROCESS_LOG"
	SupervisorStateChangePrefix = "SUPERVISOR_STATE_CHANGE"
	TickPrefix                  = "TICK"
)

type ProcessStateHandler interface {
	handleProcessState(Header, ProcessState) Result
}

type ProcessStateHandlerFunc func(Header, ProcessState) Result

func (p ProcessStateHandlerFunc) handleProcessState(h Header, ps ProcessState) Result {
	return p(h, ps)
}

type RemoteCommunicationHandler interface {
	handleRemoteCommunication(Header, RemoteCommunication) Result
}

type ProcessLogHandler interface {
}

type SupervisorStateChangeHandler interface {
}

type TickHandler interface {
}

type Listener struct {
	ProcessStateHandler
	RemoteCommunicationHandler
}

func (l *Listener) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("READY")

		//read string until end of line
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return err
		}

		//parse header line
		hdr, err := parseHeader(line)
		if err != nil {
			log.Println(err)
			return err
		}

		//read payload
		payload := make([]byte, hdr.Len)
		n, err := reader.Read(payload)
		if err != nil || n != len(payload) {
			log.Printf("Failed to read %d bytes from payload (n=%d, err=%v)", hdr.Len, n, err)
			break
		}

		fmt.Print(l.handle(hdr, payload))
	}

	log.Printf("Exiting reader loop")

	return nil
}

func (l *Listener) handle(h Header, payload []byte) Result {
	if strings.HasPrefix(h.EventName, ProcessStatePrefix) && l.ProcessStateHandler != nil {
		ps, err := parseProcessState(string(payload))
		if err != nil {
			//FIXME: should probably bail out here because supervisor will re-queue this event
			// this might end up being an infinite loop
			return RESULT_FAIL
		}
		return l.ProcessStateHandler.handleProcessState(h, ps)
	} else if strings.HasPrefix(h.EventName, RemoteCommunicationPrefix) {
		//TODO: implement
	} else if strings.HasPrefix(h.EventName, ProcessLogPrefix) {
		//TODO: implement
	} else if strings.HasPrefix(h.EventName, SupervisorStateChangePrefix) {
		//TODO: implement
	} else if strings.HasPrefix(h.EventName, TickPrefix) {
		//TODO: implement
	}

	return RESULT_OK
}
