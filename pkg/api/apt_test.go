/*
Copyright 2024 Dmitrii Ageev <dmitrii@soft-engineering.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	testCases := []struct {
		Name          string
		RequestMethod string
		RequestBody   string
		RequestUrl    string
		ExpectedCode  int
		ExpectedBody  string
		ExpectedError string
	}{
		{
			Name:          "ValidRequest",
			RequestMethod: http.MethodPost,
			RequestBody:   `{"command": "search", "arguments": "*"}`,
			RequestUrl:    urlPath,
			ExpectedCode:  http.StatusOK,
			ExpectedBody:  `{"status":"processing","id":"ABCD01234"}` + "\n",
			ExpectedError: "",
		},
		{
			Name:          "InvalidMethod",
			RequestMethod: http.MethodGet,
			RequestBody:   "",
			RequestUrl:    urlPath,
			ExpectedCode:  http.StatusMethodNotAllowed,
			ExpectedBody:  "",
			ExpectedError: "",
		},
		{
			Name:          "InvalidURL",
			RequestMethod: http.MethodPost,
			RequestBody:   `{"command": "search", "invalid_field": "*"}`,
			RequestUrl:    "/incorrect",
			ExpectedCode:  http.StatusNotFound,
			ExpectedBody:  "",
			ExpectedError: "",
		},
		{
			Name:          "InvalidJSON",
			RequestMethod: http.MethodPost,
			RequestBody:   `{"command": "search", "invalid_field": "*"}`,
			RequestUrl:    urlPath,
			ExpectedCode:  http.StatusBadRequest,
			ExpectedBody:  "",
			ExpectedError: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var requestBody io.Reader

			if tc.RequestMethod == http.MethodPost {
				requestBody = bytes.NewBufferString(tc.RequestBody)
			} else {
				requestBody = nil
			}
			request := httptest.NewRequest(tc.RequestMethod, tc.RequestUrl, requestBody)
			recorder := httptest.NewRecorder()
			HandleRequest(recorder, request)

			response := recorder.Result()
			if response.StatusCode != tc.ExpectedCode {
				t.Errorf("Expected status code `%d`, got `%d`", tc.ExpectedCode, response.StatusCode)
			}

			if tc.ExpectedBody != "" {
				expectedBody := []byte(tc.ExpectedBody)
				responseBody, err := io.ReadAll(response.Body)
				defer response.Body.Close()

				if err != nil {
					t.Errorf("Error reading response body: %s", err)
					return
				}

				if !bytes.Equal(expectedBody, responseBody) {
					t.Errorf("Expected body `%s`, got `%s`", tc.ExpectedBody, string(responseBody))
				}
			}

			/*
				if tc.ExpectedError != "" {
					// Add assertions for expected errors if needed
				}
			*/
		})
	}
}
