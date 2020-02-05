//main_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/suite"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(Handler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello World!`
	actual := recorder.Body.String()
	Equal(t, expected, actual)
}

func TestRouter(t *testing.T) {
	r := NewRouter("assets")

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status should ne ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}

func TestForNonExistentRoute(t *testing.T) {
	r := NewRouter("assets")
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status should be 405, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""
	if respString != expected {
		t.Errorf("response should be %s, got %s", expected, respString)
	}
}

func TestFileServer(t *testing.T) {
	r := NewRouter("assets")
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s got %s", expectedContentType, contentType)
	}
}

type ParseURLSuite struct {
	Suite
	Expected URLCommand
}
func (suite *ParseURLSuite) SetupTest() {
	suite.Expected = URLCommand{SourceUser: "david", DestUser: "austin", Points: 5}
}

func (suite *ParseURLSuite) TestParseURLHTTP() {
	inputURL := "http://david.gives.5.points.to/austin"
	actualOutput := parseURL(inputURL)
	Equal(suite.T(), suite.Expected, actualOutput)
}

func (suite *ParseURLSuite) TestParseURLHTTPS() {
	inputURL := "https://david.gives.5.points.to/austin"
	actualOutput := parseURL(inputURL)
	Equal(suite.T(), suite.Expected, actualOutput)
}
	
func TestParseURLSuite(t *testing.T) {
	Run(t, new(ParseURLSuite))
}