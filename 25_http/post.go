package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Post() {
	// String
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}

	// Form data
	values := url.Values{
		"Name": {"Lee"},
		"Age":  {"10"},
	}
	resp2, err := http.PostForm("http://httpbin.org/post", values)
	if err != nil {
		panic(err)
	}

	defer resp2.Body.Close()

	respBody2, err := io.ReadAll(resp2.Body)
	if err != nil {
		str := string(respBody2)
		fmt.Println(str)
	}

	// JSON
	// marshal(): struct -> json
	person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp3, err := http.Post("http://httpbin.org/post", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp3.Body.Close()

	respBody3, err := io.ReadAll(resp3.Body)
	if err != nil {
		str := string(respBody3)
		fmt.Println(str)
	}

	// XML
	pbyptes2, _ := xml.Marshal(person)
	xmlBuff := bytes.NewBuffer(pbyptes2)
	resp4, err := http.Post("http://httpbin.org/post", "application/xml", xmlBuff)
	if err != nil {
		panic(err)
	}

	defer resp4.Body.Close()

	respBody4, err := io.ReadAll(resp4.Body)
	if err != nil {
		str := string(respBody4)
		fmt.Println(str)
	}

	// http.NewRequest()
	req, err := http.NewRequest("POST", "httpbin.org/post", xmlBuff)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{}
	resp5, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp5.Body.Close()

	respBody5, err := io.ReadAll(resp5.Body)
	if err != nil {
		str := string(respBody5)
		fmt.Println(str)
	}

}

type Person struct {
	Name string
	Age  int
}
