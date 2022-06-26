package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"log"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()
	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)
	//Read data 1=MB at a time
	bufferSize := make([]byte, 0, 1024)
	for {
		read, err := reader.Read(bufferSize[:cap(bufferSize)])
		bufferSize = bufferSize[:read]
		if read == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				log.Fatal(err)
				break
			}
		}

		content := string(bufferSize)

		//In order to scale, this could be put into a separate go routine, or add a set of worker threads to parse
		//the content using channels to store and pull from with infinite for loops and cancel contexts
		if strings.Contains(content, "error") {
			fmt.Println(content)
		}
	}
}
