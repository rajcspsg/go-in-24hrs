package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type RDD2 struct {
	Broken bool
}

func (r *RDD2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func simpleFunc(r Robot) {
	fmt.Println(r)
}

func main() {
	t := T850{
		Name: "The Terminator",
	}

	r := RDD2{
		Broken: true,
	}

	simpleFunc(&r)

	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}
