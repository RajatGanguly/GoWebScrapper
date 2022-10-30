package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://amazon.com"

func main() {
	fmt.Println("WEB REQUEST")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("PANIC Happens")
		panic(err)
	}
	fmt.Printf("%T\n", response)

	defer response.Body.Close()
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("PANIC Happens")
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}

