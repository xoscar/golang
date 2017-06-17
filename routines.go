package main

import (
	"fmt"
	"time"
	"strings"
	"io/ioutil"
	"bufio"
	"os"
)

func splitWord(word string) {
	letters := strings.Split(word, "")

	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}

func routines() {
	go splitWord("Oscar")
	fmt.Println("after split")
	var wait string
	fmt.Scanln(&wait)
}

func channels() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	msg := <- channel

	fmt.Println(msg)
}

func iofunction() {
	data, err := ioutil.ReadFile("./file.json")

	if err != nil {
		fmt.Println("Error happened")
	} else {
		fmt.Println(string(data))
	}
}

func ioLines() bool {
	file, err := os.Open("./files.json")
	defer func () {
		file.Close()
		recover()
	}()

	if err != nil {
		fmt.Println("Error happened")
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		i++
		fmt.Println(scanner.Text(), i)
	}

	return true
}

func main() {
	ioLines()
}