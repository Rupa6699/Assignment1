package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getresponse()
	getcontent()
	getdata()
}

func getresponse(){

	url := "http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b"

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("failed %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}
func getcontent() {

	url := "http://www.mocky.io/v2/5ecfd630320000f1aee3d64d"

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("failed %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}
}
func getdata(){

	url := "http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e"

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("failed %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}
	

