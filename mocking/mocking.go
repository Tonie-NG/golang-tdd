package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
  Sleep()
}

type SpySleeper struct {
  Calls int
}

func (s *SpySleeper) Sleep() {
  s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const startCount = 3
const finalWord = "Go!"

func Countdown(w io.Writer, s Sleeper) {
  for i := startCount; i > 0; i--{
    fmt.Fprintln(w, i)
    s.Sleep()
  }
  fmt.Fprint(w, finalWord)
}

func main() {
  sleeper := &DefaultSleeper{}
  Countdown(os.Stdout, sleeper)
}
