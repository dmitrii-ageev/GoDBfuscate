module github.com/dmitrii-lqc/golang-project-template

go 1.22.0

replace github.com/dmitrii-lqc/golang-project-template/pkg/logger => ./pkg/logger

replace github.com/dmitrii-lqc/golang-project-template/pkg/api => ./pkg/api

replace github.com/dmitrii-lqc/golang-project-template/internal/config => ./internal/config

require (
	github.com/dmitrii-lqc/golang-project-template/pkg/api v0.0.0-00010101000000-000000000000
	github.com/dmitrii-lqc/golang-project-template/pkg/logger v0.0.0-00010101000000-000000000000
)

require (
	github.com/dmitrii-lqc/golang-project-template/internal/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/dmitrii-lqc/golang-project-template/pkg/obfuscator v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/dmitrii-lqc/golang-project-template/pkg/obfuscator => ./pkg/obfuscator
