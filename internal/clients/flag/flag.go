package flag

import "github.com/spf13/pflag"

// Values contains all values parsed from command line flags
type Values struct {
	Debug bool
}

// Parse parses flags from command line
func Parse() *Values {
	values := Values{}

	pflag.BoolVarP(&values.Debug, "debug", "d", false, "Enable debug logs")

	pflag.Parse()

	return &values
}
