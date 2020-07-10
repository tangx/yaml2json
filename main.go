package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	var b []byte

	if len(os.Args) > 1 {
		b = fileReader(os.Args[1])
	} else {
		fmt.Printf("Usage: %s path/2/file.yml\n", os.Args[0])
		os.Exit(1)
	}

	yaml2json(b)
}

func yaml2json(y []byte) {
	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", j)
}

func fileReader(file string) (b []byte) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return
}
