package mcd

import (
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/messages/auth"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

// вход в МЦД МО2, выход в МЦД МО1, формирование complete на (T+1) день.
func Test3_5_10_2(t *testing.T) {
	pan := CardPan()
	responseI := &processing.OnlinePassResponse{}
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
					"station": "2000460",
					"direction": 1,
					"sub_carrier": 9
			},
			"card": {
					"system": 2,
					"type": 2,
					"pan": "` + pan + `",
					"bin": 47617310,
					"exp": "1224",
					"emv": "` + MKEmulatorUnsuccessWithReauth() + `",
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

		r := httpService.POST(onlinePassMCD).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), responseI)
		require.NoError(t, err)

		fmt.Println(responseI)
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
					"station": "2000055",
					"direction": 2,
					"sub_carrier": 8
			},
			"card": {
					"system": 2,
					"type": 2,
					"pan": "` + pan + `",
					"bin": 47617310,
					"exp": "1224",
					"emv": "` + MKEmulatorUnsuccessWithReauth() + `",
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

		r := httpService.POST(onlinePassMCD).WithJSON(request).
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
		httpService := httpexpect.New(t, calculatorUrl)
		now := int(time.Now().UnixNano())

		req := []byte(`{
				"pan": "` + pan + `",
				"time": ` + strconv.Itoa(now) + `
			}`)
		fmt.Println(string(req))
		request := &processing.CompleteWithCalculateRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(complete).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &processing.CompleteWithCalculateResponse{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}

	time.Sleep(time.Second * 15)

	{
		httpService := httpexpect.New(t, authUrl)
		req := []byte(`{
							"pass_id": "` + responseI.Id + `"
						}
						`)
		fmt.Println(string(req))
		request := &auth.CardByPassIDRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(reAuth).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &auth.AuthResponseEvent{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}

	time.Sleep(time.Second * 5)

	{
		httpService := httpexpect.New(t, passUrl)
		req := []byte(`{
							"id": "` + responseI.Id + `"
						}
						`)
		fmt.Println(string(req))
		request := &processing.AuthRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(authStatusMCD).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		response := &processing.AuthResponse{}
		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}
}
