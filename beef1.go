package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
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

	dict := make(map[string]interface{})

	inflate(dict, x)

	buf, err = json.Marshal(dict)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(buf))
}

func inflate(dict map[string]interface{}, x interface{}) {

	switch x.(type) {
	case map[string]interface{}:
		d := x.(map[string]interface{})
		var keys []string
		for key, _ := range d {
			keys = append(keys, key)
		}
		sort.Sort(sort.StringSlice(keys))

		for _, key := range keys {
			subkeys := strings.Split(key, ".")
			if len(subkeys) == 1 {
				switch d[key].(type) {
				case float64:
					dict[key] = d[key].(float64)
				case map[string]interface{}:
					subdict := make(map[string]interface{})
					inflate(subdict, d[key].(map[string]interface{}))
					dict[key] = subdict
				}
				continue
			}

			sdict := dict
			for _, subkey := range subkeys[0 : len(subkeys)-1] {
				if _, ok := sdict[subkey]; !ok {
					sdict[subkey] = make(map[string]interface{})
				}
				sdict = sdict[subkey].(map[string]interface{})
			}
			sdict[subkeys[len(subkeys)-1]] = d[key]
		}
	}
}
