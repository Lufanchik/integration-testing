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

//проход в АО «Московско-Тверская пригородная пассажирская компания» (МТППК)
func Test3_5_2(t *testing.T) {
	response := &processing.OnlinePassResponse{}

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
						"station": "2000065",
						"direction": 1
				},
			"card": {
					"system": 1,
					"type": 2,
					"pan": "` + CardPan() + `",
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

		r := httpService.POST(offlinePassMTPPK).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}

	time.Sleep(time.Second * 5)

	{
		httpService := httpexpect.New(t, passUrl)
		now := int(time.Now().UnixNano())
		req := []byte(`{
							"id": "` + response.Id + `",
							"created": ` + strconv.Itoa(now) + `,
							"amount": 4400
						}
						`)
		fmt.Println(string(req))
		request := &processing.CompleteRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(completeMTPPK).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &processing.CompleteResponse{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}

	time.Sleep(time.Second * 5)

	{
		httpService := httpexpect.New(t, passUrl)
		req := []byte(`{
							"id": "` + response.Id + `"
						}
						`)
		fmt.Println(string(req))
		request := &processing.AuthRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(authStatusMTPPK).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &processing.AuthResponse{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}
}
