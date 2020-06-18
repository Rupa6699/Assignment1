package main

import (
	//implements formatted I/O functions
	"fmt"
	//implements some I/O utility functions
	"io/ioutil"
	//provides HTTP client and server implementations
	"net/http"
)

func main() {
	//calling getresponse function
	getresponse()
	//calling getcontent function
	getcontent()
	//calling getdata function
	getdata()
}

//implementing getresponse()
func getresponse(){

	//json data
	url := "http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b"

	res, err := http.Get(url)

	if err != nil {
		//handle error
		fmt.Printf("failed %s\n", err)
	} else {
		//ReadAll reads from src until an error or EOF and returns the data it read
		data, _ := ioutil.ReadAll(res.Body)
		//Printing the content stored in json
		fmt.Println(string(data))
	}

}
//implementing getcontent()
func getcontent() {

	//json data
	url := "http://www.mocky.io/v2/5ecfd630320000f1aee3d64d"

	res, err := http.Get(url)

	if err != nil {
		//handle error
		fmt.Printf("failed %s\n", err)
	} else {
		//ReadAll reads from src until an error or EOF and returns the data it read
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}
}
//implementing getdata()
func getdata(){
	
	//json data
	url := "http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e"

	res, err := http.Get(url)

	if err != nil {
		//handle error
		fmt.Printf("failed %s\n", err)
	} else {
		//ReadAll reads from src until an error or EOF and returns the data it read
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(data))
	}

}
	

