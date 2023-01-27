package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	log.Println(os.Args[0], os.Args[1], os.Args[2])

	read, error := os.Open(os.Args[1])
	if error != nil {
		log.Fatal("Error raised when opening:", error)
	}
	defer read.Close()

	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
	lst := [][]byte{}

	for {
		buffer := make([]byte, 100)
		content, error := read.Read(buffer)

		if error != nil {
			if error != io.EOF {
				log.Fatal(error)
			}
		}

		if error == io.EOF {
			break
		}
		lst = append(lst, buffer[:content])
	}

	sort.Slice(lst, func(i, j int) bool {
		return bytes.Compare(lst[i][:10], lst[j][:10]) < 0
	})

	write, error := os.Create(os.Args[2])
	if error != nil {
		log.Fatal("Error raised when creating:", error)
	}
	defer write.Close()

	writer := bufio.NewWriter(write)
	for _, j := range lst {
		_, err := writer.Write(j)
		if err != nil {
			log.Fatal("Error raised when writing:", err)
		}
	}

	writer.Flush()
}
