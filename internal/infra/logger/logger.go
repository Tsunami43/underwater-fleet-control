package logger

import (
	"log"
	"os"
)

type Logger struct {
	file *os.File
}

func NewLogger(filePath string) (*Logger, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{file: f}, nil
}

func (l *Logger) Log(msg string) {
	log.SetOutput(l.file)
	log.Println(msg)
}

func (l *Logger) Close() error {
	return l.file.Close()
}

