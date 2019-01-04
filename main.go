package main

import (
	"fmt"
	"log"
	"net/http"
	"testts"
)

func main() {
	//base.BasePrint()
	testg.IrisWebTest()
}

func startWeb() {
	http.HandleFunc("/", sayHello)
	er := http.ListenAndServe(":9090", nil)
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")

}