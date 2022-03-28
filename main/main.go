package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("my enclave data")
	err := os.WriteFile("./enclave_data.txt", data, 0644)
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile("./enclave_data.txt")
	if err != nil {
		panic(err)
	}
	println(fmt.Sprintf("jjj file contents: %s", file))
}
