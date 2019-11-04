package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestComplexPassMCK(t *testing.T) {
	Passes(t, ComplexPassMCK)
}

var ComplexPassMCK = Cases{
	//Пересадка с МЦК на МЦК
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММ на ММ
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММ на МЦК
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММТС на ММ
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММТС на МЦК
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММТС на ММТС
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММ на ММТС на ММ
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММ на ММТС на МЦК
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
	},
	//Пересадка с МЦК на ММ на ММТС на ММТС
	{
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MCK_SUB,
			ExpectedSum: 4200,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MM_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFree,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			Parent:      1,
		},
		&Pass{
			PaymentType: PaymentTypeFullPayment,
			RequestType: RequestTypeOnline,
			Carrier:     carriers.Carrier_MM,
			SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			ExpectedSum: 4200,
		},
	},
}
