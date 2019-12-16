package mtppk

import (
	"lab.siroccotechnology.ru/tp/common/messages/carriers"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
)

var CasesMTPPKPasses = test.Cases{
	{
		N:          "1. MTPPK, VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 4200,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					2, 3, 4,
				},
				Sum: 16800,
			},
		},
	},
	//{
	//	N:          "2. MTPPK, MASTERCARD",
	//	CardSystem: processing.CardSystem_MASTERCARD,
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeStartAggregate,
	//			AuthType:    test.AuthTypeCorrect,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeAggregate,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			Aggregate:   1,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeAggregate,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			Aggregate:   1,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeAggregate,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			Aggregate:   1,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Complete{
	//			StartPass: 1,
	//			Passes: []int{
	//				2, 3, 4,
	//			},
	//			Sum: 16800,
	//		},
	//	},
	//},
	//{
	//	N:          "3. MTPPK, FAIL",
	//	CardSystem: processing.CardSystem_MASTERCARD,
	//	T: test.T{
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeStartAggregate,
	//			AuthType:    test.AuthTypeIncorrect,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeAggregate,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			Aggregate:   1,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeAggregate,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			Aggregate:   1,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Pass{
	//			PaymentType: test.PaymentTypeAggregate,
	//			Carrier:     carriers.Carrier_MTPPK,
	//			Aggregate:   1,
	//			ExpectedSum: 4200,
	//		},
	//		&test.Complete{
	//			StartPass: 1,
	//			Passes: []int{
	//				2, 3, 4,
	//			},
	//			Sum: 16800,
	//		},
	//	},
	//},
}
