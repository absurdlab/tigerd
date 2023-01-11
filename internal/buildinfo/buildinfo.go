package buildinfo

import "time"

var (
	// Version is the version of the tigerd executable. Its value is injected at compile time.
	Version = ""
	// CompiledAt is the time at which the program is compiled. Its format is time.RFC3339. Its value is injected
	// at compile time.
	CompiledAt = ""

	defaultCompileTime = time.Now()
)

// CompileAtTime returns the time.Time representation of CompiledAt, defaulting to a default compile time on any error.
//goland:noinspection ALL
func CompileAtTime() time.Time {
	if len(CompiledAt) == 0 {
		return defaultCompileTime
	}

	ct, err := time.Parse(time.RFC3339, CompiledAt)
	if err != nil {
		return defaultCompileTime
	}

	return ct
}
