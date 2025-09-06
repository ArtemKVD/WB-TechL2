package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	conn, err := connection(address, *timeout)
	if err != nil {
		fmt.Println("connection error: ", err)
	}
	defer conn.Close()

	done := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go read(conn, done)
	go write(conn)

	select {
	case <-done:
		fmt.Println("connection closed")
	case sig := <-sigCh:
		fmt.Println("recived signal: ", sig)
	}
}

func connection(address string, timeout time.Duration) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("connection error: ", err)
	}
	return conn, err
}

func read(conn net.Conn, done chan<- struct{}) {
	defer close(done)

	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error: ", err)
			}
			return
		}
		if n > 0 {
			os.Stdout.Write(buffer[:n])
		}
	}
}

func write(conn net.Conn) {

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(conn)

	for scanner.Scan() {
		text := scanner.Text()

		_, err := writer.WriteString(text + "\n")
		if err != nil {
			fmt.Println("write error: ", err)
			return
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("flush error: ", err)
			return
		}
	}
}
