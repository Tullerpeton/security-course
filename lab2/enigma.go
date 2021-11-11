package main

import (
	"strconv"
	"strings"
)

type Enigma struct {
	rotors    [RotorCount]Rotor
	reflector Reflector
}

func NewEnigma(settings string) *Enigma {
	e := new(Enigma)

	var rotorSettings [RotorCount]string
	var reflectorSettings string

	s := strings.Split(settings, "\n")
	for i := 0; i < RotorCount; i++ {
		rotorSettings[i] = s[i]
	}
	reflectorSettings = s[RotorCount]

	for i := 0; i < RotorCount; i++ {
		e.rotors[i] = makeRotor(rotorSettings[i])
	}

	e.reflector = makeReflector(reflectorSettings)

	return e
}

func (e *Enigma) Code(text []byte) []byte {
	var encText []byte
	for _, symbol := range text {
		encSymbol := symbol

		// Прямой проход
		for i := 0; i < RotorCount; i++ {
			encSymbol = byte(e.rotors[i].getVal(int(encSymbol)))
		}

		encSymbol = byte(e.reflector.getVal(int(encSymbol)))

		// Обратный проход
		for i := RotorCount - 1; i >= 0; i-- {
			encSymbol = byte(e.rotors[i].getKey(int(encSymbol)))
		}

		encText = append(encText, encSymbol)
		e.rotate()
	}

	return encText
}

func (e *Enigma) rotate() {
	for i := 0; i < RotorCount; i++ {
		// Если предыдущий ротор сделал полный оборот, то следующий выполняется сдвиг
		if !e.rotors[i].inc() {
			break
		}
	}
}

func GenerateEnigma() (*Enigma, string) {
	// Генерация Энигмы с рандомными настройками
	e := new(Enigma)
	for i := 0; i < RotorCount; i++ {
		e.rotors[i] = generateRotor()
	}
	e.reflector = generateReflector()

	// Форматирование настроек для сохранения в файл
	settings := ""
	for i := 0; i < RotorCount; i++ {
		ring := e.rotors[i].getRing()
		for _, v := range ring {
			settings += strconv.Itoa(v) + " "
		}
		settings = settings[:len(settings)-1] + "\n"
	}

	mapping := e.reflector.getMapping()
	for i := 0; i < Length; i++ {
		settings += strconv.Itoa(i) + ":" + strconv.Itoa(mapping[i]) + " "
	}
	settings = settings[:len(settings)-1] + "\n"

	return e, settings
}
