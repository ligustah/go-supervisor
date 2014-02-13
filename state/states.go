package state

const (
	STOPPED  int64 = 0
	STARTING int64 = 10
	RUNNING  int64 = 20
	BACKOFF  int64 = 30
	STOPPING int64 = 40
	EXITED   int64 = 100
	FATAL    int64 = 200
	UNKNOWN  int64 = 1000
)
