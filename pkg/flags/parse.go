package flags

import (
	"flag"
	"log/slog"
	"os"
)

const (
	INPUT        = "input"
	SHORT_INPUT  = "i"
	OUTPUT       = "output"
	SHORT_OUTPUT = "o"
	HELP         = "help"
	VERSION      = "version"
)

// desctiption
const (
	DES_INPUT   = "path to the folder you want to back up"
	DES_OUTPUT  = "new paths where the utility will put files"
	DES_HELP    = "help"
	DES_VERSION = "shows current version"
)

// Contains a flags (args)
type Flag struct {
	Input   string // path to a folder
	Output  string // path to anather folder
	Help    bool
	Version bool
	Logger  slog.Logger
}

// Parse is constructor function
// It setup these flags:
//
//	SHORT_INPUT  = "i"
//	OUTPUT       = "output"
//	SHORT_OUTPUT = "o"
//	HELP         = "help"
//	VERSION      = "version"
func Parse(l *slog.Logger) *Flag {
	var f Flag
	f.Logger = *l // set logger

	// get home dir
	defultPath, err := os.UserHomeDir()
	if err != nil {
		l.Error("error getting home directory", "error", err)
		os.Exit(1)
	}

	// set flags
	flag.StringVar(&f.Input, INPUT, "", DES_INPUT)
	flag.StringVar(&f.Input, SHORT_INPUT, "", DES_INPUT)
	flag.StringVar(&f.Output, OUTPUT, defultPath, DES_OUTPUT)
	flag.StringVar(&f.Output, SHORT_OUTPUT, defultPath, DES_OUTPUT)

	flag.BoolVar(&f.Help, HELP, false, DES_HELP)
	flag.BoolVar(&f.Version, VERSION, false, VERSION)

	return &f

}
