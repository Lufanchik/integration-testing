package mcd

import (
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

//вход в МЦД МО1 (за пределами Москвы), выход в МЦД МО2
func Test3_5_1_2(t *testing.T) {
	pan := CardPan()
	{
		httpService := httpexpect.New(t, passUrl)
		now := int(time.Now().UnixNano())

		req := []byte(`{
				"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(now) + `,
				"tap": {
				"created": ` + strconv.Itoa(now) + `,
				"resolution": 1,
				"sign": "test",
				"terminal": {
					"id": "100",
					"station": "2000055",
					"direction": 1,
					"sub_carrier": 8
			},
			"card": {
					"system": 4,
					"type": 2,
					"pan": "` + pan + `",
					"bin": 47617310,
					"exp": "1224",
					"emv": "` + CardEmvCorrect() + `",
					"token": {
					"type": 1
				}
			}
		},
		"auth": {
			"sum": 4400,
			"type": 1
		}
	}`)
		fmt.Println(string(req))
		request := &processing.OnlinePassRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(offlinePassMCD).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &processing.OnlinePassResponse{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}

	time.Sleep(time.Second * 5)

	{
		httpService := httpexpect.New(t, passUrl)
		now := int(time.Now().UnixNano())
		req := []byte(`{
				"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(now) + `,
				"tap": {
				"created": ` + strconv.Itoa(now) + `,
				"resolution": 1,
				"sign": "test",
				"terminal": {
					"id": "100",
					"station": "2001220",
					"direction": 2,
					"sub_carrier": 7
			},
			"card": {
					"system": 4,
					"type": 2,
					"pan": "` + pan + `",
					"bin": 47617310,
					"exp": "1224",
					"emv": "` + CardEmvCorrect() + `",
					"token": {
					"type": 1
				}
			}
		},
		"auth": {
			"sum": 4400,
			"type": 1
		}
	}`)
		fmt.Println(string(req))
		request := &processing.OnlinePassRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(offlinePassMCD).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &processing.OnlinePassResponse{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}
}
