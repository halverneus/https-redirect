package cli

import (
	"flag"
	"fmt"

	"github.com/halverneus/https-redirect/lib/cli/help"
	"github.com/halverneus/https-redirect/lib/cli/server"
	"github.com/halverneus/https-redirect/lib/cli/version"
	"github.com/halverneus/https-redirect/lib/config"
)

// Execute the CLI arguments.
func Execute() (err error) {
	// Parse flag options, then parse command arguments.
	flag.Parse()
	args := Parse(flag.Args())

	job := selectRoutine(args)
	return job()
}

var (
	option struct {
		configFile  string
		helpFlag    bool
		versionFlag bool
	}
)

func init() {
	setupFlags()
}

func setupFlags() {
	flag.StringVar(&option.configFile, "config", "", "")
	flag.StringVar(&option.configFile, "c", "", "")
	flag.BoolVar(&option.helpFlag, "help", false, "")
	flag.BoolVar(&option.helpFlag, "h", false, "")
	flag.BoolVar(&option.versionFlag, "version", false, "")
	flag.BoolVar(&option.versionFlag, "v", false, "")
}

// Assignments used to simplify testing.
var (
	selectRoutine   = selectionRoutine
	unknownArgsFunc = unknownArgs
	runServerFunc   = server.Run
	runHelpFunc     = help.Run
	runVersionFunc  = version.Run
	loadConfig      = config.Load
)

func selectionRoutine(args Args) func() error {
	switch {

	// redirect help
	// redirect --help
	// redirect -h
	case args.Matches("help") || option.helpFlag:
		return runHelpFunc

	// redirect version
	// redirect --version
	// redirect -v
	case args.Matches("version") || option.versionFlag:
		return runVersionFunc

	// redirect
	case args.Matches():
		return withConfig(runServerFunc)

	// Unknown arguments.
	default:
		return unknownArgsFunc(args)
	}
}

func unknownArgs(args Args) func() error {
	return func() error {
		return fmt.Errorf(
			"unknown arguments provided [%v], try: 'help'",
			args,
		)
	}
}

func withConfig(routine func() error) func() error {
	return func() (err error) {
		if err = loadConfig(option.configFile); nil != err {
			return
		}
		return routine()
	}
}
