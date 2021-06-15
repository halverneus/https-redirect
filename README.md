# https-redirect

## Introduction

Tiny, simple HTTP-to-HTTPS redirect server. Install from any of the following locations:

- Docker Hub: https://hub.docker.com/r/halverneus/https-redirect/
- GitHub: https://github.com/halverneus/https-redirect

## Configuration

### Environment Variables

Default values are shown with the associated environment variable.

```bash
# Enable debugging for troubleshooting. If set to 'true' this prints extra
# information during execution. IMPORTANT NOTE: The configuration summary is
# printed to stdout while logs generated during execution are printed to stderr.
DEBUG=false

# Optional Hostname for binding. Leave black to accept any incoming HTTP request
# on the prescribed port.
HOST=

# If assigned, must be a valid port number.
PORT=8080
```

### YAML Configuration File

YAML settings are individually overridden by the corresponding environment
variable. The following is an example configuration file with defaults. Pass in
the path to the configuration file using the command line option
('-c', '-config', '--config').

```yaml
debug: false
host: ""
port: 8080
```

Example configuration with possible alternative values:

```yaml
debug: true
port: 80
```

## Deployment

### Without Docker

```bash
PORT=80 ./https-redirect
```

This will automatically redirect all incoming HTTP requests to HTTPS.

### With Docker

```bash
docker run -d \
    -p "80:8080" \
    halverneus/https-redirect:latest
```

This will automatically redirect all incoming HTTP requests to HTTPS on the host machine.

Any of the variables can also be modified:

```bash
docker run -d \
    -e PORT=8888 \
    -p 80:8888 \
    halverneus/https-redirect:latest
```

### Getting Help

```bash
./redirect help
# OR
docker run -it halverneus/https-redirect:latest help
```