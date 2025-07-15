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

type hmock struct {
	cfg       Config
	client    *http.Client
	transport *hmockTransport
}

type hmockTransport struct {
	responder func(*http.Request) (*http.Response, error)
	logger    *slog.Logger
}

// RoundTrip handles the HTTP request and returns the response.
func (mt *hmockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if mt.logger != nil {
		mt.logger.Debug("handling HTTP request", "method", req.Method, "url", req.URL.String())
	}

	resp, err := mt.responder(req)
	if err != nil {
		if mt.logger != nil {
			mt.logger.Error("responder returned an error", "error", err)
		}
		return nil, err
	}

	if mt.logger != nil {
		mt.logger.Debug("returning response", "status", resp.StatusCode)
	}

	return resp, nil
}

// New initializes and returns a new hmock instance.
func New(config ...Config) *hmock {
	cfg := configDefault(config...)

	transport := &hmockTransport{
		responder: cfg.Responder,
		logger:    cfg.Logger,
	}

	client := &http.Client{
		Transport: transport,
	}

	if transport.logger != nil {
		transport.logger.Debug("hmock is initialized")
	}

	return &hmock{
		cfg:       cfg,
		client:    client,
		transport: transport,
	}
}

// Client returns the HTTP client used by hmock.
func (m *hmock) Client() *http.Client {
	return m.client
}
