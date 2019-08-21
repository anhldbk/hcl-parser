package main

import (
	"encoding/json"

	"github.com/gopherjs/gopherjs/js"
	"github.com/hashicorp/hcl"
)

func main() {
	js.Module.Get("exports").Set("parse", parse)
}

// Parse a HCL string into a JSON object
func parse(input string) (output *js.Object, err error) {
	var ast interface{}
	err = hcl.Unmarshal([]byte(input), &ast)
	if err != nil {
		return
	}

	data, err := json.MarshalIndent(ast, "", "    ")
	if err != nil {
		return
	}
	output = js.Global.Get("JSON").Call("parse", string(data))
	return
}
