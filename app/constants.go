package app

import (
	"sync/atomic"
)

// HEALTH is a constant for the health URI
const HEALTH = "/health"

// GET is a constant for GET request handling
const GET = "GET"

// ADDRESS_PORT is a constant for the hardcoded address and port
const ADDRESS_PORT = ":8080"

// check is the Value (bool) that will say if the service is healthy
var check atomic.Value
