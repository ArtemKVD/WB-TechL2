package main

import (
	"ntpclient/ntpclient"
	"os"
)

func main() {
	os.Exit(ntpclient.Run())
}
