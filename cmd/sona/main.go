package main

import (
	"fmt"
	"os"

	"github.com/valentino7504/sona/internal/audio"
)

func main() {
	cmdArgs := os.Args
	fileName := cmdArgs[1]
	player, err := audio.NewAudioPlayer(fileName, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	player.Play()
}
