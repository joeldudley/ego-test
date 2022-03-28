package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	dir      = "./enclave"
	filename = "enclave_data.txt"
)

func main() {
	var err error
	var f *os.File
	var contents []byte
	var value uint64

	path := fmt.Sprintf("%s/%s", dir, filename)

	// We create the file and its parent directories.
	check(os.MkdirAll(dir, 0770))
	f, err = os.OpenFile(path, os.O_CREATE, 0644)
	f.Close()
	check(err)

	// We read the current value stored in the file.
	contents, err = os.ReadFile(path)
	check(err)

	// We write back the value, incremented by one.
	value, err = strconv.ParseUint(string(contents), 10, 64)
	updatedValue := []byte(strconv.FormatUint(value+1, 10))
	check(os.WriteFile(path, updatedValue, 0644))

	// We read back and print the value.
	contents, err = os.ReadFile(path)
	check(err)
	println(fmt.Sprintf("jjj file contents: %s", contents))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
