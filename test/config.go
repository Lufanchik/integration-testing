package test

import "os"

var (
	ProcessingApiUrl  = os.Getenv("PROCESSING_API_URL")
	PassUrl           = os.Getenv("PASS_URL")
	ApmApiUrl         = os.Getenv("APM_API_URL")
	globalRequestType = RequestTypeOnline
)
