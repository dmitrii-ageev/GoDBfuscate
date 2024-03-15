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

package config

import "log/syslog"

const (
	Application = "Application"
	LogLevel    = syslog.LOG_INFO
)

func LogPrefix() string {
	var prefix string
	switch LogLevel {
	case syslog.LOG_DEBUG:
		prefix = "DEBUG"
	case syslog.LOG_INFO:
		prefix = "INFO"
	case syslog.LOG_NOTICE:
		prefix = "NOTICE"
	case syslog.LOG_WARNING:
		prefix = "WARNING"
	case syslog.LOG_ERR:
		prefix = "ERROR"
	case syslog.LOG_CRIT:
		prefix = "CRITICAL"
	case syslog.LOG_ALERT:
		prefix = "ALERT"
	case syslog.LOG_EMERG:
		prefix = "EMERGENCY"
	default:
		prefix = "UNKNOWN"
	}
	return prefix
}
