package test

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog/gherkin"
	"gopkg.in/resty.v1"
)

type ApiFeature struct {
	resp    *resty.Response
	BaseUrl string
}

func (a *ApiFeature) ResetResponse(interface{}) {
	a.resp = &resty.Response{}
}

func (a *ApiFeature) IDoARequest(method string, path string) (err error) {
	var resp = &resty.Response{}

	if method == "GET" {
		resp, err = resty.R().Get(a.BaseUrl + path)
	} else if method == "DELETE" {
		resp, err = resty.R().Delete(a.BaseUrl + path)
	} else {
		return fmt.Errorf("unexpected method type %s, can be either GET or DELETE", method)
	}

	if err != nil {
		return
	}

	a.resp = resp

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()
	return
}

func (a *ApiFeature) IDoARequestWithBody(method string, path string, body *gherkin.DocString) (err error) {

	var resp = &resty.Response{}

	var bodyRaw []byte
	var bodyString interface{}

	// re-encode the body string
	if err = json.Unmarshal([]byte(body.Content), &bodyString); err != nil {
		return
	}
	if bodyRaw, err = json.MarshalIndent(bodyString, "", "  "); err != nil {
		return
	}

	if method == "POST" {
		resp, err = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(bodyRaw).
			Post(a.BaseUrl + path)
	} else if method == "PUT" {
		resp, err = resty.R().
			SetBody(bodyRaw).
			Put(a.BaseUrl + path)
	} else {
		return fmt.Errorf("unexpected method type %s, can be either POST or PUT", method)
	}

	if err != nil {
		return
	}

	a.resp = resp

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()
	return
}

func (a *ApiFeature) TheResponseShouldBeAndMatchThisJson(status int, body *gherkin.DocString) (err error) {
	err = a.TheResponseShouldBe(status)
	if err != nil {
		return
	}

	var expected, actual []byte
	var exp, act interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &exp); err != nil {
		return
	}
	if expected, err = json.MarshalIndent(exp, "", "  "); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Body(), &act); err != nil {
		return
	}
	if actual, err = json.MarshalIndent(act, "", "  "); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"expected json length: %d does not match actual: %d:\n%s",
			len(expected),
			len(actual),
			string(actual),
		)
	}

	for i, b := range actual {
		if b != expected[i] {
			return fmt.Errorf(
				"expected JSON does not match actual, showing up to last matched character:\n%s",
				string(actual[:i+1]),
			)
		}
	}
	return
}

func (a *ApiFeature) TheResponseShouldBe(status int) error {
	if status != a.resp.StatusCode() {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", status, a.resp.StatusCode())
	}
	return nil
}
