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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dmitrii-lqc/golang-project-template/pkg/logger"
	"github.com/dmitrii-lqc/golang-project-template/pkg/obfuscator"
)

var jobInProgress = false

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var response Response

	if r.Method != http.MethodPost {
		logger.Error("Got request with method " + r.Method + ", expected POST instead!")
		http.Error(w, fmt.Sprintf("%d Method not allowed", http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != urlPath {
		logger.Error("Got request with invalid URL: " + r.URL.Path + ", expected " + urlPath + " instead!")
		http.Error(w, fmt.Sprintf("%d Not found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Decode the request body
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error("Failed to decode JSON input: " + err.Error())
		http.Error(w, fmt.Sprintf("%d Failed to decode JSON input", http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Validate the request body
	switch request.Command {
	case "obfuscate":
		if len(request.Arguments) == 0 {
			logger.Error("Missing arguments in the request body")
			http.Error(w, fmt.Sprintf("%d Missing arguments in the request body", http.StatusBadRequest), http.StatusBadRequest)
			return
		} else {
			// Call obfuscator
			// TODO: Add Arguments check-up
			err := obfuscator.Obfuscate(request.Arguments, fmt.Sprintf("%s-obfuscated", request.Arguments))
			if err != nil {
				logger.Error("Failed to obfuscate: " + err.Error())
				http.Error(w, fmt.Sprintf("%d Failed to obfuscate", http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	case "status":
		if len(request.Arguments) == 0 {
			logger.Error("Missing arguments in the request body")
			http.Error(w, fmt.Sprintf("%d Missing arguments in the request body", http.StatusBadRequest), http.StatusBadRequest)
			return
		} else if jobInProgress {
			// Call status checker
			response = Response{
				Status: "processing",
				ID:     "ABCD01234",
			}
		} else {
			response = Response{
				Status: "available",
				ID:     "",
			}
		}
	default:
		logger.Error("Invalid command in the request body")
		http.Error(w, fmt.Sprintf("%d Invalid command in the request body", http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
