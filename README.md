# `hcl2json`

![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/pbar1/hcl2json/latest?label=docker&style=flat-square)

### Get

Either download the appropriate binary from the _Releases_ page, or run the following:

```shell script
docker run --rm pbar1/hcl2json -- --help
```

### Usage
```
hcl2json

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
  -t, --toml      output TOML
```

### Example

```shell script
♪ ~ cat <<EOF | hcl2json | jq
terraform {          
  required_version = "~> 0.12.0" 
}

resource "null_resource" "test" {}
EOF
```

Outputs the following...
```json
{
  "resource": [
    {
      "null_resource": [
        {
          "test": [
            {}
          ]
        }
      ]
    }
  ],
  "terraform": [
    {
      "required_version": "~> 0.12.0"
    }
  ]
}
```
