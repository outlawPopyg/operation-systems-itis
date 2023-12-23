package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	producerOut, producerIn, _ := os.Pipe()
	bcIn, bcInController, _ := os.Pipe()
	bcOutController, bcOut, _ := os.Pipe()
	counter := 0
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGUSR1)
	go func() {
		for {
			_ = <-sigc
			fmt.Fprintf(os.Stderr, "Produced: %v\n", counter)
		}
	}()
	producerPID, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if producerPID == 0 {
		syscall.Dup2(int(producerIn.Fd()), 1)
		syscall.Exec("./producer", []string{}, []string{})
	}
	bcPID, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if bcPID == 0 {
		syscall.Dup2(int(bcIn.Fd()), 0)
		syscall.Dup2(int(bcOut.Fd()), 1)
		syscall.Exec("/usr/bin/bc", []string{}, []string{})
	}

	var str = make([]byte, 6)
	read, _ := syscall.Read(int(producerOut.Fd()), str)

	for read > 1 {
		in := string(str[0:5])
		syscall.Write(int(bcInController.Fd()), str)

		str = make([]byte, 6)
		syscall.Read(int(bcOutController.Fd()), str)
		fmt.Printf("%v = %v", in, string(str))
		counter++
		read, _ = syscall.Read(int(producerOut.Fd()), str)
	}
}
