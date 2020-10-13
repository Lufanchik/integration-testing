package aggregate

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var aggregateMMCard = &processing.Card{
	System: processing.CardSystem_VISA,
	Type:   processing.CardType_DEBIT,
	Pan:    test.CardPan(),
	Bin:    test.CardBin(),
	Exp:    test.CardExp(),
	Token: &processing.Token{
		Type: processing.Token_SAMSUNG_PAY,
	},
}

var MetroAggregate = test.Cases{
	{
		N:              "1. Metro Aggregate Passes pan:" + aggregateMMCard.Pan,
		CardSystem:     processing.CardSystem_VISA,
		Card:           aggregateMMCard,
		NotDoubleCheck: true,
		T: test.T{
			&test.Pass{
				PaymentType:          test.PaymentTypePayment,
				Carrier:              carriers.Carrier_MM,
				IsInitAggregate:      true,
				ExpectedSumAggregate: 13200,
				Aggregate:            1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				Aggregate:   1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				Aggregate:   1,
			},
			&test.CompleteWithCalculate{
				Pan: aggregateMMCard.Pan,
			},
		},
	},
}
