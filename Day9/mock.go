package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3

	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, err := fmt.Fprintln(out, i)
		if err != nil {
			return
		}
	}
	sleeper.Sleep()
	_, err := fmt.Fprint(out, finalWord)
	if err != nil {
		return
	}
}

func main() {
	sleeper := &ConfigurableSleeper{5 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
