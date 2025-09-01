package ntpclient

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetCurrentTime() (time.Time, error) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	time := time.Now().Add(response.ClockOffset)
	if err != nil {
		return time, err
	}

	return time, nil
}

func Run() int {
	currentTime, err := GetCurrentTime()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error:", err)
		return 1
	}

	fmt.Println(currentTime)
	return 0
}
