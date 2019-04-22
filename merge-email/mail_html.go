package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"strings"
)

//Res is
type Res struct {
	Result     string      `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}

//envs
var (
	API_USER  = "sealyun"
	FROM      = "support@mail.sealyun.com"
	FROM_USER = "sealyun"
	KEY       = ""
)

func SendHtmlMail(key, to, subject, html string) {
	RequestURI := "http://api.sendcloud.net/apiv2/mail/send"
	PostParams := url.Values{
		"apiUser":        {API_USER},
		"apiKey":         {key},
		"from":           {FROM},
		"fromName":       {FROM_USER},
		"to":             {to}, //to此时为地址列表
		"subject":        {subject},
		"html":           {html},
		"useAddressList": {"false"},
	}

	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("dump send err : %s\n", to)
		return
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(BodyByte))
		fmt.Printf("dump send err : %s\n", to)
		return
	}

	res := &Res{}
	err = json.Unmarshal(BodyByte, res)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("dump send err : %s\n", to)
		return
	}
	if res.StatusCode != 200 {
		fmt.Println(string(BodyByte))
		fmt.Printf("dump send err : %s\n", to)
	}
	fmt.Printf("dump send success: %s\n", to)
}
