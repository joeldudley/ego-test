package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/edgelesssys/ego/ecrypto"
)

const path = "./enclave/enclave_data.txt"

func main() {
	// We create the enclave data file, if it doesn't already exist.
	f, err := os.OpenFile(path, os.O_CREATE, 0o644)
	f.Close()
	check(err)

	// We read the current value stored in the file, defaulting to zero.
	sealedContents, err := os.ReadFile(path)
	check(err)
	var contents []byte
	if len(sealedContents) == 0 {
		contents = []byte("0")
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
	check(os.WriteFile(path, sealedContents, 0o644))

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
