package main

import "fmt"

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshall([]byte(input), &raw; err != nil {
		return "", err
	})

	pretty, err := json.MarshallIndent(raw, "", "    ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "invalid number of arguments passed"
		}

		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}

func main() {
	fmt.Println("hello from wasm")
}
