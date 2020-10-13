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
	CommentsURL       = os.Getenv("COMMENTS_SERVICE_URL")
	CardURL           = os.Getenv("CARD_SERVICE_URL")
	ReviseApiUrl      = os.Getenv("REVISE_SERVICE_URL")
	ResolveApiUrl     = os.Getenv("RESOLVE_SERVICE_URL")
	TWPGApiUrl        = os.Getenv("TWPG_SERVICE_URL")
	GlobalRequestType = RequestTypeOnline
	TimeAfterRequest  = time.Millisecond * 50 * 10
)
