// HTTTClient com Timeout

package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"jean"}`))
	resp, err := c.Post("http://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
