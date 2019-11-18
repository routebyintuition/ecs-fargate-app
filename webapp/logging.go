package main

import (
	"io"
	"log"
	"os"
)

var (
	// Info is the info log level for the app
	Info *log.Logger

	// Warning is the warn log level for the app
	Warning *log.Logger

	// Error is the error log level for the app
	Error *log.Logger
)

// Simple logging method
func logInit() {

	infoHandle := io.MultiWriter(os.Stdout)
	warningHandle := io.MultiWriter(os.Stdout)
	errorHandle := io.MultiWriter(os.Stdout)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime)
}
