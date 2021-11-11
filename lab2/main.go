package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	Length     = 256
	RotorCount = 3
	Settings   = "settings.txt"
	HelpMessage = "enigma input_file output_file [-n]"
)

func main() {
	generateSettings := flag.Bool("n", false, `[-n] - generate and set new random settings`)
	flag.Parse()

	if len(os.Args) < 3 {
		fmt.Println(HelpMessage)
		return
	}

	var enigma *Enigma
	var newSettings string
	if *generateSettings {
		enigma, newSettings = GenerateEnigma()
		if err := ioutil.WriteFile(Settings, []byte(newSettings), 0666); err != nil {
			log.Fatal(err)
		}
	} else {
		settings, err := ioutil.ReadFile(Settings)
		if err != nil {
			log.Fatal(err)
		}
		enigma = NewEnigma(string(settings))
	}

	in := os.Args[1]
	out := os.Args[2]

	data, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	encText := enigma.Code(data)
	if err := ioutil.WriteFile(out, encText, 0666); err != nil {
		log.Fatal(err)
	}
}
