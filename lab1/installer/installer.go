package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
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

	file, err := os.OpenFile("./lab1/key.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha256.Sum256([]byte(serialNumber))
	file.Write([]byte(hex.EncodeToString(hash[:])))
}
