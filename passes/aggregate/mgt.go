package aggregate

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var aggregateMgtCard = &processing.Card{
	System: processing.CardSystem_VISA,
	Type:   processing.CardType_DEBIT,
	Pan:    test.CardPan(),
	Bin:    test.CardBin(),
	Exp:    test.CardExp(),
	Token: &processing.Token{
		Type: processing.Token_SAMSUNG_PAY,
	},
}

var MgtAggregate = test.Cases{
	{
		N:              "1. Mgt Aggregate Passes pan:" + aggregateMgtCard.Pan,
		CardSystem:     processing.CardSystem_VISA,
		Card:           aggregateMgtCard,
		NotDoubleCheck: true,
		T: test.T{
			&test.Pass{
				PaymentType:          test.PaymentTypePayment,
				Carrier:              carriers.Carrier_MGT,
				IsInitAggregate:      true,
				ExpectedSumAggregate: 9200,
				Aggregate:            1,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MGT,
				Aggregate:   1,
			},
			&test.CompleteWithCalculate{
				Pan: aggregateMgtCard.Pan,
			},
		},
	},
}
