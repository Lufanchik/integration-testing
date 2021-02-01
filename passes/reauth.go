package passes

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var ReAuth = test.Cases{
	{
		N: "1. ReAuth",
		T: test.T{
			&test.ReAuth{
				Id: "3b422d69-b7d8-50be-baf7-6579ee0299c6",
			},
		},
	},
}
