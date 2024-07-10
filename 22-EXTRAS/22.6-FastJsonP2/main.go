package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var p fastjson.Parser

	jsonData := `{"user": {"name": "John Doe", "age": 30}}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := v.GetObject("user")
	fmt.Printf("User name: %v\n", user.Get("name"))
	fmt.Printf("User age: %v\n", user.Get("age"))

	userJSON := v.GetObject("user").String()

	var userStruct User
	if err := json.Unmarshal([]byte(userJSON), &userStruct); err != nil {
		panic(err)
	}

	fmt.Printf("User name: %v is %v years old.\n", userStruct.Name, userStruct.Age)
}
