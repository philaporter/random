package app

import (
	"github.com/alexcesaro/log/stdlog"
	"sync/atomic"
)

const GEN_INFO = "INFO_100: Generic information "

const GEN_WARN = "WARN_200: Generic warning "

const GEN_ERROR = "ERROR_300: Generic error "

const HEALTH = "/health"

const GET = "GET"

const ADDRESS_PORT = "0.0.0.0:80"

var check atomic.Value

var logger = stdlog.GetFromFlags()
