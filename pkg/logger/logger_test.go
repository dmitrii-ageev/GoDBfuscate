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

package logger

import (
	"bytes"
	"github.com/dmitrii-lqc/golang-project-template/internal/config"
	"testing"
)

func TestLogger(t *testing.T) {
	var (
		stdoutBuffer bytes.Buffer
		stderrBuffer bytes.Buffer
	)

	initLoggers(&stdoutBuffer, &stderrBuffer, 0)

	Message("Test Message")
	stdout := stdoutBuffer.String()
	if stdout != "["+config.Application+"] "+"Test Message\n" {
		t.Errorf("Expected error log: '[%s] Test Message', got: '%s'", config.Application, stdout)
	}

	stdoutBuffer.Reset()
	Debug("Test Debug Message")
	stdout = stdoutBuffer.String()
	if stdout != "["+config.Application+"] "+"DEBUG: Test Debug Message\n" {
		t.Errorf("Expected error log: '[%s] DEBUG: Test Debug Message', got: '%s'", config.Application, stdout)
	}

	stdoutBuffer.Reset()
	Info("Test Info Message")
	stdout = stdoutBuffer.String()
	if stdout != "["+config.Application+"] "+"INFO: Test Info Message\n" {
		t.Errorf("Expected error log: '[%s] INFO: Test Info Message', got: '%s'", config.Application, stdout)
	}

	stdoutBuffer.Reset()
	Warning("Test Warning Message")
	stdout = stdoutBuffer.String()
	if stdout != "["+config.Application+"] "+"WARNING: Test Warning Message\n" {
		t.Errorf("Expected error log: '[%s] WARNING: Test Warning Message', got: '%s'", config.Application, stdout)
	}

	Error("Test Error Message")
	stderr := stderrBuffer.String()
	if stderr != "["+config.Application+"] "+"ERROR: Test Error Message\n" {
		t.Errorf("Expected error log: '[%s] ERROR: Test Error Message', got: '%s'", config.Application, stderr)
	}
}
