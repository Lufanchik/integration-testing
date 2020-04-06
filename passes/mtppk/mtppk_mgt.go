package mtppk

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesMTPPK_MGT = test.Cases{
	{
		N:          "1. МТППК - МГТ, VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				AuthType:    test.AuthTypeCorrect,
				Carrier:     carriers.Carrier_MTPPK,
				ExpectedSum: 42,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
				//ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 42,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 42,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
				//ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				Aggregate:   1,
				ExpectedSum: 42,
			},
			&test.Complete{
				StartPass: 1,
				Passes: []int{
					3, 4, 6,
				},
				Sum: 42,
			},
		},
	},
	{
		N:          "2. МТППК - МГТ, VISA",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
				ExpectedSum: 4400,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOnline,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOffline,
				ExpectedSum: 777,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOnline,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOffline,
				Aggregate:   3,
				ExpectedSum: 777,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOnline,
				Aggregate:   3,
				ExpectedSum: 777,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOnline,
				Aggregate:   3,
				ExpectedSum: 777,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				RequestType: test.RequestTypeOffline,
				//ExpectedSum: 4200,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeStartAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOnline,
				ExpectedSum: 123,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOffline,
				Aggregate:   10,
				ExpectedSum: 123,
			},
			&test.Pass{
				PaymentType: test.PaymentTypeAggregate,
				Carrier:     carriers.Carrier_MTPPK,
				RequestType: test.RequestTypeOnline,
				Aggregate:   10,
				ExpectedSum: 123,
			},
			&test.Complete{
				StartPass: 10,
				Passes: []int{
					11, 12,
				},
				Sum: 123,
			},
			&test.Complete{
				StartPass: 3,
				Passes: []int{
					5, 6, 8,
				},
				Sum: 777,
			},
		},
	},
}
