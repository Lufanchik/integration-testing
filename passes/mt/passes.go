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
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "2.МЦК-МЦК",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "3.ММТС-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "4.ММ-МЦК-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePrepayed,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "5.ММТС-МЦК-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "6.МЦК-ММ-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "7.МЦК-ММТС-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "8.МЦК-ММ-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "9.ММТС-ММ-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOnline,
			},
		},
	},
	{
		N:        "10.MМ-MМ-2",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "11.МЦК-МЦК",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "12.ММТС-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "13.ММ-МЦК-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				PaymentType: test.PaymentTypePrepayed,
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "14.ММТС-МЦК-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "15.МЦК-ММ-ММТС",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "16.МЦК-ММТС-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "17.МЦК-ММ-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MMTS_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
	{
		N:        "18.ММТС-ММ-ММ",
		PassType: pass.PassType_PASS_MT,
		T: test.T{
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MCK_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
			&test.Pass{
				Carrier:     carriers.Carrier_MM,
				SubCarrier:  carriers.SubCarrier_MM_SUB,
				RequestType: test.RequestTypeOffline,
			},
		},
	},
}
