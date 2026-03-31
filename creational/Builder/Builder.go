/*
   We got our bob the builder here, the Builder Pattern.
   This one is very simple, It allows you to build an object at your pace that must be contructed at once with all parameters.
   With this builder class you can use polymorphism to support multiple ways of setting a param or leave it as default.
   This one helps with DRY principle as well as SOLID principles.
   Given example is very bad example of request builder, kids. Dont do this at home.
*/

package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
)

func main() {
	parms := map[string]string{
		"id": "42",
	}
	builder1 := RequestBuilder{}
	getRequest := builder1.SetUrl("https://api.example.com/v1/data").SetMethod("GET").AddHeader("Authorization: Bearer TOKEN_123").AddHeader("Accept: application/json").SetParams(parms).AddParam("type", "user").Build()
	getRequest.ExecuteRequest()

	payload := map[string]interface{}{
		"name" : "John Doe",
		"role" : "Admin",
	}
	builder2 := RequestBuilder{}
	postRequest := builder2.SetUrl("https://api.example.com/v1/users").SetMethod("POST").SetPayload(payload).SetTimeout(5000).Build()
	postRequest.ExecuteRequest()
}

type HTTPRequest struct {
	url string
	method string
	params map[string]string
	headers []string
	payload []byte
	connTimeout int
	reqTimeout int
}

func (r *HTTPRequest) ExecuteRequest() {
	fmt.Println("If you really think i would really hit this req, high is stupidity in you")
	fmt.Println("--- Debugging HttpRequest ---")
	fmt.Println("URL: " + r.url)
	fmt.Println("Method: " + r.method)
	fmt.Println("Params: ")
	if len(r.params) > 0 {
		for k, v := range r.params {
			fmt.Println("\t" + k + "=" + v)
		}
	} else {
		fmt.Println("\tNone")
	}
	fmt.Println("Headers: ")
	if len(r.headers) > 0 {
		for _, h :=  range r.headers {
			fmt.Println(h)
		}
	} else {
		fmt.Println("\tNone")
	}
	if len(r.payload) > 0 {
		fmt.Println("I aint printing the payload, size of it is " + fmt.Sprintf("%d", len(r.payload)) + " bytes")
	}
	fmt.Println("Connection Timeout: " + fmt.Sprintf("%d", r.connTimeout) + "ms")
	fmt.Println("Connection Timeout: " + fmt.Sprintf("%d", r.reqTimeout) + "ms")
	fmt.Println("-----------------------------")
}

type RequestBuilder struct {
	url string
	method string
	params map[string]string
	headers []string
	payload []byte
	connTimeout int
	reqTimeout int
}

func (b *RequestBuilder) SetUrl(url string) *RequestBuilder {
	b.url = url
	return b
}

func (b *RequestBuilder) SetMethod(method string) *RequestBuilder {
	b.method = method
	return b
}

func (b *RequestBuilder) SetParams(params map[string]string) *RequestBuilder {
	b.params = params
	return b
}

func (b *RequestBuilder) AddParam(key string, value string) *RequestBuilder {
	b.params[key] = value
	return b
}

func (b *RequestBuilder) SetHeaders(headers []string) *RequestBuilder {
	b.headers = headers
	return b
}

func (b *RequestBuilder) AddHeader(header string) *RequestBuilder {
	b.headers = append(b.headers, header)
	return b
}

func (b *RequestBuilder) SetPayload(payload any) *RequestBuilder {
    switch v := payload.(type) {
    case map[string]any:
        jsonData, _ := json.Marshal(v)
        b.payload = jsonData
    case string:
        b.payload = []byte(v)
    case []byte:
        b.payload = v
    default:
        b.payload = []byte(fmt.Sprintf("%v", v))
    }
	return b
}

func (b *RequestBuilder) SetTimeout(timeout ...int) *RequestBuilder {
	if len(timeout) > 0 {
		if len(timeout) >= 2 {
			b.connTimeout = timeout[0]
			b.reqTimeout = timeout[1]
		} else {
			b.connTimeout = timeout[0]
			b.reqTimeout = timeout[0]
		}
	}
	return b
}

func (b *RequestBuilder) Build() HTTPRequest {
	if b.url == "" {
		panic("Building HTTP request without url, are we?")
	}
	if b.method == "" {
		panic("You forgot to set method")
	}
	return HTTPRequest{
		url: b.url,
		method: b.method,
		params: maps.Clone(b.params),
		headers: slices.Clone(b.headers),
		payload: slices.Clone(b.payload),
		connTimeout: b.connTimeout,
		reqTimeout: b.reqTimeout,
	}
}