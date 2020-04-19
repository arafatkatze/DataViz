// Package utils provides common utility functions.
//
// Provided functionalities:
// - sorting
// - comparators
package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(value.(int64)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(value.(uint64)), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value.(float64)), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return fmt.Sprintf("%+v", value)
	}
}

// WriteDotStringToPng takes a content of a dot file in a string and makes a graph using Graphviz
// to ouput an image
func WriteDotStringToPng(fileName string, dotFileString string) (ok bool) {
	byteString := []byte(dotFileString) // Converting the string to byte slice to write to a file
	tmpFile, _ := ioutil.TempFile("", "TemporaryDotFile")
	tmpFile.Write(byteString)            // Writing the string to a temporary file
	dotPath, err := exec.LookPath("dot") // Looking for dot command
	if err != nil {
		fmt.Println("Error: Running the Visualizer command. Please install Graphviz")
		return false
	}
	dotCommandResult, err := exec.Command(dotPath, "-Tpng", tmpFile.Name()).Output() // Running the command
	if err != nil {
		fmt.Println("Error: Running the Visualizer command. Please install Graphviz")
		return false
	}
	ioutil.WriteFile(fileName, dotCommandResult, os.FileMode(int(0777)))
	return true
}
