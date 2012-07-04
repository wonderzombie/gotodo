package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readNotesDir(dirName string) (files []string, err error) {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	files = make([]string, len(fileInfos))
	for i, value := range fileInfos {
		files[i] = value.Name()
	}

	return
}

func printFiles(files []string) {
	for _, f := range files {
		fmt.Printf("* %s\n", f)
	}
}

func find(needle string, haystack []string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func printFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", contents)
}

func appendToFile(filename string, line string) {
	return
}

func main() {
	default_path := filepath.Join(os.Getenv("HOME"), ".notes")

	files, err := readNotesDir(default_path)
	if err != nil {
		fmt.Printf("Unable to open %s: %s\n", default_path, err)
	}

	if len(os.Args) == 1 {
		printFiles(files)
		return
	} else {
		todoCommand := os.Args[1]
		fmt.Printf("Command was: %v\n", todoCommand)

		if find(todoCommand, files) {
			filename := filepath.Join(default_path, todoCommand)

			// Append or list?
			if len(os.Args) > 2 {
				// append remaining args to file
				line := strings.Join(os.Args[2:], " ")
				appendToFile(filename, line)
			} else {
				// list
				printFile(filename)
			}
		}
	}

}
