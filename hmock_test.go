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
	"errors"
	"io"
	// "log/slog"
	"net/http"
	// "os"
	"strings"
	"testing"
)

func TestConfigDefault(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cfg := ConfigDefault
		if cfg.Responder == nil {
			t.Error("expected default responder to be set")
		}
	})
}

func TestClient(t *testing.T) {
	m := New()
	client := m.Client()
	if client == nil {
		t.Fatal("expected client to be returned")
	}
}

func TestIntegration(t *testing.T) {
	// handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	// logger := slog.New(handler)

	t.Run("default responder integration", func(t *testing.T) {

		m := New(Config{
			// Logger: logger,
		})

		client := m.Client()

		resp, err := client.Get("http://example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code 200, got %d", resp.StatusCode)
		}
	})

	t.Run("custom responder integration", func(t *testing.T) {
		m := New(Config{
			Responder: func(req *http.Request) (*http.Response, error) {
				if req.URL.Path == "/users" {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`[{"name": "byterio"}]`)),
					}, nil
				}
				if req.URL.Path == "/users/create" {
					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(strings.NewReader(`{"name": "hmock"}`)),
					}, nil
				}

				return nil, errors.New("not found")

			},
			// Logger: logger,
		})

		client := m.Client()

		resp, err := client.Get("http://example.com/users")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code 200, got %d", resp.StatusCode)
		}

		resp, err = client.Get("http://example.com/users/create")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status code 201, got %d", resp.StatusCode)
		}
	})
}
