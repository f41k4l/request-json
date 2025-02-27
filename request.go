package requestjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Client *http.Client = http.DefaultClient

type Request struct {
	BaseURL string
	headers http.Header
}

func New(baseURL string) *Request {
	return &Request{
		BaseURL: baseURL,
	}
}

func (req *Request) SetHeader(key, value string) {
	req.headers.Set(key, value)
}

func (req *Request) AddHeader(key, value string) {
	req.headers.Add(key, value)
}

// Do sends a request to the specified URL and returns the response.
// The payload can be any type that implements the `json.Marshaler` interface, and the response will be decoded into the provided struct.
func (req *Request) Do(method, path string, payload any, response any) (err error) {

	buffer := new(bytes.Buffer)
	if payload != nil {
		err = json.NewEncoder(buffer).Encode(payload)
		if err != nil {
			err = fmt.Errorf("failed encoding request body, %w", err)
			return
		}
	}

	r, err := http.NewRequest(method, req.BaseURL+path, buffer)
	if err != nil {
		err = fmt.Errorf("failed creating request, %w", err)
		return
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
	for key, values := range req.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}

	res, err := Client.Do(r)
	if err != nil {
		err = fmt.Errorf("failed sending request, %w", err)
		return
	}

	defer res.Body.Close()

	if status := res.StatusCode; status > 299 {
		body, _ := io.ReadAll(res.Body)
		err = fmt.Errorf("bad status code '%d', %s", status, body)
		return
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			err = fmt.Errorf("failed decoding response body, %w", err)
		}
	}

	return
}

// GET wrape Do for sending a GET request to the specified path and returns the response.
func (req *Request) GET(path string, response any) (err error) {
	return req.Do(http.MethodGet, path, nil, response)
}

// POST wrape Do for sending a POST request to the specified path and returns the response.
func (req *Request) POST(path string, body any, response any) (err error) {
	return req.Do(http.MethodPost, path, body, response)
}

// PUT wrape Do for sending a PUT request to the specified path and returns the response.
func (req *Request) PUT(path string, body any, response any) (err error) {
	return req.Do(http.MethodPut, path, body, response)
}

// DELETE wrape Do for sending a DELETE request to the specified path and returns the response.
func (req *Request) DELETE(path string, response any) (err error) {
	return req.Do(http.MethodDelete, path, nil, response)
}

// HEAD wrape Do for sending a HEAD request to the specified path and returns the response.
func (req *Request) HEAD(path string, response any) (err error) {
	return req.Do(http.MethodHead, path, nil, response)
}

// PATCH wrape Do for sending a PATCH request to the specified path and returns the response.
func (req *Request) PATCH(path string, body any, response any) (err error) {
	return req.Do(http.MethodPatch, path, body, response)
}

// OPTIONS wrape Do for sending an OPTIONS request to the specified path and returns the response.
func (req *Request) OPTIONS(path string, response any) (err error) {
	return req.Do(http.MethodOptions, path, nil, response)
}
