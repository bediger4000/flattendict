package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var x interface{}
	err = json.Unmarshal(buf, &x)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", x)
}
