package main

import (
	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser

	jsonData := `{"name":"John","age":30,"city":"New York", "bool": true, "float": 3.14, "array": [1, 2, 3], "object": {"key": "value"}}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	name := v.GetStringBytes("name")
	age := v.GetInt("age")
	city := v.GetStringBytes("city")
	boolValue := v.GetBool("bool")
	floatValue := v.GetFloat64("float")

	println(string(name), age, string(city), boolValue, floatValue)

	array := v.GetArray("array")
	for i := range array {
		println(array[i].GetInt())
	}

	object := v.GetObject("object")
	key := object.Get("key")
	println(string(key.GetStringBytes()))

}
