package help

import (
	"fmt"
)

// Run print operation.
func Run() error {
	fmt.Println(Text)
	return nil
}

var (
	// Text for directly accessing help.
	Text = `
NAME
    https-redirect
SYNOPSIS
    https-redirect
    https-redirect [ -c | -config | --config ] /path/to/config.yml
    https-redirect [ help | -help | --help ]
    https-redirect [ version | -version | --version ]
DESCRIPTION
    The HTTPS Redirect is intended to be a tiny, fast and simple solution
    for redirecting incoming connections from HTTP to HTTPS. The features
	included are limited to make to binding to a host name and port. If you
	want really awesome reverse proxy features, I recommend Nginx.
DEPENDENCIES
    None... not even libc!
ENVIRONMENT VARIABLES
    DEBUG
        When set to 'true' enables additional logging, including the
        configuration used and an access log for each request. IMPORTANT NOTE:
        The configuration summary is printed to stdout while logs generated
        during execution are printed to stderr. Default value is 'false'.
    HOST
        The hostname used for binding. If not supplied, contents will be served
        to a client without regard for the hostname.
    PORT
        The port used for binding. If not supplied, defaults to port '8080'.
CONFIGURATION FILE
    Configuration can also managed used a YAML configuration file. To select the
    configuration values using the YAML file, pass in the path to the file using
    the appropriate flags (-c, --config). Environment variables take priority
    over the configuration file. The following is an example configuration using
    the default values.
    Example config.yml with defaults:
    ----------------------------------------------------------------------------
    debug: false
    host: ""
    port: 8080
    ----------------------------------------------------------------------------
    Example config.yml with possible alternative values:
    ----------------------------------------------------------------------------
    debug: true
    port: 80
    ----------------------------------------------------------------------------
USAGE
    COMMAND
		https-redirect
		    Available listening on port 8080 of host.
		export PORT=80
		https-redirect
			Available listening on port 80 (standard HTTP) of host.
		export PORT=80
        https-redirect -c config.yml
            Result: Runs with values from config.yml, but with the port being
                    overridden by the PORT environment variable.
`
)
