package flags

const (
	INPUT   = "input"
	OUTPUT  = "output"
	HELP    = "help"
	VERSION = "0.0.0"
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
	Version int
}
