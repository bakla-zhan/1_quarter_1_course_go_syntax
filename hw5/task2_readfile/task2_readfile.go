package main

import (
	"fmt"
	"os"
)

// исходная программа чтения файла из методички
func main() {
	file, err := os.Open("../task4_copy_file/chat1.txt")
	if err != nil {
		return
	}
	defer file.Close()

	/* Ну например можно было бы производить чтение файла построчно,
	т.к. если размер файла большой, например 2ГБ, то с использованием file.Stat()
	для операции копирования должно будет выделено все 2ГБ памяти.

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}*/

	data := make([]byte, 128)

	for {
		n, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Println(string(data[:n]))
	}
}
