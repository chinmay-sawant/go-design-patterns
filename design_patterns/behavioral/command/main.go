package main

import "fmt"

// Command Pattern
//
// 5-Year-Old Explanation:
// Imagine a TV Remote Control.
// Each button is a "Command".
// One button knows how to "Turn On TV".
// Another button knows how to "Mute Volume".
// You (the Invoker) just press the button. You don't need to know how the TV works inside.
// You can also "Undo" if you pressed the wrong button!

// Command interface
type Command interface {
	Execute()
}

// Receiver (The TV)
type TV struct {
	IsOn bool
}

func (t *TV) On() {
	t.IsOn = true
	fmt.Println("TV is ON")
}

func (t *TV) Off() {
	t.IsOn = false
	fmt.Println("TV is OFF")
}

// -- Concrete Commands --

type TurnOnCommand struct {
	tv *TV
}

func (c *TurnOnCommand) Execute() {
	c.tv.On()
}

type TurnOffCommand struct {
	tv *TV
}

func (c *TurnOffCommand) Execute() {
	c.tv.Off()
}

// Invoker (The Remote Button)
type RemoteButton struct {
	command Command
}

func (b *RemoteButton) Press() {
	b.command.Execute()
}

func main() {
	fmt.Println("--- Command Pattern: TV Remote ---")

	myTV := &TV{}

	// Setup commands
	onCommand := &TurnOnCommand{tv: myTV}
	offCommand := &TurnOffCommand{tv: myTV}

	// Create buttons
	onButton := &RemoteButton{command: onCommand}
	offButton := &RemoteButton{command: offCommand}

	// Press buttons
	fmt.Println("User presses ON button:")
	onButton.Press()

	fmt.Println("User presses OFF button:")
	offButton.Press()
}
