package resolve

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesResolve = test.Cases{
	{
		N:                    "RESOLVE_PASS",
		SkipIdempotencyCheck: true,
		T: test.T{
			&test.ProcessRevisePass{},
		},
	},
}
