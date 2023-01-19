package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)



type heander struct{
	Key string
	Value string
}

var(
	localhost ="http://localhost:8080"
	testHost = ""
)

func PerformRequest(method, path string, req, res interface{}, headers ...heander)(*http.Response, error)  {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err 
	}

	client := &http.Client{}
	request, err := http.NewRequestWithContext(context.Background(),method, fmt.Sprintf("%s%s", localhost, path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err 
	}
	for _, h := range headers {
		request.Header.Add(h.Key, h.Value)
	}

	request.Header.Add("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(resp_body, res)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
