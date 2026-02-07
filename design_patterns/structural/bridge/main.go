package main

import "fmt"

// Bridge Pattern
//
// 5-Year-Old Explanation:
// Imagine you have a TV Remote and a TV.
// The Remote (the clicker) is one thing. The TV (the screen) is another thing.
// You can have different kinds of Remotes (Basic, Fancy) and different kinds of TVs (Sony, Samsung).
// The Bridge pattern separates the "Remote" from the "TV" so you can mix and match them!
// A Fancy Remote can work with a Sony TV OR a Samsung TV.

// -- Implementation (The Device/TV) --
type Device interface {
	Run()
	IsEnabled() bool
	Enable()
	Disable()
	SetVolume(percent int)
}

// Tv is a specific device
type Tv struct {
	on     bool
	volume int
}

func (t *Tv) Run() {
	fmt.Println("TV is running...")
}

func (t *Tv) IsEnabled() bool {
	return t.on
}

func (t *Tv) Enable() {
	fmt.Println("TV: Turned ON")
	t.on = true
}

func (t *Tv) Disable() {
	fmt.Println("TV: Turned OFF")
	t.on = false
}

func (t *Tv) SetVolume(percent int) {
	t.volume = percent
	fmt.Printf("TV: Volume set to %d\n", t.volume)
}

// Radio is another device
type Radio struct {
	on     bool
	volume int
}

func (r *Radio) Run() {
	fmt.Println("Radio is playing music...")
}

func (r *Radio) IsEnabled() bool {
	return r.on
}

func (r *Radio) Enable() {
	fmt.Println("Radio: Turned ON")
	r.on = true
}

func (r *Radio) Disable() {
	fmt.Println("Radio: Turned OFF")
	r.on = false
}

func (r *Radio) SetVolume(percent int) {
	r.volume = percent
	fmt.Printf("Radio: Volume set to %d\n", r.volume)
}

// -- Abstraction (The Remote) --
type RemoteControl interface {
	TogglePower()
	VolumeDown()
	VolumeUp()
}

// BasicRemote is a simple remote
type BasicRemote struct {
	device Device
}

func (r *BasicRemote) TogglePower() {
	fmt.Println("Remote: Power button pressed.")
	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}
}

func (r *BasicRemote) VolumeDown() {
	fmt.Println("Remote: Volume Down pressed.")
	r.device.SetVolume(0)
}

func (r *BasicRemote) VolumeUp() {
	fmt.Println("Remote: Volume Up pressed.")
	r.device.SetVolume(100)
}

// AdvancedRemote can do more (mute)
type AdvancedRemote struct {
	BasicRemote // Inherits from BasicRemote
}

func (r *AdvancedRemote) Mute() {
	fmt.Println("Advanced Remote: Mute pressed.")
	r.device.SetVolume(0)
}

func main() {
	fmt.Println("--- Bridge Pattern: Remote and Devices ---")

	testDevice := func(device Device) {
		fmt.Println("Tests with Basic Remote.")
		basicRemote := BasicRemote{device: device}
		basicRemote.TogglePower()
		basicRemote.VolumeUp()
		basicRemote.TogglePower()

		fmt.Println("\nTests with Advanced Remote.")
		advancedRemote := AdvancedRemote{BasicRemote{device: device}}
		advancedRemote.TogglePower()
		advancedRemote.Mute()
		advancedRemote.TogglePower()
	}

	fmt.Println("\n--> Connecting to TV")
	tv := &Tv{}
	testDevice(tv)

	fmt.Println("\n--> Connecting to Radio")
	radio := &Radio{}
	testDevice(radio)
}
