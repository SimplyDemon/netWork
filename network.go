package network

import (
	"net/http"
	"bytes"
	"log"
	"io/ioutil"
)

type Request struct {
	Url     string
	Body    string
	Headers map[string]string
}

type Response struct {
	Body    string
	Headers http.Header
}



func (request *Request) SetUrl(setUrl string) {
	request.Url = setUrl
}

func (request *Request) AddHeader(name string, value string) {
	request.Headers[name] = value
}

func (request *Request) SetBody(setBody string) {
	request.Body = setBody
}

func (request *Request) Send(method string, response Response) Response {
	req, err := http.NewRequest(method, request.Url, bytes.NewBuffer([]byte(request.Body))) // chose method, enter url and body or request
	for name, value := range request.Headers {
		//cycle for set Headers
		req.Header.Set(name, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	response.Body = string(respBody)
	response.Headers = resp.Header

	return response
}