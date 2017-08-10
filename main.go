package main

import (
	"fmt"
	"github.com/sssgun/mp3utils/mp3"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		log.Printf("Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Couldn't open %s.\n", filename)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Printf("Couldn't open %s.\n", filename)
		os.Exit(1)
	}

	data := make([]byte, stat.Size())
	log.Printf("Buffer of %d\n\n", len(data))

	file.Read(data)

	seek := 0
	frame, err := mp3.ReadFrame(data, 0)

	var frameLength int
	frameLength, err = frame.Header.GetFrameLength()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	seek += frameLength

	for i := 0; i < 2; i++ {
		frame, err := mp3.ReadFrame(data, seek)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		frameLength, err = frame.Header.GetFrameLength()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		seek += frameLength

		log.Println(frame)
	}
}
