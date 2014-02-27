package config

import (
	"io"
	"text/template"
)

//define program options
const (
	Command               = "command"
	ProcessName           = "process_name"
	NumProcs              = "numprocs"
	Directory             = "directory"
	UMask                 = "umask"
	Priority              = "priority"
	AutoStart             = "autostart"
	AutoRestart           = "autorestart"
	StartSecs             = "startsecs"
	StartRetries          = "startretries"
	ExitCodes             = "exitcodes"
	StopSignal            = "stopsignal"
	StopWaitSecs          = "stopwaitsecs"
	User                  = "user"
	RedirectStderr        = "redirect_stderr"
	StdoutLogfile         = "stdout_logfile"
	StdoutLogfileMaxBytes = "stdout_logfile_maxbytes"
	StoudtLogfileBackups  = "stdout_logfile_backups"
	StdoutCaptureMaxBytes = "stdout_capture_maxbytes"
	StderrLogfile         = "stderr_logfile"
	StderrLogfileMaxBytes = "stderr_logfile_maxbytes"
	StderrLogfileBackups  = "stderr_logfile_backups"
	StderrCaptureMaxBytes = "stderr_capture_maxbytes"
	Environment           = "environment"
	ServerURL             = "serverurl"
)

const (
	programTemplateText string = `[program:{{ .Name }}]
{{ range $name, $value := .Options }}{{ $name }} = {{ $value }}
{{ end }}`
)

var (
	programTemplate = template.Must(template.New("supervisorProgram").Parse(programTemplateText))
)

func GenerateProgramConfig(name string, options map[string]string, out io.Writer) error {
	data := struct {
		Name    string
		Options map[string]string
	}{
		name, options,
	}

	return programTemplate.Execute(out, data)
}
