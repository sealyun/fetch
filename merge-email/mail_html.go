package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"strings"
)

const (
	API_USER  = "sealyun"
	FROM      = "support@mail.sealyun.com"
	FROM_USER = "sealyun"
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
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(BodyByte))
}
