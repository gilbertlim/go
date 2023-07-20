package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get() {
	// http.Get()
	// 간단 하지만 헤더 추가 등 세밀한 작업 불가
	res, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(data))

	// http.NewRequest()
	// 세부 작업 가능
	// request 객체 생성 후 header를 추가하고, http.client 객체를 통해 요청을 보냄
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Cawler")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)
	fmt.Println(str)
}
