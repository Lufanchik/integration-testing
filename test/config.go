package test

import (
	"os"
	"time"
)

var (
	ProcessingApiUrl  = os.Getenv("PROCESSING_API_URL")
	PassUrl           = os.Getenv("PASS_URL")
	ApmApiUrl         = os.Getenv("APM_API_URL")
	WebApiUrl         = os.Getenv("WEB_API_URL")
	GlobalRequestType = RequestTypeOffline
	TimeAfterRequest  = time.Millisecond * 500
	TimeBeforeRecheck = time.Millisecond * 1000
)
