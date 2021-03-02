package restore

import (
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"time"
)

var McdRestore = test.Cases{
	{
		N: "mcd_restore",
		T: test.T{
			&test.CompleteWithCalculate{
				Pan:  "500358C1BB426F061A46743752B01BD1D1F930F7865AA439BFFA49CD11CB1F50",
				Date: time.Now().Add(-time.Hour * 24 * 8),
			},
			//&test.McdRestore{
			//	Pan:  "500358C1BB426F061A46743752B01BD1D1F930F7865AA439BFFA49CD11CB1F50",
			//	Date: time.Now().Add(-time.Hour * 24 * 8),
			//},
		},
	},
}
