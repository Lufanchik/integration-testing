package case6_ae

import (
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/aggregation"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	passUrl = "http://localhost:9090"

	onlinePassMM   = "/mm/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	onlinePassMCD  = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	offlinePassMCD = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
	offlinePassAE  = "/ae/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
)

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}

// разрешено 2001270
// запрещено 2001340
// виза 2
// мир 4

// 1 вход ае, выход мцд, комплекс, агрегация, оффлайн
// 2 вход ае, выход мцд, не комплекс, агрегация, оффлайн
// 3 вход мцд, выход ае, комплекс (отмена мцд), агрегация, оффлайн
// 4 вход мцд, выход ае, не комплекс, агрегация, оффлайн

func Test_1(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "1"

		reqAE1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001270",
						"direction": 1,
						"sub_carrier": 0
					},
				"card": {
						"system": 2,
						"type": 1,
						"pan": "` + pan + `",
						"bin": 47617310,
						"exp": "1224",
						"emv": "` + test.CardEmvCorrect() + `",
						"token": {
							"type": 1
						}
				}
			},
			"auth": {
				"sum": 4600,
				"type": 1
			}
		}`)

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001270",
						"direction": 2,
						"sub_carrier": 1
					},
				"card": {
						"system": 2,
						"type": 1,
						"pan": "` + pan + `",
						"bin": 47617310,
						"exp": "1224",
						"emv": "` + test.CardEmvCorrect() + `",
						"token": {
							"type": 1
						}
				}
			},
			"auth": {
				"sum": 4600,
				"type": 1
			}
		}`)

		Request(t, reqAE1, offlinePassAE)
		Request(t, reqMCD1, offlinePassMCD)
	}
}

func Test_2(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "2"

		reqAE1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001270",
						"direction": 1,
						"sub_carrier": 0
					},
				"card": {
						"system": 2,
						"type": 1,
						"pan": "` + pan + `",
						"bin": 47617310,
						"exp": "1224",
						"emv": "` + test.CardEmvCorrect() + `",
						"token": {
							"type": 1
						}
				}
			},
			"auth": {
				"sum": 4600,
				"type": 1
			}
		}`)

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001340",
						"direction": 2,
						"sub_carrier": 1
					},
				"card": {
						"system": 2,
						"type": 1,
						"pan": "` + pan + `",
						"bin": 47617310,
						"exp": "1224",
						"emv": "` + test.CardEmvCorrect() + `",
						"token": {
							"type": 1
						}
				}
			},
			"auth": {
				"sum": 4600,
				"type": 1
			}
		}`)

		Request(t, reqAE1, offlinePassAE)
		Request(t, reqMCD1, offlinePassMCD)
	}
}

func Request(t *testing.T, req []byte, pe string) {
	fmt.Println(string(req))
	httpService := httpexpect.New(t, passUrl)

	request := &processing.OnlinePassRequest{}
	err := json.Unmarshal(req, request)
	require.NoError(t, err)

	r := httpService.POST(pe).WithJSON(request).
		Expect().
		Status(http.StatusOK)

	logRequest(r)

	response := &processing.OnlinePassResponse{}
	err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
	require.NoError(t, err)

	fmt.Println(response)
}
