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
	"io"
	"log"
	"log/syslog"
	"os"

	"github.com/dmitrii-lqc/golang-project-template/internal/config"
)

var (
	initialised bool = false

	stdout *log.Logger
	stderr *log.Logger
)

func Init() {
	initLoggers(os.Stdout, os.Stderr)
}

func initLoggers(infoWriter, errorWriter io.Writer, flags ...int) {
	flag := log.Ldate | log.Ltime | log.Lshortfile
	if len(flags) > 0 {
		flag = flags[0]
	}
	stdout = log.New(infoWriter, "["+config.Application+"] ", flag)
	stderr = log.New(errorWriter, "["+config.Application+"] ", flag)
	initialised = true
}

func Message(message string) {
	if initialised {
		stdout.Println(message)
	} else {
		initLoggers(os.Stdout, os.Stderr)
		stdout.Println(message)
	}
}

func Debug(message string) {
	Message("DEBUG: " + message)
}

func Error(message string) {
	if initialised {
		stderr.Println("ERROR: " + message)
	} else {
		initLoggers(os.Stdout, os.Stderr)
		stderr.Println("ERROR: " + message)
	}
}

func Fatal(message string) {
	if initialised {
		stderr.Fatalln("FATAL: " + message)
	} else {
		initLoggers(os.Stdout, os.Stderr)
		stderr.Fatalln("FATAL: " + message)
	}
}

func Warning(message string) {
	Message("WARNING: " + message)
}

func Info(message string) {
	Message("INFO: " + message)
}

func Log(message string, level ...syslog.Priority) {
	if len(level) > 0 && level[0] > config.LogLevel {
		Message(config.LogPrefix() + ": " + message)
		return
	}
}
