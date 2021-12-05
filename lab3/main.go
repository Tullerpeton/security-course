package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/security-course/lab3/des"
)

const (
	Key  = "./lab3/config/key.txt"
	Path = "./lab3/test_data/"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v input_file\n", os.Args[0])
		return
	}

	inputFile := os.Args[1]

	key, err := ioutil.ReadFile(Key)
	if err != nil {
		log.Fatal(err)
	}

	desEcb := des.NewDES(string(key))
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	pwd, _ := os.Getwd()

	path := strings.Split(inputFile, "/")
	fileName := path[len(path)-1]
	encData := desEcb.Encode(data)
	encFile := Path + "enc_" + fileName
	ioutil.WriteFile(pwd+"/"+encFile, encData, 0666)

	decData := desEcb.Decode(encData)
	decFile := Path + "dec_" + fileName
	ioutil.WriteFile(pwd+"/"+decFile, decData, 0666)

	fmt.Println("Successfully done!")
}
