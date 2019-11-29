package passes

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesCancel = test.Cases{
	{
		N: "MM_successful_authorization_and_cancel",
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

	{
		N: "MCK_successful_authorization_and_cancel",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},

	{
		N: "MMTS_successful_authorization_and_cancel",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},

	{
		N: "MM_failed_authorization_and_cancel",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},

	{
		N: "MCK_failed_authorization_and_cancel",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},

	{
		N: "MMTS_failed_authorization_and_cancel",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
				AuthType:    test.AuthTypeIncorrect,
			},
			&test.Cancel{
				Target: 1,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},
}
