module github.com/dmitrii-ageev/GoDBfuscate/pkg/api

go 1.22.0

replace github.com/dmitrii-ageev/GoDBfuscate/pkg/logger => ../logger

replace github.com/dmitrii-ageev/GoDBfuscate/internal/config => ../../internal/config

require (
	github.com/dmitrii-ageev/GoDBfuscate/pkg/logger v0.0.0-00010101000000-000000000000
	github.com/dmitrii-ageev/GoDBfuscate/pkg/obfuscator v0.0.0-00010101000000-000000000000
)

require github.com/dmitrii-ageev/GoDBfuscate/internal/config v0.0.0-00010101000000-000000000000 // indirect

replace github.com/dmitrii-ageev/GoDBfuscate/pkg/obfuscator => ../obfuscator
