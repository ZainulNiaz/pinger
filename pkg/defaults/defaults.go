package defaults

import (
	"os"
	"path/filepath"
	"time"
)

var (
	// MinControllerRetryInterval is the minimum value of the retry interval that the controller accepts.
	MinControllerRetryInterval = time.Second * 4

	// ControllerRetryInterval is the default value of the retry interval used by
	// the controller.
	ControllerRetryInterval = time.Second * 30

	// MinControllerTimeout is the minimum value of the permissible timeout
	// that the controller accepts.
	MinControllerTimeout = time.Second * 5

	// DefaultHTTPMethod is the default method to use for HTTP Input
	DefaultHTTPMethod = "GET"

	// DefaultGRPCRequestTimeout is the default timeout for the GRPC request we
	// are making between the server and the agent and vice versa.
	DefaultGRPCRequestTimeout = time.Second * 10

	// DefaultHTTPProbeTimeout is the http request timeout for HTTP prober.
	DefaultHTTPProbeTimeout = time.Second * 10

	// JWTExpireInterval is interval for which JWT is valid
	JWTExpireInterval = time.Hour * 24

	// JWTContextKey is the key by which JWT claims are stored in Gin Context bucket
	JWTContextKey = "currentUser"

	// JWTAuthType is the token type of the authorization header
	JWTAuthType = "Bearer"

	// DefaultAgentPrometheusMetricsPort is the default port for the agent to host
	// prometheus metrics enpoint on.
	DefaultAgentPrometheusMetricsPort = 9008

	// DefaultAgentPort is the default value of port to run the status agent on
	DefaultAgentPort = 9009

	// DefaultControllerType is the default type of the controller, other types can be
	// specifed based on the type of checks/probes that the controller executes
	DefaultControllerType = "default"

	// DefaultInvalidDuration is the default value of time duration which will be considered
	// as a failed controller execution
	DefaultInvalidDuration = time.Second * 360

	// DefaultStatusPageConfigPath is the default path of the config file for the status page
	// agent. This is mostly used in case when we run the agent in a standalone mode.
	DefaultStatusPageConfigPath = filepath.Join(os.Getenv("HOME"), "/.sp.yml")
)
