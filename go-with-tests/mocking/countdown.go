package mocking

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

func main() {
	// sleeper := DefaultSleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)

	fmt.Printf("start time %v\n", time.Now())
	cSleeper := ConfigurableSleeper{3 * time.Second, time.Sleep}
	cSleeper.Sleep()
	fmt.Printf("end time %v\n", time.Now())
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

// INITIAL IMPLEMENTATION
//	func Countdown(out io.Writer) {
//		for i := countdownStart; i > 0; i-- {
//			fmt.Fprintln(out, i)
//			time.Sleep(1 * time.Second)
//		}
//		fmt.Fprint(out, finalWord)
//	}

// type DefaultSleeper struct{}

// func (ds DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

// IMPLEMENTATION REFACTOR
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	// time.Sleep(cs.duration)
	cs.sleep(cs.duration)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
