package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
		m := d.(map[string]interface{})
		// sort keys to make output repeatable
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		if len(keys) == 0 {
			fmt.Printf("%s\n\t%q: {}", comma, hdr)
			return
		}
		sort.Sort(sort.StringSlice(keys))
		for _, name := range keys {
			newhdr := name
			if hdr != "" {
				newhdr = fmt.Sprintf("%s.%s", hdr, name)
			}
			flatten(comma, newhdr, m[name])
			comma = ","
		}
	}
}
