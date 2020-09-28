package cards

import (
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
)

var FaceList = test.Cases{
	{
		N:                    "1. Face list",
		CardSystem:           processing.CardSystem_MIR,
		SkipIdempotencyCheck: true,
		T: test.T{
			&test.CardGetFull{
				Kind:     cards.DiffExportKind_FACE_IDS_WHITELIST,
				FileType: cards.FileType_SQ_LITE_BINARY_TYPE,
			},
		},
	},
}
