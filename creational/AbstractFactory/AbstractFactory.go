/*
   Extension of factory pattern, This one is so confusing that i cant even think of any quip on it.
   If I were to compare this to regular factory, difference is this one is interface of object, not method.
   You can view this as the format a factory must follow in order to become multipurpose and not just wrapper.
   This pattern is the peak implementation of abstraction, there is nothing concrete in abstract factory.
*/

package main

import "fmt"

func main() {
	//Suppose this is client code
	flFactory := GetLogFactory("File")
	fl := flFactory.CreateLogger()
	fl.Log("Hello World")
	flc := flFactory.CreateLogConnection()
	flc.Open()
	flc.Open()
	flc.Close()

	clFactory:= GetLogFactory("Console")
	cl := clFactory.CreateLogger()
	cl.Log("Hmm")
	clc := clFactory.CreateLogConnection()
	clc.Close()
	clc.Open()
	clc.Close()
}

func GetLogFactory(factoryType string) ILogFactory {
	//Miss enum so much
	switch factoryType {
	case "File":
		return &FileLogFactory{}
	case "Console":
		return &ConsoleLogFactory{}
	default:
		return &FileLogFactory{}
	}
}

type ILogFactory interface {
	CreateLogger() ILogger
	CreateLogConnection() ILogConnection
}

type FileLogFactory struct {}

func (f *FileLogFactory) CreateLogger() ILogger {
	return &FileLogger{}
}

func (f *FileLogFactory) CreateLogConnection() ILogConnection {
	return &FileLogConnection{}
}

type ConsoleLogFactory struct {}

func (c *ConsoleLogFactory) CreateLogger() ILogger {
	return &ConsoleLogger{}
}

func (c *ConsoleLogFactory) CreateLogConnection() ILogConnection {
	return &ConsoleLogConnection{}
}

type ILogger interface {
	Log(message string)
}

type FileLogger struct {}

func (f *FileLogger) Log(message string) {
	fmt.Println("Data logged in File")
}

type ConsoleLogger struct {}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Data printed on Console")
}

type ILogConnection interface {
	Open()
	Close()
}

type FileLogConnection struct {
	connection int
}

func (f *FileLogConnection) Open() {
	if f.connection == 0 {
		f.connection = 1
		fmt.Println("Connection Opened for file logger")
	} else {
		fmt.Println("Connection is already made for file logger")
	}
}

func (f *FileLogConnection) Close() {
	if f.connection == 0 {
		fmt.Println("No connection present for file logger")
	} else {
		f.connection = 0
		fmt.Println("Connection Closed for file logger")
	}
}

type ConsoleLogConnection struct {
	connection int
}

func (c *ConsoleLogConnection) Open() {
	if c.connection == 0 {
		c.connection = 1
		fmt.Println("Connection Opened for console logger")
	} else {
		fmt.Println("Connection is already made for console logger")
	}
}

func (c *ConsoleLogConnection) Close() {
	if c.connection == 0 {
		fmt.Println("No connection present for console logger")
	} else {
		c.connection = 0
		fmt.Println("Connection Closed for console logger")
	}
}