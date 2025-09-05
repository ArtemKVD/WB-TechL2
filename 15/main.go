package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			fmt.Println("Exit")
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) > 1 {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Printf("cd error: %v", err)
				}
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Printf("pwd error: %v", err)
			} else {
				fmt.Println(dir)
			}
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) > 1 {
				pid, err := strconv.Atoi(args[1])

				if err != nil {
					fmt.Println("invalid PID")
				} else {
					process, err := os.FindProcess(pid)
					if err != nil {
						fmt.Printf("kill error: %v", err)
					} else {
						err := process.Kill()
						if err != nil {
							fmt.Printf("kill error: %v", err)
						}
					}
				}
			}
		case "ps":
			processes, err := ps.Processes()
			if err != nil {
				fmt.Printf("ps error: %v", err)
				break
			}

			fmt.Println("PID", "PPID", "COMMAND")
			for _, process := range processes {
				fmt.Println(process.Pid(), process.PPid(), process.Executable())
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}
