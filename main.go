package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	var b []byte

	if len(os.Args) == 0 {
		b = stdinReader()
	} else {
		b = fileReader(os.Args[1])
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

func stdinReader() []byte {
	// scan能够按行扫描输入
	sc := bufio.NewScanner(os.Stdin)

	// var s []byte
	var s string
	for sc.Scan() {
		txt := sc.Text()
		s += txt + "\n"
	}

	return []byte(s)
}
