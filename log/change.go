package log

import "fmt"

func Change(strLevel string, pathname string, flag int) *Logger {
	newLogger := make(chan *Logger)
	go func() {
		logger, err := New(strLevel, pathname, flag)
		if err != nil {
			fmt.Println(err)
			return
		}
		newLogger <- logger
	}()
	return <-newLogger
}
