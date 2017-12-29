package log

func Change(strLevel string, pathname string, flag int) *Logger {
	newLogger := make(chan *Logger)
	go func() {
		logger, err := New(strLevel, pathname, flag)
		if err != nil {
			logger = nil
			return
		}
		newLogger <- logger
	}()
	return <-newLogger
}
