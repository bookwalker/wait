package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func wait(address string, timeout time.Duration) error {
	start := time.Now()
	for ;; {
		conn, _ := net.DialTimeout("tcp", address, timeout)
		if conn != nil {
			defer conn.Close()
			fmt.Println(fmt.Sprintf("Connection to [%s] established after %d seconds", address, time.Since(start) / time.Second))
			return nil
		} else {
			fmt.Println(fmt.Sprintf("Connect to [%s] failed. Retrying in 1 sec", address))
			time.Sleep(1 * time.Second)
			if timeout != 0 {
				timeElapsed := time.Since(start)
				if timeElapsed >= timeout {
					return errors.New(fmt.Sprintf("Could not connect to [%s] within timeout %d seconds", address, timeout / time.Second))
				}
			}
		}
	}
}

func main() {
	address := flag.String("address", "", "a hostname or ip-address including port. Eg. 127.0.0.1:80")
	timeout := flag.Int("timeout", 10, "a timeout in seconds or 0 for no timeout")
	flag.Parse()

	if len(*address) > 0 && !strings.Contains(*address, ":") {
		_, _ = fmt.Fprintln(os.Stderr, "port not set")
		flag.Usage()
		os.Exit(1)
	}

	if len(*address) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "address not set")
		flag.Usage()
		os.Exit(1)
	}

	err := wait(*address, time.Duration(*timeout) * time.Second)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("%v", err))
		os.Exit(1)
	}
}
