package logger

import (
	"fmt"
	"log"
)

// Fatal logs a fatal error and exits the program if any argument is non-nil
func Fatal(args ...interface{}) {
	for _, arg := range args {
		if err, ok := arg.(error); ok && err != nil {
			log.Fatal(args...)
		}
	}
}

func Println(msg string)  {
    log.Println(msg)
}


func Errorf(format string, args ...interface{}) error {
	for _, arg := range args {
		if err, ok := arg.(error); ok && err != nil {
			// Log the error
			log.Printf(format, args...)
			// Return the formatted error
			return fmt.Errorf(format, args...)
		}
	}
	return nil
}