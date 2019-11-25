package http_test

import "os"

var (
	processingApiUrl  = os.Getenv("PROCESSING_API_URL")
	passUrl           = os.Getenv("PASS_URL")
	apmApiUrl         = os.Getenv("APM_API_URL")
	reviseApi         = os.Getenv("REVISE_URL")
	globalRequestType = RequestTypeOnline
)
