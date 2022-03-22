package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	values := map[string]string{"link": "https://kanobu.ru/videogames/"}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post("http://localhost:8000", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bodyBytes))
}
