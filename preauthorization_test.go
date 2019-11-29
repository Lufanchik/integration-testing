package http_test

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"testing"
)

func TestSimplePreauthorization(t *testing.T) {
	Run(t, casesSimplePreauthorization)
}

var (
	casesSimplePreauthorization = Cases{
		//{
		//	&Pass{
		//		PaymentType: PaymentTypePayment,
		//		RequestType: RequestTypeOnline,
		//		Carrier:     carriers.Carrier_MM,
		//		SubCarrier:  carriers.SubCarrier_MM_SUB,
		//		ExpectedSum: 4200,
		//		AuthType:    AuthTypeIncorrect,
		//	},
		//},
	}
)

func TestSimplePreauthorizationone(t *testing.T) {
	Run(t, casesSimplePreauthorizationone)
}

var (
	casesSimplePreauthorizationone = Cases{
		{
			//&Pass{
			//	PaymentType: PaymentTypePayment,
			//	RequestType: RequestTypeOnline,
			//	Carrier:     carriers.Carrier_MM,
			//	SubCarrier:  carriers.SubCarrier_MCK_SUB,
			//	ExpectedSum: 4200,
			//	AuthType:    AuthTypeIncorrect,
			//},
		},
	}
)

func TestSimplePreauthorizationtwo(t *testing.T) {
	Run(t, casesSimplePreauthorizationtwo)
}

var (
	casesSimplePreauthorizationtwo = Cases{
		{
			N: "МЦД МО/МСК 1 - ММ - МЦД_МСК/МО_2 - МЦД МСК/МО 1", //52
			T: T{
				&Pass{
					PaymentType: PaymentTypePayment,
					RequestType: RequestTypeOnline,
					Carrier:     carriers.Carrier_MM,
					SubCarrier:  carriers.SubCarrier_MMTS_SUB,
					ExpectedSum: 4200,
					AuthType:    AuthTypeIncorrect,
				},
			},
		},
	}
)
