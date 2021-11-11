package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Rotor struct {
	pos  int
	ring [Length]int
}

func makeRotor(settings string) Rotor {
	var r Rotor
	s := strings.Split(settings, " ")
	for k, v := range s {
		r.ring[k], _ = strconv.Atoi(v)
	}

	return r
}

func generateRotor() Rotor {
	var r Rotor
	for i := 0; i < Length; i++ {
		r.ring[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(Length, func(i, j int) {
		r.ring[i], r.ring[j] = r.ring[j], r.ring[i]
	})

	return r
}

func (r *Rotor) inc() bool {
	r.pos++
	defer func() {
		r.pos %= Length
	}()
	return r.pos == Length
}

func (r *Rotor) getVal(in int) int {
	return r.ring[(in+r.pos)%Length]
}

func (r *Rotor) getKey(val int) int {
	for i := 0; i < Length; i++ {
		if r.ring[i] == val {
			key := (i - r.pos) % Length
			return key
		}
	}
	return -1
}

func (r Rotor) getRing() [Length]int {
	return r.ring
}
