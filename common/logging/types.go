package logging

type (
	//ILogger logging wrapper
	ILogger interface {
		InfoLn(...interface{}) //INFO
		Infof(string, ...interface{})
		Error(error, ...interface{})
		Errorf(error, string, ...interface{})
		Panic(...interface{})
		Fatal(...interface{})
		Debugln(...interface{})
		Debugf(string, ...interface{})
	}

	//Fields type for standardize custom field input
	Fields map[string]interface{}
)

//Logger custom log
type Logger struct {
	fields Fields
}

var (
	logger Logger
)
