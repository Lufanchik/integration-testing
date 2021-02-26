package restore

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"time"
)

var McdRestore = test.Cases{
	{
		N: "mcd_restore",
		T: test.T{
			&test.McdRestore{
				Pan:  "",
				Date: time.Now(),
			},
		},
	},
}
