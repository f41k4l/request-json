# request-json

Simple library for requesting json in pure go

## Usage

Basic usage :

```go
package main

import (
  "http"
	"github.com/f41k4l/request-json"
)

func main() {
  req := &requestjson.Request{
    BaseURL: "http://example.com",
  }

  req.SetHeader("Authorization", "Bearer your_token")

  var response struct {
    Data string `json:"data"`
  }

  err := req.Do(http.MethodGet, "/api/data", nil, &response)
  if err != nil {
    panic(err)
  }

  fmt.Println(response.Data)
}
```

Customize http client :

```go
package main

import (
  "time"
	"github.com/f41k4l/request-json"
)

func init() {
  requestjson.Client = &http.Client{
    Timeout: time.Second * 10,  // set timeout request to 10 seconds
    Transport: &http.Transport{
      InsecureSkipVerify: true, // skip ssl verification for testing purposes
    }
  }
}
```
