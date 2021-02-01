package resolve

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesResolve = test.Cases{
	{
		N:                    "RESOLVE_PASS",
		SkipIdempotencyCheck: true,
		T: test.T{
			&test.ProcessRevisePass{},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				RevisePass:  1,
			},
		},
	},
}
