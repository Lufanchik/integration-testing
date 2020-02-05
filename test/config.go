package test

import (
	"os"
	"time"
)

var (
	ProcessingApiUrl  = os.Getenv("PROCESSING_API_URL")
	AuthServiceUrl    = os.Getenv("AUTH_SERVICE_URL")
	PassUrl           = os.Getenv("PASS_URL")
	ApmApiUrl         = os.Getenv("APM_API_URL")
	WebApiUrl         = os.Getenv("WEB_API_URL")
	GlobalRequestType = RequestTypeOnline
	TimeAfterRequest  = time.Millisecond * 50
)
