package mt

import (
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var CasesMetroComplexMT = test.Cases{
	{
		N:        "1.MМ-MМ-2",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
		},
	},
	{
		N:        "2.МЦК-МЦК",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
		},
	},
	{
		N:        "3.ММТС-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
		},
	},
	{
		N:        "4.ММ-МЦК-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePrepayed,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
			},
		},
	},
	{
		N:        "6.ММТС-МЦК-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
		},
	},
	{
		N:        "7.МЦК-ММ-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
		},
	},
	{
		N:        "8.МЦК-ММТС-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
		},
	},
	{
		N:        "9.МЦК-ММ-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MMTS_SUB,
			},
		},
	},
	{
		N:        "10.ММТС-ММ-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MCK_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
			&test.Pass{
				Carrier:    carriers.Carrier_MM,
				SubCarrier: carriers.SubCarrier_MM_SUB,
			},
		},
	},
}
