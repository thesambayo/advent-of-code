package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	absoluteInputDirPath = "../../../.inputs"
)

// getInputFromCwdPath uses the current working directory main.go `<repo>/go/<year>/day<number>/main.go`
// to get the file path for the input file
// which is expected to the in the `<repo>/.inputs/<year>/day<number>.txt` format
func getInputFromCwdPath() (string, error) {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return "", errors.New("Could not get current working directory")
	}

	filePathInfo := strings.Split(currentWorkingDirectory, "/")
	inputInfo := filePathInfo[len(filePathInfo)-2:]
	inputPath := fmt.Sprintf("%v/%v.txt", absoluteInputDirPath, strings.Join(inputInfo, "/"))
	return inputPath, nil
}

// readFile uses file path to read a file and returns the file to be used as desired.
// It is advised to call closeFile anywhere readFile is called
func readFile(filepath string) (*os.File, func(), error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, errors.New("Unable to read file")
	}

	closeFile := func() {
		defer file.Close()
	}
	return file, closeFile, nil
}

func ReadOneLineFile() (string, error) {
	var inputLine string
	filepath, err := getInputFromCwdPath()
	if err != nil {
		return inputLine, err
	}

	file, closeFile, err := readFile(filepath)
	defer closeFile()
	if err != nil {
		return inputLine, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine = scanner.Text()
	}
	return inputLine, nil
}

// ReadMultiLinesFile reads and returns the input for cwd as slice of string
func ReadMultiLinesFile() ([]string, error) {
	inputSlice := make([]string, 0)
	filepath, err := getInputFromCwdPath()
	if err != nil {
		return inputSlice, err
	}

	file, closeFile, err := readFile(filepath)
	defer closeFile()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputSlice = append(inputSlice, scanner.Text())
	}
	return inputSlice, nil
}

// ReadMultiLinesFile reads and returns the input for cwd as slice of string
func ReadMultiColumnsFile() ([][]string, error) {
	inputSlice := make([][]string, 0)
	filepath, err := getInputFromCwdPath()
	if err != nil {
		return nil, err
	}

	file, closeFile, err := readFile(filepath)
	defer closeFile()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	// Initialize slices for the columns
	var column1 []string
	var column2 []string
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line) // strings.Fields splits the string by whitespace
		column1 = append(column1, columns[0])
		column2 = append(column2, columns[1])
	}
	inputSlice = append(inputSlice, column1, column2)

	return inputSlice, nil
}

func ReadMultiColumnsIntFile() ([][]int, error) {
	inputSlice := make([][]int, 0)
	filepath, err := getInputFromCwdPath()
	if err != nil {
		return nil, err
	}

	file, closeFile, err := readFile(filepath)
	defer closeFile()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	// Initialize slices for the columns
	var column1 []int
	var column2 []int
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line) // strings.Fields splits the string by whitespace
		columnItem1, _ := strconv.Atoi(columns[0])
		columnItem2, _ := strconv.Atoi(columns[1])
		column1 = append(column1, columnItem1)
		column2 = append(column2, columnItem2)
	}
	inputSlice = append(inputSlice, column1, column2)

	return inputSlice, nil
}

// ----------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------
// MultipleLinesNumbers returns the input file if it contains multiple lines of text which are numbers
// func MultipleLinesNumbers(fileName string) ([]int, error) {
// 	inputSlice := make([]int, 0)
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return inputSlice, err
// 	}
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		thisNumber, _ := strconv.Atoi(scanner.Text())
// 		inputSlice = append(inputSlice, thisNumber)
// 	}
// 	err = file.Close()
// 	if err != nil {
// 		return inputSlice, err
// 	}
// 	return inputSlice, nil
// }
//
//
//
//
