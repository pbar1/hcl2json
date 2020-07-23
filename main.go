package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

var version string

func main() {
	outputFormat := "json"
	filename := "-"
	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--help" || arg == "-help" {
			fmt.Println(`hcl2json

Converts Hashicorp Configuration Langauge (HCL) to JavaScript Object Notation (JSON).
Can also output YAML and TOML. If multiple output format command line flags and/or
filename arguments are given, the rightmost wins. If no filename or - is given, reads
from stdin.'

Usage:
  hcl2json [FLAGS] [FILENAME]

Examples:
  Concatenate all Terraform files in a directory convert the result to JSON via stdin
  > cat *.tf | hcl2json

  Convert single HCL file to YAML
  > hcl2json -y example.hcl

Flags:
  -h, --help      help for hcl2json
  -v, --version   print program version
  -j, --json      output JSON (default)
  -y, --yaml      output YAML
  -t, --toml      output TOML`)
			os.Exit(0)
		}
		if arg == "-v" || arg == "--version" || arg == "-version" {
			fmt.Println(version)
			os.Exit(0)
		}
		if arg == "-t" || arg == "--toml" || arg == "-toml" {
			outputFormat = "toml"
		} else if arg == "-y" || arg == "--yaml" || arg == "-yaml" {
			outputFormat = "yaml"
		} else if arg == "-j" || arg == "--json" || arg == "-json" {
			outputFormat = "json"
		} else {
			filename = arg
		}
	}

	var in, out []byte
	var err error

	if filename == "-" {
		in, err = ioutil.ReadAll(os.Stdin)
	} else {
		in, err = ioutil.ReadFile(filename)
	}
	check(err, "unable to read input from "+filename)

	var v interface{}
	err = hcl.Unmarshal(in, &v)
	check(err, "unable to unmarshal hcl input from "+filename)

	if outputFormat == "json" {
		out, err = json.Marshal(v)
	} else if outputFormat == "yaml" {
		out, err = yaml.Marshal(v)
	} else if outputFormat == "toml" {
		out, err = toml.Marshal(v)
	}
	check(err, "unable to marshal "+outputFormat+" output")

	fmt.Println(string(out))
}

func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}
