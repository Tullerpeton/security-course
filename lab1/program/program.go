package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main()  {
	data, err := os.ReadFile("./lab1/key.txt")
	if err != nil {
		log.Fatal(err)
	}

	var serialNumber string
	out, err := exec.Command("/usr/sbin/ioreg", "-l").Output()
	for _, l := range strings.Split(string(out), "\n") {
		if strings.Contains(l, "IOPlatformSerialNumber") {
			s := strings.Split(l, " ")
			serialNumber = s[len(s)-1]
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.Sum256([]byte(serialNumber))
	if string(data) != hex.EncodeToString(hash[:]) {
		fmt.Println("\nYou need to purchase a license")
	} else {
		fmt.Println("You are welcome!")
	}
}
