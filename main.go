package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	dir = flag.String("dir", "", "Directory to watch")
)

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetAvailableName(path string) string {
	var suffix string
	var extension string

	_, err := os.Stat(path)

	if err != nil {
		return path
	}

	for i := 0; err == nil; i++ {
		splitted_path := strings.Split(path, ".")
		path = splitted_path[0]
		if len(splitted_path) > 1 {
			extension = "." + splitted_path[1]
		} else {
			extension = ""
		}
		suffix = fmt.Sprintf(" (%d)", i)
		_, err = os.Stat(path + suffix + extension)

	}
	return path + suffix + extension
}

func Rename(path string) {
	splitted_path := strings.Split(path, "/")
	file_name := splitted_path[len(splitted_path)-1]

	splitted_filename := strings.Split(file_name, "_")

	// We Recover the filename, the space in the filename have been replaced by "_",
	// so we need to check where the filename ends

	if len(splitted_filename) <= 4 {
		return
	}
	file_name = splitted_filename[4]
	for i := 5; i < len(splitted_filename); i++ {
		if splitted_filename[i][0] == '{' {
			break
		}
		file_name += "_" + splitted_filename[i]
	}

	splitted_path[len(splitted_path)-1] = file_name
	new_path := strings.Join(splitted_path, "/")

	new_path = GetAvailableName(new_path)

	err := os.Rename(path, new_path)
	handleErr(err)
}

func WalkDir(path string) {
	files, err := os.ReadDir(path)
	handleErr(err)
	for _, file := range files {
		if file.Type().IsDir() {
			WalkDir(path + "/" + file.Name())
		} else {
			Rename(path + "/" + file.Name())
		}
	}
}

func main() {
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		return
	}

	WalkDir(*dir)
}
