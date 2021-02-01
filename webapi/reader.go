package webapi

import (
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"time"
)

var ReaderCase = test.Cases{
	{
		N: "1. Reader get face list",
		T: test.T{
			&test.ReaderConfiguration{
				FaceList: &test.FaceList{
					Time:     time.Now(),
					FileType: cards.FileType_SQ_LITE_BINARY_TYPE,
				},
				StopList: &test.StopList{
					Time:     time.Now(),
					FileType: cards.FileType_SQ_LITE_BINARY_TYPE,
				},
			},
		},
	},
}
