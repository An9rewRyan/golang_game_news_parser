package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	values := map[string]string{"link": "https://kanobu.ru/videogames/"}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post("http://localhost:8000", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	// request_body := []byte(`{"link": "https://kanobu.ru/videogames/"}`)
	// req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(request_body))
	// // req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
}
