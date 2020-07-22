package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"gopkg.in/yaml.v2"
)

func main() {
	outputYAML := false
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" || arg == "-help" {
			fmt.Println(`hcl2json 1.0.0

Converts Hashicorp Configuration Langauge (HCL) to JavaScript Object Notation (JSON).
Expects input from stdin.

Usage:
  cat *.tf | hcl2json

Flags:
  -h, --help   help for hcl2json
  -y, --yaml   output YAML`)
			os.Exit(0)
		}
		if arg == "-y" || arg == "--yaml" || arg == "-yaml" {
			outputYAML = true

		}
	}

	in, err := ioutil.ReadAll(os.Stdin)
	check(err)

	var v interface{}
	err = hcl.Unmarshal(in, &v)
	check(err)

	var out []byte
	if outputYAML {
		out, err = yaml.Marshal(v)
	} else {
		out, err = json.Marshal(v)
	}
	check(err)

	fmt.Println(string(out))
}

func check(err error) {
	if err != nil {
		if _, err := fmt.Fprintln(os.Stderr); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}
