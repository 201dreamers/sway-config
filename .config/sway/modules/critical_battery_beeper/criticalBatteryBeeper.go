package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getBatteryStatus() (string, error) {
	status, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/status")
	if err != nil {
		return "", err
	}
	return strings.Trim(string(status), "\n\t "), nil
}

func getBatteryCapacity() (uint8, error) {
	capacity, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		return uint8(0), err
	}

	uintegerCapacity, err := strconv.ParseUint(
		strings.Trim(string(capacity), "\n\t "),
		10, 8,
	)
	if err != nil {
		return uint8(0), err
	}

	return uint8(uintegerCapacity), nil
}

func main() {
	var (
		cmdPlaySound *exec.Cmd
	)
	musicPlayerPathToExecutable, err := exec.LookPath("mpg123")
	if err != nil {
		fmt.Println("ERROR: mpg123 Not found")
		return
	}
	pathToAudio := "/home/fern/.config/sway/modules/critical_battery_beeper/critical_battery.mp3"

	for {
		batteryStatus, _ := getBatteryStatus()
		if batteryStatus == "Discharging" {
			batteryCapacity, _ := getBatteryCapacity()
			if batteryCapacity <= uint8(10) {
				cmdPlaySound = exec.Command(musicPlayerPathToExecutable, "-q", pathToAudio)
				cmdPlaySound.Run()
			}
			time.Sleep(time.Minute)
		}
	}

}
