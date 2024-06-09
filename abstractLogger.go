package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger interface {
	Log(message string) error
}

type ConsoleLogger struct {
	Writer io.Writer
}

func hellowsWorlds() string {
	return `Hello World!`
}
func (c *ConsoleLogger) Log(message string) error {
	_, err := fmt.Fprintln(c.Writer, message)
	return err
}

type FileLogger struct {
	File *os.File
}

func (f *FileLogger) Log(message string) error {
	_, err := fmt.Fprintln(f.File, message)
	return err
}

type RemoteLogger struct {
	Address string
}

func (r *RemoteLogger) Log(message string) error {
	// Имитация отправки сообщения на удаленный сервер
	fmt.Println("Sending message to remote server:", message)
	return nil
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

// func main(){
// 		//AbstractLogger
// 	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
// 	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	fileLogger := &FileLogger{File: file}
// 	remoteLogger := &RemoteLogger{Address: "http://example.com/logger"}

// 	loggers := []Logger{consoleLogger, fileLogger, remoteLogger}
// 	LogAll(loggers, "This is a test log message.")
// }
