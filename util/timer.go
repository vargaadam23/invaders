package util

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	Start    float64
	Duration float64
}

func (timer *Timer) StartTimer(duration float64) {
	timer.Start = rl.GetTime()
	timer.Duration = duration
}

func (timer *Timer) TimerDone() bool {
	return rl.GetTime()-timer.Start >= timer.Duration
}

func (timer *Timer) TimerElapsed() float64 {
	return rl.GetTime() - timer.Start
}
