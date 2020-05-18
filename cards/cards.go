package cards

import (
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CardsStopList = test.Cases{
	{
		N:          "1. Stop list",
		CardSystem: processing.CardSystem_VISA,
		T: test.T{
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMKEmulatorUnsuccess,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeUnsuccessWithReauth,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePayment,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				AuthType:    test.AuthTypeMKEmulatorSuccess,
			},
			&test.CardStopList{
				ExpectedStatus: cards.CardStatus_BLACKLISTED_CARD,
				Passes:         []int{2},
			},
			&test.ForceReauth{
				Passes: []int{2},
			},
			&test.CardStopList{
				ExpectedStatus: cards.CardStatus_NOT_LISTED,
				Passes:         []int{2},
			},
		},
	},
}
