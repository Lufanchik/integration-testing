package case5_in_path

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

	onlinePassMM  = "/mm/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	onlinePassMCD = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
)

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}

// 1 В пути с агрегацией
// 2 В пути с без агрегации
// 3 В пути + выход онлайн в мск без агрегации
// 4 В пути + выход онлайн в мск с агрегацией
// 5 В пути + выход оффлайн в мск без агрегации
// 6 В пути + выход оффлайн в мск с агрегацией
// 7 В пути + выход онлайн в мо без агрегации
// 8 В пути + выход онлайн в мо с агрегацией
// 9 В пути + выход оффлайн в мо без агрегации
// 10 В пути + выход оффлайн в мо с агрегацией
// 11 В пути + вход онлайн в мск без агрегации
// 12 В пути + вход онлайн в мск с агрегацией
// 13 В пути + вход оффлайн в мск без агрегации
// 14 В пути + вход оффлайн в мск с агрегацией
// 15 В пути + вход онлайн в мо без агрегации
// 16 В пути + вход онлайн в мо с агрегацией
// 17 В пути + вход оффлайн в мо без агрегации
// 18 В пути + вход оффлайн в мо с агрегацией

func Test_1_1(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "1"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
						"direction": 1,
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

		Request(t, reqMCD1, onlinePassMCD)
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
