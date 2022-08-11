package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/k0kubun/pp"
)

func init() {
	pp.SetColorScheme(pp.ColorScheme{
		Bool:            pp.NoColor,
		Integer:         pp.NoColor,
		Float:           pp.NoColor,
		String:          pp.NoColor,
		StringQuotation: pp.NoColor,
		EscapedChar:     pp.NoColor,
		FieldName:       pp.NoColor,
		PointerAdress:   pp.NoColor,
		Nil:             pp.NoColor,
		Time:            pp.NoColor,
		StructName:      pp.NoColor,
		ObjectLength:    pp.NoColor,
	})
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("jsonToGo", js.FuncOf(jsonToGo))
	<-c
}

func jsonToGo(this js.Value, args []js.Value) any {
	inputJson := args[0].String()

	var mData map[string]interface{}
	if err := json.Unmarshal([]byte(inputJson), &mData); err != nil {
		var sData []interface{}
		if err := json.Unmarshal([]byte(inputJson), &sData); err != nil {
			return js.ValueOf(`json: cannot unmarshal array into Go value of type "map[string]interface{}" or "[]interface{}"`)
		}
		return js.ValueOf(pp.Sprint(sData))
	}

	return js.ValueOf(pp.Sprint(mData))
}
