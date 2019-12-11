package passes

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesCancel = test.Cases{
	//{
	//	N: "MM_successful_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MM_SUB,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},
	//
	//{
	//	N: "MCK_successful_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},
	//
	//{
	//	N: "MMTS_successful_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},
	//
	//{
	//	N: "MM_failed_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MM_SUB,
	//			ExpectedSum: 4200,
	//			AuthType:    test.AuthTypeIncorrect,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},
	//
	//{
	//	N: "MCK_failed_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MCK_SUB,
	//			ExpectedSum: 4200,
	//			AuthType:    test.AuthTypeIncorrect,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},
	//
	//{
	//	N: "MMTS_failed_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
	//			ExpectedSum: 4200,
	//			AuthType:    test.AuthTypeIncorrect,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},
	//
	//{
	//	N: "MM_successful_authorization_and_cancel",
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MM_SUB,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Cancel{
	//			Target: 1,
	//			Reason: processing.CancelPassRequest_CSS,
	//		},
	//	},
	//},

	{
		N:          "MM - отмена поездки",
		CardSystem: processing.CardSystem_VISA,
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
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 3,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 5,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 7,
				Reason: processing.CancelPassRequest_CSS,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				ExpectedSum: 4200,
			},
			&test.Cancel{
				Target: 9,
				Reason: processing.CancelPassRequest_CSS,
			},
		},
	},

	{
		N: "45. Закрытие комплексной поездки по максимальному количеству пересадок за пределами МСК (ММТС – МЦК – ММ – МЦД МСК вход – МЦД МО выход)",
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				ExpectedSum: 4200,
			},
			//&test.Cancel{
			//	Target: 1,
			//	Reason: processing.CancelPassRequest_CSS,
			//},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				//ExpectedSum: 4200,
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				Parent:      1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeFree,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MSK,
				Terminal: &processing.Terminal{
					Station:   "2000075", //ТУШИНО
					Direction: processing.TerminalDirection_INGRESS,
				},
				Parent: 1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MCD,
				SubCarrier:  carriers.SubCarrier_MCD2_MO,
				Terminal: &processing.Terminal{
					Station:   "2000460", //НАХАБИНО
					Direction: processing.TerminalDirection_EGRESS,
				},
				Ingress:     4,
				ExpectedSum: 700,
			},
		},
	},
}
