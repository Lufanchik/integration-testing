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

//вход в МЦД МСК2, выход в МЦД МСК1
func Test3_5_4_3(t *testing.T) {
	pan := CardPan()
	responseE := &processing.OnlinePassResponse{}
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
					"direction": 1,
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
			"sum": 4200,
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

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), responseE)
		require.NoError(t, err)

		fmt.Println(responseE)
	}

	time.Sleep(time.Second * 5)

	{
		httpService := httpexpect.New(t, passUrl)
		req := []byte(`{
							"id": "` + responseE.Id + `"
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

	time.Sleep(time.Second * 5)
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
					"station": "2001060",
					"direction": 2,
					"sub_carrier": 6
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
			"sum": 4200,
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

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), responseI)
		require.NoError(t, err)

		fmt.Println(responseI)
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

	time.Sleep(time.Second * 5)

	{
		httpService := httpexpect.New(t, passUrl)
		req := []byte(`{
							"id": "` + responseE.Id + `"
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
