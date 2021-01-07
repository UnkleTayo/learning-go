package main

import "fmt"

type Guitarist interface {

	// Base guitar and accoustic guitarist will implement the Guitarist inteface which returns a function
	playGuitar()
}

// This should make you stop screaming 
type BaseGuitarist struct {
	Name string
}

// This should make you stop screaming 
type AcousticGuitarist struct {
	Name string
}

// This should make you stop screaming 
func (b BaseGuitarist) playGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

// This should make you stop screaming 
func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func guitarist() {
	var player BaseGuitarist
	player.Name = "Marcelo"
	player.playGuitar()

	var player2 AcousticGuitarist
    player2.Name = "Ringo"
    player2.PlayGuitar()

}