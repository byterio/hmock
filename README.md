# HMock - HTTP Mock Client for Go

[![Go Version][GoVer-Image]][GoDoc-Url] [![License][License-Image]][License-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url] [![Go Report Card][ReportCard-Image]][ReportCard-Url]

[GoVer-Image]: https://img.shields.io/badge/Go-1.24%2B-blue
[GoDoc-Url]: https://pkg.go.dev/github.com/byterio/hmock
[GoDoc-Image]: https://pkg.go.dev/badge/github.com/byterio/hmock.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/byterio/hmock
[ReportCard-Image]: https://goreportcard.com/badge/github.com/byterio/hmock?style=flat
[License-Url]: https://github.com/byterio/hmock/blob/main/LICENSE
[License-Image]: https://img.shields.io/github/license/byterio/hmock

HMock is a lightweight, flexible HTTP mock client for Go that helps you test your HTTP-dependent code with ease.

## Features

✨ **Simple API** - Works with standard `http.Client` interface  
🔧 **Customizable Responders** - Define mock responses for specific requests  
📝 **Structured Logging** - Built-in support for `slog` logging  
🚀 **Zero Dependencies** - Lightweight and easy to integrate  
🧪 **Testing Friendly** - Perfect for unit and integration tests

## Installation

```bash
go get -u github.com/byterio/hmock
```

## Quick Start

```go
package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/byterio/hmock"
)

func main() {
	// Create a mock client with custom responder
	mock := hmock.New(hmock.Config{
		Responder: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("Hello, HMock!")),
			}, nil
		},
		Logger: slog.Default(),
	})

	// Use the mock client
	client := mock.Client()
	resp, err := client.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body)) // Output: Hello, HMock!
}
```

## Advanced Usage

### Custom Response Based on Request

```go
mock := hmock.New(hmock.Config{
	Responder: func(req *http.Request) (*http.Response, error) {
		if req.URL.Path == "/api/users" {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`[{"id":1,"name":"Byterio"}]`)),
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			}, nil
		}
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       http.NoBody,
		}, nil
	},
})
```

### Error Simulation

```go
mock := hmock.New(hmock.Config{
	Responder: func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("network error")
	},
})
```

## Configuration Options

| Option      | Description                                                              | Default Value                  |
| ----------- | ------------------------------------------------------------------------ | ------------------------------ |
| `Responder` | Function that generates responses for HTTP requests                      | Returns 200 OK with empty body |
| `Logger`    | Optional `slog.Logger` for logging requests and responses (nil disables) | nil (disabled)                 |

## Feedback and Contributions

If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/byterio/hmock/issues) on GitHub.

We welcome contributions! Fork the repository, make your changes, and submit a pull request.

## Support

If you enjoy using HMock, please consider giving it a star! Your support helps others discover the project and encourages further development.

## License

HMock is open-source software released under the Apache License, Version 2.0. You can find a copy of the license in the [LICENSE](https://github.com/byterio/hmock/blob/main/LICENSE) file.
