package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	registerURL := "http://localhost:8080/auth/register"
	requestBody := RegisterRequestBody{
		Email:    "test@example.com",
		Password: "password1234",
	}
	var tests = []struct {
		name  string
		input RegisterRequestBody
		want  int
	}{
		{"Register", requestBody, http.StatusCreated},
		{"Register with same email", requestBody, http.StatusConflict},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.input)

			req, err := http.NewRequest("POST", registerURL, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("error making http request: %s\n", err)
			}
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response body: %s\n", err)
			}
			var result2 map[string]interface{}

			_ = json.Unmarshal(resBody, &result2)
			fmt.Printf("response body: %s\n", resBody)
			assert.Equal(t, tt.want, int(result2["status"].(float64)))
		})
	}
}

func TestLogin(t *testing.T) {
	loginURL := "http://localhost:8080/auth/login"
	loginBody := LoginRequestBody{
		Email:    "test@example.com",
		Password: "password1234",
	}
	loginBody2, loginBody3 := LoginRequestBody{
		Email:    "test@example.com",
		Password: "password12345",
	}, LoginRequestBody{
		Email:    "test1@example.com",
		Password: "password1234",
	}
	var tests = []struct {
		name  string
		input LoginRequestBody
		want  int
	}{
		{"Login with registered account", loginBody, http.StatusOK},
		{"Login with incorrect password", loginBody2, http.StatusNotFound},
		{"Login with non-existing email", loginBody3, http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tt.input)

			req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("error making http request: %s\n", err)
			}
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response body: %s\n", err)
			}
			var result2 map[string]interface{}

			_ = json.Unmarshal(resBody, &result2)

			fmt.Printf("response body: %s\n", resBody)
			assert.Equal(t, tt.want, int(result2["status"].(float64)))
		})
	}
}
