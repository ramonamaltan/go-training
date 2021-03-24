package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println(err) // open no-file.txt: no such file or directory
		log.Println(err) // also get datetimestamp
		// log.Fatalln(err) // log message with exit status 1
		// log.Panicln(err) // log message plus panic(err)

		// panic(err) // stops normal execution of the current goroutine
		// with panic defered functions run and recover can be used
	}
}

// ideally choose between logging errors
