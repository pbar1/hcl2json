# `hcl2json`

### Usage
```
hcl2json 1.0.0

Converts Hashicorp Configuration Langauge (HCL) to JavaScript Object Notation (JSON).
Expects input from stdin.

Usage:
  cat *.tf | hcl2json

Flags:
  -h, --help   help for hcl2json
  -y, --yaml   output YAML
```

### Example

```
♪ ~ cat <<EOF | hcl2json | jq
terraform {          
  required_version = "~> 0.12.0" 
}

resource "null_resource" "test" {}
EOF

# outputs the following...
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
