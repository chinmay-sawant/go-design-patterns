package main

import "fmt"

// Facade Pattern
//
// 5-Year-Old Explanation:
// Imagine you want to play a video game.
// You have to:
// 1. Turn on the TV.
// 2. Turn on the Sound System.
// 3. Turn on the Game Console.
// 4. Select the right input on TV.
// That's too many buttons!
// A "Facade" is like a magical single button that says "PLAY GAME".
// You press it, and it does ALL those things for you automatically.

// -- Complex Subsystem Parts --

type TV struct{}
func (t *TV) On() { fmt.Println("TV: Turning ON") }
func (t *TV) Off() { fmt.Println("TV: Turning OFF") }

type SoundSystem struct{}
func (s *SoundSystem) On() { fmt.Println("Sound: Turning ON") }
func (s *SoundSystem) SetVolume(vol int) { fmt.Printf("Sound: Volume set to %d\n", vol) }
func (s *SoundSystem) Off() { fmt.Println("Sound: Turning OFF") }

type GameConsole struct{}
func (g *GameConsole) On() { fmt.Println("Console: Turning ON") }
func (g *GameConsole) StartGame() { fmt.Println("Console: Starting the game...") }
func (g *GameConsole) Off() { fmt.Println("Console: Turning OFF") }

// -- The Facade --

type GameFacade struct {
	tv      *TV
	sound   *SoundSystem
	console *GameConsole
}

func NewGameFacade() *GameFacade {
	return &GameFacade{
		tv:      &TV{},
		sound:   &SoundSystem{},
		console: &GameConsole{},
	}
}

// PlayGame is the simple button
func (g *GameFacade) PlayGame() {
	fmt.Println("\n>>> Master Button: PLAY GAME <<<")
	g.tv.On()
	g.sound.On()
	g.sound.SetVolume(50)
	g.console.On()
	g.console.StartGame()
	fmt.Println(">>> Ready to play! <<<")
}

// StopGame is another simple button
func (g *GameFacade) StopGame() {
	fmt.Println("\n>>> Master Button: STOP GAME <<<")
	g.console.Off()
	g.sound.Off()
	g.tv.Off()
	fmt.Println(">>> Goodnight! <<<")
}

func main() {
	fmt.Println("--- Facade Pattern: The Master Button ---")

	facade := NewGameFacade()
	
	// Start everything with one call
	facade.PlayGame()

	// Stop everything with one call
	facade.StopGame()
}
