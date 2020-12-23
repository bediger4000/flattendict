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

	flatten("", x)
}

func flatten(hdr string, d interface{}) {
	// fmt.Printf("%q: %T\n", hdr, d)
	switch d.(type) {
	case float64:
		fmt.Printf("\t%q: %f,\n", hdr, d.(float64))
	case map[string]interface{}:
		if hdr != "" {
			fmt.Printf("\t%q: ", hdr)
		}
		fmt.Println("{")
		for name, interf := range d.(map[string]interface{}) {
			newhdr := name
			if hdr != "" {
				newhdr = fmt.Sprintf("%s.%s", hdr, name)
			}
			flatten(newhdr, interf)
		}
		fmt.Println("}")
	}
}
