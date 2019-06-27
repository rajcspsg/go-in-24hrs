package main

import "fmt"

type Alarm struct {
	Time  string
	Sound string
}

func newAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Kalxon",
	}
	return a
}

func main() {
	b := newAlarm("morning")
	fmt.Println(b)
}
