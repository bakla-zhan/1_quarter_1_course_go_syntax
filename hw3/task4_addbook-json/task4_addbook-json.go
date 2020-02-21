package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// за основу взят телефонный справочник из методички
	addressBook := make(map[string][]int)

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)

	fmt.Println(addressBook)

	for name, numbers := range addressBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	} // за основу взят телефонный справочник из методички

	// внесение в справочник дополниельных данных
	addressBook["Candy"] = []int{89260001122}
	fmt.Println(addressBook)

	// создание последовательности байт json
	aBookJSON, err := json.Marshal(addressBook)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(aBookJSON)
	//fmt.Println(string(aBookJSON))

	// запись в файл
	errMakeFile := ioutil.WriteFile("addressBook.json", aBookJSON, 0644)
	if errMakeFile != nil {
		log.Fatal(errMakeFile)
	}
}
