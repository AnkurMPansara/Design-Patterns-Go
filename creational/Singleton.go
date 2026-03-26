/*
   Singleton pattern, it is the most abused pattern to the point that people call it anti-pattern.
   But it is absolutely necessary pattern to use when you dont want to create new object every time you call it.
   In nutshell, this is basically glorified global variable.
*/

package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	reader1 := getFileReader()
	fmt.Println(reader1.readFile("README.md"))
	reader2 := getFileReader()
	fmt.Println(reader2.readFile("README.md"))
}

var once sync.Once
var instance *FileReader

type FileReader struct {
	sync.RWMutex
	fileMap map[string]string
}

func getFileReader() *FileReader {
	once.Do(func() {
		fmt.Println("Instance initiated")
		instance = &FileReader{
			fileMap: make(map[string]string),
		}
	})
	return instance
}

func (f *FileReader) loadFile(filePath string) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Its not really my fault if path was wrong")
	}
	f.fileMap[filePath] = string(fileData)
}

func (f *FileReader) readFile(filePath string) string {
	fileContent := ""
	f.RLock()
	fileContent, exists := f.fileMap[filePath]; 
	f.RUnlock()
	if exists {
		fmt.Println("file Already present in reader");
		return fileContent
	}
	f.Lock()
	defer f.Unlock()
	if fileContent, exists = f.fileMap[filePath]; exists {
		return fileContent
	}
	f.loadFile(filePath)
	return f.fileMap[filePath]
}

func (f *FileReader) closeRead(filePath string) bool {
	f.Lock()
	defer f.Unlock()
	if _, ok := f.fileMap[filePath]; ok {
		delete(f.fileMap, filePath)
		return true
	}
	return false
}