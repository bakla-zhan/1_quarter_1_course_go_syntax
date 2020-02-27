package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	f1 = flag.String("file1", "chat1.txt", "Your file origin")
	f2 = flag.String("file2", "chat2.txt", "Your file destination")
)

func main() {
	flag.Parse()

	file1, err := os.Open(*f1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file2, err := os.Create(*f2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file2.Close()
	defer file1.Close()

	// getting size of file
	stat, err := file1.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make([]byte, stat.Size())

	_, err = file1.Read(data)
	file2.Write(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file copied")
}
