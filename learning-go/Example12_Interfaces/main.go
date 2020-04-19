// Example 12a
package main

type Logger interface {
	Log(message string)
}

// we might have various types of loggers:
type SqlLogger struct{}
type ConsoleLogger struct{}
type FileLogger struct{}

// how to use one: just like any other type
// it could be a structure's field:
type Server struct {
	logger Logger
}

// or a function parameter (or return value)
func process(logger Logger) {
	logger.Log("Hello")
}

func main() {

}
