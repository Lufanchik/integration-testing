package mcd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"io"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	passUrl       = "http://localhost:9090"
	authUrl       = "http://localhost:9091"
	calculatorUrl = "http://localhost:9098"
	//passUrl = "http://processing-api-gateway.stage.svc.cluster.local:9090"
	//authUrl = "http://auth-service.stage.svc.cluster.local:9091"
	//calculatorUrl = "http://calculator-service.stage.svc.cluster.local:9098"

	complete         = "/twirp/sirocco.CalculatorAPI/ProcessCompleteWithCalculate"
	onlinePassMCD    = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	onlinePassMM     = "/mm/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	onlinePassAE     = "/ae/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	offlinePassMCD   = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
	offlinePassMTPPK = "/mtppk/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
	offlinePassAE    = "/ae/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
	completeMTPPK    = "/mtppk/twirp/sirocco.ProcessingAPI/ProcessComplete"
	authStatusMTPPK  = "/mtppk/twirp/sirocco.ProcessingAPI/AuthStatus"
	authStatusMCD    = "/mcd/twirp/sirocco.ProcessingAPI/AuthStatus"
	cancel           = "/mcd/twirp/sirocco.ProcessingAPI/CancelPass"

	emv  = "YyN0wu0RVJz1xgVMq7qb2udtmu8CoK5CVjyuYxIU01pXyk1TyfshuTd/zZ50hDerUXzHjAcoOGC+RBxvtL1l69Kc291LxR79jd2nQAFVT8yyLPYqPfbdQS1vg9J+UTAd+5wggbT0/RJyTJv5QL0bD5sq6xwKNFwLG8su//bSsBtaE6MqrRaubRnpYwpUweMkNPTOtJHa4h6ptZ5YmmcHHN9kSJJwFWAKcPqVSLUwEn/G//uOdkUgfR5K5D5M4svASszGMWUd313qON/kOcGQbsL5FCXbwdvzXjiiejLQaYyGldflhI6znRYOxritVCIroVv8IsZdtAJQTiOh2o7cfQ=="
	emv2 = "ozT3j/s7Z1APFxp6WXXUv6plrLAjn+vHMAc+MBTLDywxryqIShgyjH9qPD1nxrHDhcLL0HFjIaKdMbgIuZ24EfvbjRLwXtxfeq6oeSVcmivZzFMg5kw1TgmJyCycNgNB6XC6V4EZtb8KLrFOchKoalorbl2nK83VUL5b4GBPAulwoGbJ1VjUwvWt0pDaM7u9NNn0O36fZxzcsurCdCdeLNYaEvy1BbTCiOg9S7avM7VyteDMt9K0nJQXe6Hgc5UqCW5Er4QiJFuaGRWnfi9DCb2hWz4m8cnfttgYhat1ry8vCLuLZy7Y2QvvDGZ0h1PtFo3hS12H2NReObMGyHHvFw=="
	emv3 = "hmOgBFzlRO2/YgK1vcr2dyf257LncfWB0RDS+yY+y8opsiZLGC45dZECUmHu/4Fcp66FKpDT5YlnZbuoW6uwyrSzQG+trMPhF3+kMc+ifPIPn+ytQcWhxnI/5IuS9upaaWwvlNcpmZYbIpuW+9LXMI6XD9/UwTmErNTZBWD7k2ldVB0bS+xPVMG3betrWuJ9x/XGRANf+YK5pALSp2sqwB06sYgxfaI26yYeLehLynOCBIJzwYRidbq0SneBdpLZJszpzNFEWIiIVvkrlmZVYD7q0gLcyFH2agblrBViZGDOcmtCzsqewdKFAXbTePi/B/LtO/oKZ803Bu4vIxtDkA=="
	emv4 = "hy8eRekMl0BOWAEDOslK2foDc680FctBaH6IPVM5u/ZnJ5G7+HoQaLth/K9WvKi0T80LQ9vXIaGiNAhvzlfOpAKhVmF+jcuMSDZxMrZEq0HcvcuAzLDcFulJci7IbBcPalFJJz8z15eIgXzSS5wq/ANqVKYGjZc+0733M8eb0wLCjNQ8qlv5rJ959GgZZni/qm4egn05xxphe1xs9DLBoWgq3iS1031BVmDy0DXVzJqKp6bsVjPsqWIrOA3GtiLExEcvGs/inSVXl1WBKDbw9KHgVvFbSWo3JnjtL+GYEOi8sbprb1AvPuBc53gsj+n7BYTZtdjT82QvcDVcmpE85A=="
)

//вход в Московский центральный диаметр (МЦД) МСК1 (в пределах Москвы), выход в МЦД МСК2.
func Test3_5_1_1(t *testing.T) {
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
					"station": "2001060",
					"direction": 1,
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

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}
func CardPan() string {
	gofakeit.Seed(time.Now().UnixNano())
	return Generate(strconv.Itoa(gofakeit.CreditCardNumber()))
}

func Generate(d string) string {
	reader := bytes.NewReader([]byte(d))
	hash := sha256.New()

	_, err := io.Copy(hash, reader)
	if err != nil {
		panic(fmt.Errorf("sha: copy: %w", err))
	}

	hashInBytes := hash.Sum(nil)[:32]
	str := hex.EncodeToString(hashInBytes)

	return strings.ToUpper(str)
}

func CardEmvCorrect() string {
	return "ca5zfeN8X0VQGfocwEl/8WI1IGyfrOPVST9rtzVBBpPDYVKm/f1/r+gyPwRBANXrWt13q/ADa6VEoQU8u+Og0FPk2IzepajqqPEpxzHGZANjE5fygv7Hk/kblwu6Ktxj76AhU3Te1nlr5LhrfcsOLr3LbEfr2YXOi8GRX2FC/AHNJRumHCyF6r7aBB4EwoAM/gBk53y1TIvM84eq7G4b+z4w0p/le+FXb4DzuryOsj633DVEaWtMbv4A+HgqpdubQyFizkbDmiRLlcEkXztuqEJX/c1jFrR4LhA4/gU9YPcd0YggZQ53gqJ8b57HljbVAosjwuvWE4JaG4sF71sRUg=="
}
func CardEmvIncorrect() string {
	return "WovYAaagxXfxea2U/5cMzTN1MEQIJZAUTovU5NDItD4qB6PyP0rzjg3nnXjKH3XVaVMjUSHJ5YiIJ2Ji4jX3E5bj+Ufe5BntEUhjqyphe8HQz+jJdhWfld+Bm61C8yeeq3qhYbf7zgswrh3d2Gd5L6h4PlKVbbCuGLI8KOmTmfoqQXX5dtF+ZHum7l5BUIvGn3nFj1Fbkye0iKcxHXyToZd/l/M9FuRy9/klAgKYPScYlYWRSwH2I5HZs5qDKJi/cXscLjYoF6h9xxRJMsXJr68BQ5E1bx9sG5mlBb3Pzeytns8Qct9pziuUdDUnFEd2xgx07ul7jNu40k9BARcyfQ=="
}
func MKEmulatorUnsuccessWithReauth() string {
	return "b2X1GbUBWsFcfaLtjS9myMw94DLWYPKgiVgkP7rxkm/N4H6wqZbe/mfmhVIX+rjU3M1ISY9zgpFHUkqNrZLTFv7UlG3/gARhZ4fB+hrJdv7mfWfnAfbm7OknmRZiJpZzrcOiLJAhHaSnA75L1JzZdqC0gySK7qdW9nJTlCxvotNIxvztxAIb4Y9YUE/X0F5HDH4Im0yXzWa9RZB1hekxKkask5mncqTO8SHqt15btJDXoJUogFDFGkGCYhUfQJ8qNYc++Y+77ZoSniaN0bBkTR2gYwVkgB/0RXTYWxbj7U5bocNH3ztHNcc9TGHdPzyzM2Jjiy+DvLy7w0uM2uRVGQ=="
}
