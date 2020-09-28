package twpg

import (
	"github.com/google/uuid"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/messages/twpg"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CaseTWPGReverseOrder = test.Cases{
	{
		N:          "1. TWPG Reverse order",
		CardSystem: processing.CardSystem_MIR,
		CustomerId: uuid.New().String(),
		T: test.T{
			&test.TWPGCreateAndPayOrderStep{},
			&test.TWPGOrderStatus{
				OrderStatus: twpg.OrderStatus_APPROVED,
			},
			&test.TWPGReverseOrder{},
			&test.TWPGOrderStatus{
				OrderStatus: twpg.OrderStatus_REVERSED,
			},
		},
	},
}
