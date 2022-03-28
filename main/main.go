package main

import (
	"fmt"
	"github.com/edgelesssys/ego/ecrypto"
	"os"
	"strconv"
)

const (
	dir      = "./enclave"
	filename = "enclave_data.txt"
)

func main() {
	path := fmt.Sprintf("%s/%s", dir, filename)

	// We create the file and its parent directories.
	check(os.MkdirAll(dir, 0770))
	f, err := os.OpenFile(path, os.O_CREATE, 0644)
	f.Close()
	check(err)

	// We read the current value stored in the file, defaulting to zero.
	sealedContents, err := os.ReadFile(path)
	check(err)
	var contents []byte
	if len(sealedContents) == 0 {
		contents = []byte{}
	} else {
		contents, err = ecrypto.Unseal(sealedContents, nil)
		check(err)
	}

	// We increment the current value by one and write it back.
	value, err := strconv.ParseUint(string(contents), 10, 64)
	check(err)
	updatedValue := []byte(strconv.FormatUint(value+1, 10))
	sealedContents, err = ecrypto.SealWithUniqueKey(updatedValue, nil)
	check(err)
	check(os.WriteFile(path, sealedContents, 0644))

	// We read back and print the value.
	sealedContents, err = os.ReadFile(path)
	check(err)
	contents, err = ecrypto.Unseal(sealedContents, nil)
	println(fmt.Sprintf("jjj file contents: %s", contents))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
