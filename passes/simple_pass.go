package passes

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesSimplePass = test.Cases{
	{
		N: "Simple pass && cancel",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
}
