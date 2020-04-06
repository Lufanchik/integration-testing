package apm

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var (
	CasesApmGateway = test.Cases{
		{
			N: "simple apm",
			T: test.T{
				&test.AbsGetRegistry{},
				&test.Login{},
			},
		},
	}
)
