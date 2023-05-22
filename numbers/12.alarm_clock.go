package numbers

import (
	"fmt"
	"strings"
	"time"
)

// Alarm Clock :
// A simple clock where it plays a sound after X number of minutes/seconds or at a particular time.
func AlarmClock() {
	var (
		clockTypes     = []string{"\033[33mAlarm(a)", "Timer(t)\033[0m"}
		clockType      string
		clockSound     string
		timerTypes     = []string{"By Hours(h)", "By Minutes(m)", "By Seconds(s)"}
		timerType      string
		timerInputTime int
		alarmHour      int
		alarmMinute    int
	)

	fmt.Print("Which type of clock do you want to use?\n")
	fmt.Printf("Clock types: %v\n", clockTypes)
	fmt.Scanln(&clockType)
	clockType = strings.ToLower(clockType)

	switch clockType {
	case "alarm", "a":
		fmt.Print("Enter alarm's time(hour + minute)\n(e.g. 20 30 = 20:30 )\n")
		fmt.Scan(&alarmHour, &alarmMinute)

		fmt.Print("Enter alarm's sound: ")
		fmt.Scanln(&clockSound)

		alarm := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), alarmHour, alarmMinute, 0, 0, time.Local)
		duration := time.Until(alarm)

		time.Sleep(duration)

		fmt.Println(clockSound)

	case "timer", "t":
		fmt.Println("Choose your timer's type: ", timerTypes)
		fmt.Scanln(&timerType)
		timerType = strings.ToLower(timerType)

		switch timerType {
		case "by hours", "h":
			fmt.Print("Enter timer's time: ")
			fmt.Scanln(&timerInputTime)

			fmt.Print("Enter timer's sound: ")
			fmt.Scanln(&clockSound)

			fmt.Printf("Timer set for %d hours\n", timerInputTime)
			time.Sleep(time.Duration(timerInputTime) * time.Hour)

			fmt.Println(clockSound)

		case "by minutes", "m":
			fmt.Print("Enter timer's time: ")
			fmt.Scanln(&timerInputTime)

			fmt.Print("Enter timer's sound: ")
			fmt.Scanln(&clockSound)

			fmt.Printf("Timer set for %d minutes\n", timerInputTime)
			time.Sleep(time.Duration(timerInputTime) * time.Minute)

			fmt.Println(clockSound)

		case "by seconds", "s":
			fmt.Print("Enter timer's time: ")
			fmt.Scanln(&timerInputTime)

			fmt.Print("Enter timer's sound: ")
			fmt.Scanln(&clockSound)

			fmt.Printf("Timer set for %d seconds\n", timerInputTime)
			time.Sleep(time.Duration(timerInputTime) * time.Second)

			fmt.Println(clockSound)

		default:
			fmt.Println("Invalid timer type")
		}

	default:
		fmt.Println("Please choose a valid type of clock")
		AlarmClock()
	}
}
