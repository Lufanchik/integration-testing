package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
)

func TestCancel(t *testing.T) {
	Run(t, casesCancel)
}

var (
	casesCancel = Cases{
		{
			N: "MM_successful_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
				},
				&Cancel{
					Target: 1,
					Reason: processing.CancelPassRequest_CSS,
				},
			},
		},

		{
			N: "MCK_successful_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
				},
				&Cancel{
					Target: 1,
					Reason: processing.CancelPassRequest_CSS,
				},
			},
		},

		{
			N: "MMTS_successful_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
				},
				&Cancel{
					Target: 1,
					Reason: processing.CancelPassRequest_CSS,
				},
			},
		},

		{
			N: "MM_failed_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MM_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Cancel{
					Target: 1,
					Reason: processing.CancelPassRequest_CSS,
				},
			},
		},

		{
			N: "MCK_failed_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MCK_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Cancel{
					Target: 1,
					Reason: processing.CancelPassRequest_CSS,
				},
			},
		},

		{
			N: "MMTS_failed_authorization_and_cancel",
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
				&Cancel{
					Target: 1,
					Reason: processing.CancelPassRequest_CSS,
				},
			},
		},
	}
)
