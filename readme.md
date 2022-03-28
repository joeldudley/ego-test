This projects shows how to create a read and write to a persistent sealed file with EGo.

# How it works

A counter is stored, sealed by the enclave, in `enclave/enclave_data.txt`. Because the counter is sealed, it cannot be 
read from outside the enclave. `main/enclave.json` configures this file as a mount for the enclave, meaning that it is 
persistent across enclave restarts.

The program in `main/main.go` reads the counter (creating the file and initialising the counter to zero if it doesn't 
exist), unseals and increments it, then seals and rewrites it to disk.

# How to run it

1. Update `main/enclave.json` so that the mount's source points to the `enclave/` directory in this repo
2. Step into the `main` folder: `cd main`
3. Build and sign the binary with EGo: `ego-go build && ego sign main`
4. Run the binary with EGo in simulation mode: `OE_SIMULATION=1 ego run main`
