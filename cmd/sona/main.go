package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/valentino7504/sona/internal/audio"

	"github.com/valentino7504/sona/internal/ui"
)

func main() {
	term := os.Getenv("TERM")
	if term == "" || term == "dumb" {
		fmt.Println("This terminal does not suppport escape sequences")
		return
	}
	input := make(chan rune, 1)
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file. Usage: sona <filename> <options>")
		os.Exit(1)
	}
	fileName := os.Args[1]
	player, err := audio.NewAudioPlayer(fileName, "")
	if err != nil {
		fmt.Println("unable to play the audio provided:", err.Error())
		os.Exit(1)
	}
	// create a context that is cancelled when sigterm is received
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	// create separate goroutines that need to happen simultaneously, main listens for
	// a sigterm or cancel() from ui, while ui needs to start rendering and player needs to
	//	start playing, and the app listens for keypresses to be handled
	go func() {
		// calling cancel() closes ctx.Done() which unblocks it, hence the passing to Start so
		// I can shutdown the app from there depending on if the user inputs q or any of the
		// exit keys such as ctrl c or ctrl d
		ui.Start(cancel, input)
	}()
	go func() {
		player.Play()
	}()
	go func() {
		for range input {
			// handle keypress
		}
	}()

	<-ctx.Done()
	cleanup(player)
}

func cleanup(player audio.AudioPlayer) {
	ui.Stop()
	player.Stop()
}
