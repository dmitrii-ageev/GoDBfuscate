module github.com/dmitrii-ageev/GoDBfuscate

go 1.22.0

replace github.com/dmitrii-ageev/GoDBfuscate/pkg/logger => ./pkg/logger

replace github.com/dmitrii-ageev/GoDBfuscate/pkg/api => ./pkg/api

replace github.com/dmitrii-ageev/GoDBfuscate/internal/config => ./internal/config

require (
	github.com/dmitrii-ageev/GoDBfuscate/pkg/api v0.0.0-00010101000000-000000000000
	github.com/dmitrii-ageev/GoDBfuscate/pkg/logger v0.0.0-00010101000000-000000000000
)

require (
	github.com/dmitrii-ageev/GoDBfuscate/internal/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/dmitrii-ageev/GoDBfuscate/pkg/obfuscator v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/dmitrii-ageev/GoDBfuscate/pkg/obfuscator => ./pkg/obfuscator
