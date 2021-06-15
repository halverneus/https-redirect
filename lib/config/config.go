package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	// Get the desired configuration value.
	Get struct {
		Debug bool   `yaml:"debug"`
		Host  string `yaml:"host"`
		Port  uint16 `yaml:"port"`
	}
)

const (
	debugKey = "DEBUG"
	hostKey  = "HOST"
	portKey  = "PORT"
)

var (
	defaultDebug = false
	defaultHost  = ""
	defaultPort  = uint16(8080)
)

func init() {
	// init calls setDefaults to better support testing.
	setDefaults()
}

func setDefaults() {
	Get.Debug = defaultDebug
	Get.Host = defaultHost
	Get.Port = defaultPort
}

// Load the configuration file.
func Load(filename string) (err error) {
	// If no filename provided, assign envvars.
	if filename == "" {
		overrideWithEnvVars()
		return
	}

	// Read contents from configuration file.
	var contents []byte
	if contents, err = ioutil.ReadFile(filename); nil != err {
		return
	}

	// Parse contents into 'Get' configuration.
	if err = yaml.Unmarshal(contents, &Get); nil != err {
		return
	}

	overrideWithEnvVars()
	return
}

// Log the current configuration.
func Log() {
	// YAML marshalling should never error, but if it could, the result is that
	// the contents of the configuration are not logged.
	contents, _ := yaml.Marshal(&Get)

	// Log the configuration.
	fmt.Println("Using the following configuration:")
	fmt.Println(string(contents))
}

// overrideWithEnvVars the default values and the configuration file values.
func overrideWithEnvVars() {
	// Assign envvars, if set.
	Get.Debug = envAsBool(debugKey, Get.Debug)
	Get.Host = envAsStr(hostKey, Get.Host)
	Get.Port = envAsUint16(portKey, Get.Port)
}

// envAsStr returns the value of the environment variable as a string if set.
func envAsStr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// envAsUint16 returns the value of the environment variable as a uint16 if set.
func envAsUint16(key string, fallback uint16) uint16 {
	// Retrieve the string value of the environment variable. If not set,
	// fallback is used.
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}

	// Parse the string into a uint16.
	base := 10
	bitSize := 16
	valueAsUint64, err := strconv.ParseUint(valueStr, base, bitSize)
	if nil != err {
		log.Printf(
			"Invalid value for '%s': %v\nUsing fallback: %d",
			key, err, fallback,
		)
		return fallback
	}
	return uint16(valueAsUint64)
}

// envAsBool returns the value for an environment variable or, if not set, a
// fallback value as a boolean.
func envAsBool(key string, fallback bool) bool {
	// Retrieve the string value of the environment variable. If not set,
	// fallback is used.
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}

	// Parse the string into a boolean.
	value, err := strAsBool(valueStr)
	if nil != err {
		log.Printf(
			"Invalid value for '%s': %v\nUsing fallback: %t",
			key, err, fallback,
		)
		return fallback
	}
	return value
}

// strAsBool converts the intent of the passed value into a boolean
// representation.
func strAsBool(value string) (result bool, err error) {
	lvalue := strings.ToLower(value)
	switch lvalue {
	case "0", "false", "f", "no", "n":
		result = false
	case "1", "true", "t", "yes", "y":
		result = true
	default:
		result = false
		msg := "unknown conversion from string to bool for value '%s'"
		err = fmt.Errorf(msg, value)
	}
	return
}
