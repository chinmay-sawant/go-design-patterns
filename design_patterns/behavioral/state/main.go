package main

import "fmt"

// State Pattern
//
// 5-Year-Old Explanation:
// Imagine a Traffic Light.
// It can be in different "States": GREEN, YELLOW, or RED.
// When it is GREEN, cars GO.
// When it is RED, cars STOP.
// The light behaves DIFFERENTLY depending on what color (State) it is right now.

type State interface {
	Next(light *TrafficLight)
}

type TrafficLight struct {
	state State
}

func (t *TrafficLight) SetState(s State) {
	t.state = s
}
func (t *TrafficLight) Change() {
	t.state.Next(t)
}

// -- Concrete States --

// RedState
type RedState struct{}

func (r *RedState) Next(t *TrafficLight) {
	fmt.Println("RED LIGHT: Stop! ... Changing to Green.")
	t.SetState(&GreenState{})
}

// GreenState
type GreenState struct{}

func (g *GreenState) Next(t *TrafficLight) {
	fmt.Println("GREEN LIGHT: Go! ... Changing to Yellow.")
	t.SetState(&YellowState{})
}

// YellowState
type YellowState struct{}

func (y *YellowState) Next(t *TrafficLight) {
	fmt.Println("YELLOW LIGHT: Slow down! ... Changing to Red.")
	t.SetState(&RedState{})
}

func main() {
	fmt.Println("--- State Pattern: Traffic Light ---")

	// Start with Green
	light := &TrafficLight{state: &GreenState{}}

	// Cycle through steps
	for i := 0; i < 6; i++ {
		light.Change()
	}
}
