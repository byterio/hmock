// Copyright 2025 Byterio
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hmock

import (
	"log/slog"
	"net/http"
)

// Config holds the configuration for hmock.
type Config struct {
	// Responder generates responses for HTTP requests.
	// Defaults to returning a status OK (200) with empty body.
	Responder func(*http.Request) (*http.Response, error)

	// Logger is an optional logger used by hmock.
	// Defaults to nil (disabled).
	Logger *slog.Logger
}

// ConfigDefault provides the default configuration for hmock.
var ConfigDefault = Config{
	Responder: func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       http.NoBody,
			Header: http.Header{
				"Content-Type":   []string{"text/plain; charset=utf-8"},
				"Content-Length": []string{"0"},
			},
		}, nil
	},
	Logger: nil,
}

// configDefault sets default values for the provided config.
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	// Set default values.
	if cfg.Responder == nil {
		cfg.Responder = ConfigDefault.Responder
	}

	return cfg
}
