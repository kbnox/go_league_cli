package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fileName := flag.String("file", "", "The file containing the scores for each game.")
	flag.Parse()
	if *fileName == "" {
		log.Fatal("File not provided.")
	}
	fmt.Println("Hello Go")
}
