// Env is a small set of routines to make easier and safer to deal with environment variables, in a more structured way
package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// EnvChecker holds instructions to assert an environment variable
type Directives struct {
	VarName      string
	DefaultValue string
	Mandatory    bool
	DebugPrint   bool
}

func debugLog(msg string, debugPrint bool) {
	if debugPrint {
		log.Println(msg)
	}
}

// Check Test environment variables according given directives
func Check(varName, defaultValue string, mandatory, debugPrint bool) error {
	if os.Getenv(varName) != "" {
		debugLog(fmt.Sprintf(`environment variable "%s" asserted`, varName), debugPrint)
		return nil
	}

	if defaultValue != "" {
		if err := os.Setenv(varName, defaultValue); err != nil {
			return nil
		}

		debugLog(fmt.Sprintf(`environment variable "%s" asserted with default value`, varName), debugPrint)
		return nil
	}

	if mandatory {
		return fmt.Errorf(`required environment variable "%s" isn't set`, varName)
	}

	return nil
}

// AsString returns the env var value as string
func AsString(key, defaultValue string) string {
	if os.Getenv(key) != `` {
		return os.Getenv(key)
	}

	return defaultValue
}

// AsStringSlice returns the env var value as []string
func AsStringSlice(key, separator string, defaultValue []string) []string {
	if os.Getenv(key) != `` {
		return strings.Split(os.Getenv(key), separator)
	}

	if len(defaultValue) > 0 {
		return defaultValue
	}

	return []string{}
}

// AsInt returns the env var value as int
func AsInt(key string, defaultValue int) int {
	if os.Getenv(key) != `` {
		if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
			return i
		}
	}

	return defaultValue
}

// AsInt64 returns the env var value as int64
func AsInt64(key string, defaultValue int64) int64 {
	if os.Getenv(key) != `` {
		if i, err := strconv.ParseInt(os.Getenv(key), 10, 64); err == nil {
			return i
		}
	}

	return defaultValue
}

// AsIntSlice returns the env var value as []int
func AsIntSlice(key, separator string, defaultValue []int) []int {
	if os.Getenv(key) != `` {
		a := strings.Split(os.Getenv(key), separator)

		is := make([]int, len(a))

		for i, x := range a {
			is[i], _ = strconv.Atoi(x)
		}

		return is
	}

	if len(defaultValue) > 0 {
		return defaultValue
	}

	return []int{}
}

// AsFloat64 returns the env var value as float64
func AsFloat64(key string, defaultValue float64) float64 {
	if os.Getenv(key) != `` {
		if f, err := strconv.ParseFloat(os.Getenv(key), 64); err == nil {
			return f
		}
	}

	return defaultValue
}

// AsBool returns the env var value as boolean
func AsBool(key string, defaultValue bool) bool {
	if os.Getenv(key) != `` {
		if b, err := strconv.ParseBool(os.Getenv(key)); err == nil {
			return b
		}
	}

	return defaultValue
}

// CheckMany Test multiple environment variables at once
func CheckMany(d ...Directives) error {
	for _, d := range d {
		if err := Check(d.VarName, d.DefaultValue, d.Mandatory, d.DebugPrint); err != nil {
			return err
		}
	}

	return nil
}
