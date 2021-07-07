package utils

import (
	"fmt"
	"net/http"
	"time"
)

func DoGetReq(BaseUrl string, Time_Out int) (*http.Response, error) {

	client := &http.Client{
		Timeout: time.Duration(Time_Out) * time.Second,
	}

	req, err := http.NewRequest("GET", BaseUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//set headers from utils stuffs in req
	GetHeaders(req)

	res, err := client.Do(req)

	return res, err

}
