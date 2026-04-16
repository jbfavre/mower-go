package main

import (
	"bufio"
	"fmt"
	"log"
	"mower-go/src/mower"
	"os"
)

func main() {
	var area = mower.Rectangle{5, 5}
	// var mower = mower.Mower{Area: area, X: 0, Y: 0, Orientation : mower.N}
	var mower = mower.NewMower(0, 0, area, mower.N)
	fmt.Println("OK")
	mower.Print()
	mower.Advance()
	mower.Print()
	mower.TurnRight()
	mower.Print()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Advance()
	mower.Print()

	inputFile, err := os.Open("file.txt")
	if err != nil {
		log.Fatal("error opening input file", err)
	}

	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
