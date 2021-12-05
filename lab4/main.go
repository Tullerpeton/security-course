package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/security-course/lab4/rsa"
)

const (
	Path = "./lab4/test_data/"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v input_file bits\n", os.Args[0])
		return
	}

	inputFile := os.Args[1]
	bits, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
		return
	}

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	path := strings.Split(inputFile, "/")
	fileName := path[len(path)-1]

	log.Printf("INPUT [%s] %d bytes\n", inputFile, len(data))

	rsa := rsa.NewRSA(bits)

	start := time.Now()
	encData := rsa.Encode(data)
	duration := time.Since(start)
	fmt.Println("Encode time:", duration)
	log.Println("Encode time:", duration)

	encFile := Path + "enc_" + fileName
	ioutil.WriteFile(encFile, encData, 0666)

	start = time.Now()
	decData := rsa.Decode(encData)
	duration = time.Since(start)
	fmt.Println("Decode time:", duration)
	log.Println("Decode time:", duration)

	decFile := Path + "dec_" + fileName
	ioutil.WriteFile(decFile, decData, 0666)

	fmt.Println("Successfully done!")
}
