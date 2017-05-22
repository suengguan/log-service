package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	_ "app-service/log-service/routers"
	"model"
)

const (
	base_url = "http://localhost:8080/v1/action"
)

func Test_Action_GetAll(t *testing.T) {
	// get all actions
	res, err := http.Get(base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
