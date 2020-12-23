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

	fmt.Println("{")
	flatten("", "", x)
	fmt.Println("\n}")
}

func flatten(comma string, hdr string, d interface{}) {
	switch d.(type) {
	case float64:
		fmt.Printf("%s\n\t%q: %f", comma, hdr, d.(float64))
	case map[string]interface{}:
		for name, interf := range d.(map[string]interface{}) {
			newhdr := name
			if hdr != "" {
				newhdr = fmt.Sprintf("%s.%s", hdr, name)
			}
			flatten(comma, newhdr, interf)
			comma = ","
		}
	}
}
