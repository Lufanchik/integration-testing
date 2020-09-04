package passes

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var AuthResponse = test.Cases{
	{
		N: "1. AuthResponse",
		T: test.T{
			&test.AuthResponse{},
		},
	},
}
