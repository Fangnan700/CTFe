package log

import (
	"io"
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitLogger() {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)

	InfoLogger = log.New(multiWriter, "<INFO>", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "<ERROR>", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("logger has been initialized")
}
