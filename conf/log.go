package conf

type LogFormat string

const (
	TextFormat = LogFormat("text")
	JSONFormat = LogFormat("json")
)

type LogTo string

const (
	ToFile   = LogTo("file")
	ToStdout = LogTo("stdout")
)
