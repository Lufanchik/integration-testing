package mk_emulator

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesEmulatorMK = test.Cases{
	{
		N:          "Switch ON Success mode",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMKEmulatorSuccess,
			},
		},
	},
	{
		N:          "Success pass",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeCorrect,
			},
		},
	},
	{
		N:          "Switch OFF Success mode",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMKEmulatorUnsuccess,
			},
		},
	},
	{
		N:          "Unsuccess pass with reauth",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeUnsuccessWithReauth,
			},
		},
	},
	{
		N:          "Unsuccess pass without reauth",
		CardSystem: processing.CardSystem_MIR,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				RequestType: test.RequestTypeOnline,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeUnsuccessWithoutReauth,
			},
		},
	},
	//{
	//	N:          "Switch ON Random mode",
	//	CardSystem: processing.CardSystem_MIR,
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypePayment,
	//			RequestType: test.RequestTypeOnline,
	//			Carrier:     carriers.Carrier_MM,
	//			SubCarrier:  carriers.SubCarrier_MM_SUB,
	//			AuthType:    test.AuthTypeMKEmulatorRandom,
	//		},
	//	},
	//},
}
