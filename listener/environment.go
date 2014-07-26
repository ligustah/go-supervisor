package listener

import (
	supervisor "github.com/Ligustah/go-supervisor"
	"os"
)

var (
	SupervisorProcessName = os.Getenv("SUPERVISOR_PROCESS_NAME")
	SupervisorGroupName   = os.Getenv("SUPERVISOR_GROUP_NAME")
	SupervisorURL         = os.Getenv("SUPERVISOR_SERVER_URL")
	SupervisorEnabled     = os.Getenv("SUPERVISOR_ENABLED")
)

func GetSupervisor() supervisor.Supervisor {
	if SupervisorEnabled != "1" {
		panic("Must be started by supervisor")
	}
	return supervisor.New(SupervisorURL, nil)
}
