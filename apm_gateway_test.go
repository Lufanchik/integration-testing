package http_test

import (
	"testing"
)

var (
	apmGateway = Cases{
		//две бесплатные поездки, и обрыв комплексной поездки по монорельсу (может быть только последним)
		{
			&AbsGetRegistry{},
			&Login{},
		},
	}
)

func TestApmGateway(t *testing.T) {
	Run(t, apmGateway)
}
