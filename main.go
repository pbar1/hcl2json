package main

import (
	"encoding/json"
	"encoding/xml"
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
	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" || arg == "-help" {
			fmt.Println(`hcl2json

Converts Hashicorp Configuration Langauge (HCL) to JavaScript Object Notation (JSON).
Expects input from stdin. Can also output YAML, TOML, and XML.

Usage:
  cat *.tf | hcl2json

Flags:
  -h, --help      help for hcl2json
  -v, --version   print program version
  -j, --json      output JSON (default)
  -y, --yaml      output YAML
  -t, --toml      output TOML
  -x, --xml       output XML`)
			os.Exit(0)
		}
		if arg == "-v" || arg == "--version" || arg == "-version" {
			fmt.Println(version)
			os.Exit(1)
		}
		if arg == "-x" || arg == "--xml" || arg == "-xml" {
			outputFormat = "xml"
		}
		if arg == "-t" || arg == "--toml" || arg == "-toml" {
			outputFormat = "toml"
		}
		if arg == "-y" || arg == "--yaml" || arg == "-yaml" {
			outputFormat = "yaml"
		}
		if arg == "-j" || arg == "--json" || arg == "-json" {
			outputFormat = "json"
		}
	}

	in, err := ioutil.ReadAll(os.Stdin)
	check(err)

	var v interface{}
	err = hcl.Unmarshal(in, &v)
	check(err)

	var out []byte
	if outputFormat == "json" {
		out, err = json.Marshal(v)
	} else if outputFormat == "yaml" {
		out, err = yaml.Marshal(v)
	} else if outputFormat == "toml" {
		out, err = toml.Marshal(v)
	} else if outputFormat == "xml" {
		out, err = xml.Marshal(v)
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
