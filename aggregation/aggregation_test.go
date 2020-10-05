package aggregation

import (
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/test"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	passUrl = "http://localhost:9090"
	authUrl = "http://localhost:9091"
	//passUrl = "http://processing-api-gateway.test.svc.cluster.local:9090"
	//authUrl = "http://auth-service.test.svc.cluster.local:9091"

	onlinePass   = "/mm/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	complete     = "/mtppk/twirp/sirocco.ProcessingAPI/ProcessComplete"
	cancel       = "/mtppk/twirp/sirocco.ProcessingAPI/CancelPass"
	activeReAuth = "/twirp/sirocco.AuthAPI/ActiveReAuth"

	// 17.02.2020
	//emv = "nOK9Iah6LQpqHTPB7zmtdqtfAvQ0UYLmKQc52V6MrDCqsmorbH2HppQs7eRNQDZAHWGv77r60q6wr2WjR90eixyjL+84wpRVWhMLN9+pJQ0C74D6/fcIJErYCek9SwxhLgi2Sp0fJqUj3G40EeupmVST1CeX5HcprnrDnueaOPc2fUbUm/Q1RpM63bbxY2dX/JF2RaoSogTos7xbQKnzb1sJs4ISrw7+4UIl68mg7yN3cB2CnnydKWjCfaQkIwieGDyXfm4Z0XuEZeyqz1gik+dgc/xrfSdlaBfnVeSCyxX94yMTw8s1KwgVrrnM0KJuptRYyFBKvjmgI2sZMHW1Ew=="

	// 18.02.2020
	//emv = "SM9Y3VWP81jHw3JeLVxy9f3s3gcUNCk5meuQ0rb4SFBl4rph4Wo90gYXETNce6Ofn7WV1pF5XJOE7yNEtb6tfTQReFNOt84WpPONyCiUO6y82GguE8if+TvyJc+bg+bsUXc9MPfR6hYfNM+hArFTi1+NWZdsMBjrytOp4IzkuBCW3D92coqq0SX8/CgsWuNW0Hk5dtkSIFOxoT3RQZLZXdi4s/BFeouipDTV14NCEXVC6yJruPUqBTqK68eW2vvoQTivFZtfzwYjxQU3z5xhxTZIPipXHAaHIME7V1POUJ7I4lQ6ss+gz10Zw1etVvJ0WiBLY9fw4eBouTn2zS/r+Q=="

	// 15.04.2020
	//emv ="NhIdUq5lxb9vE5sHixNo/AOswgAh0cvJM0ntZdNhYkz4/y5JAAGImUEvtR8FkOMW5V6C5ney0qVScYH7hEWmPxxMDiZ0Q35QkAni1b8ojXhYvTswj53QuSp1s7mbb+CbfhBkM0Y1zHPNmgGorApsJiedUg573nlD5nFmjVsSUCLe7MDiU/hyGkQN/Hb7BDZFZVaRGqYCNVghIMoCxijFYq7XaX5XlqREqM75Jtc0L2U4QmFxD3vzpbWPSwbvUAySyLw0Z30/fbyrheOiXy7zrC8TyNpPV97D6Dl29D08bCB8nHpfKITQRsf+6mEGSx4gn7kOm0Crwmbc/Mk3lN5dIA=="

	// 16.04.2020
	//emv = "in1srhBjh8bW0dzKnHiAyBZp/rDjhb8JhppE+FgISzUbuvuQAsjaECUiGtEjkVz16Jc2EPB6ULI2jAAtnmjLEEOFWPb+e3mSCRrlluK65oyyswYPxKpGzzf57fn7iLiK/7BcmckG8tx3v0PsT/kXPrN9WKf074T4bwnoUJqhSBPOMkeLneTtJQycqCFs9kno6P9f0nfUEJR8oBeb8CugJWh7ym0OQtUQLzyvhu3QJ3DjmSy7Nu6LZ0VQdhas/T+T8xCBE+KC3GH5zOJGXyQp8q4GlgbComP7wXBVl9srEEo9PavjAeVviHlfpXTUPhy14H+7mET5UMIIrnNbXbWsoA=="

	//emv = "ftlQAWAPd4U4mhUACaR8FiF3V4I/47oy71Yx4Gyb79cLQ/bK9K1Ojq7N9vZwkPfCfOqlBEDf6G0oEz43EyrxaqPnc8/Ef0UpvpuzlWw6VdeKZ/fSp7l8LHhqPqQRcEk3kOnQd73z7GnHCxHYxKTk47dgiQpONH0pfXTEfEAahLKmtaxrseiv0mShzEOZ7jJtRiPt5hCFUVFfKsShM+yuvkaVsNyHUlyGQQi6GRjFj8c32PCl2n0VjX/iy9rVhEdphO6SRQb+ABtrwWWBrWRWeDlvGD+/V2aYfYDohR0rmbKN/f6jQ+H1hzyT4960YFZSiph7Sr+b656J2yyBLh7otQ=="

	emv  = "YyN0wu0RVJz1xgVMq7qb2udtmu8CoK5CVjyuYxIU01pXyk1TyfshuTd/zZ50hDerUXzHjAcoOGC+RBxvtL1l69Kc291LxR79jd2nQAFVT8yyLPYqPfbdQS1vg9J+UTAd+5wggbT0/RJyTJv5QL0bD5sq6xwKNFwLG8su//bSsBtaE6MqrRaubRnpYwpUweMkNPTOtJHa4h6ptZ5YmmcHHN9kSJJwFWAKcPqVSLUwEn/G//uOdkUgfR5K5D5M4svASszGMWUd313qON/kOcGQbsL5FCXbwdvzXjiiejLQaYyGldflhI6znRYOxritVCIroVv8IsZdtAJQTiOh2o7cfQ=="
	emv2 = "ozT3j/s7Z1APFxp6WXXUv6plrLAjn+vHMAc+MBTLDywxryqIShgyjH9qPD1nxrHDhcLL0HFjIaKdMbgIuZ24EfvbjRLwXtxfeq6oeSVcmivZzFMg5kw1TgmJyCycNgNB6XC6V4EZtb8KLrFOchKoalorbl2nK83VUL5b4GBPAulwoGbJ1VjUwvWt0pDaM7u9NNn0O36fZxzcsurCdCdeLNYaEvy1BbTCiOg9S7avM7VyteDMt9K0nJQXe6Hgc5UqCW5Er4QiJFuaGRWnfi9DCb2hWz4m8cnfttgYhat1ry8vCLuLZy7Y2QvvDGZ0h1PtFo3hS12H2NReObMGyHHvFw=="
)

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}

func Test_1_1(t *testing.T) {
	response := &processing.OnlinePassResponse{}
	{
		httpService := httpexpect.New(t, passUrl)
		now := int(time.Now().UnixNano())
		//Метро, онлайн, мир - успешная авторизация, акция
		//	req := []byte(`{
		//	"id": "` + uuid.New().String() + `",
		//		"created": ` + strconv.Itoa(now) + `,
		//		"tap": {
		//		"created": ` + strconv.Itoa(now) + `,
		//			"resolution": 1,
		//			"sign": "test",
		//			"terminal": {
		//				"id": "3213",
		//				"station": "2000285",
		//				"direction": 1,
		//				"sub_carrier": 3
		//			},
		//		"card": {
		//				"system": 4,
		//				"type": 1,
		//				"pan": "61312C7296E8F21FC86A237E632C34FD10FE7C532A130F98EBFA4E037DB672FD",
		//				"bin": 47617310,
		//				"exp": "1224",
		//				"emv": "` + test.CardEmvCorrect() + `",
		//				"token": {
		//					"type": 1
		//				}
		//		}
		//	},
		//	"auth": {
		//		"sum": 4200,
		//		"type": 1
		//	}
		//}`)

		// Метро, оффлайн, виза - успешный обычный чеккард, не нашли чеккард
		req := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(now) + `,
				"tap": {
				"created": ` + strconv.Itoa(now) + `,
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
						"pan": "61312C7296E8F21FC86A237E632C34FD10FE7C532A130F98EBFA4E037DB67212",
						"bin": 47617310,
						"exp": "1224",
						"emv": "` + test.CardEmvCorrect() + `",
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

		//Метро, оффлайн, мастер - успешный обычный чеккард, не нашли чеккард
		//	req := []byte(`{
		//	"id": "` + uuid.New().String() + `",
		//		"created": ` + strconv.Itoa(now) + `,
		//		"tap": {
		//		"created": ` + strconv.Itoa(now) + `,
		//			"resolution": 1,
		//			"sign": "test",
		//			"terminal": {
		//				"id": "3213",
		//				"station": "2000285",
		//				"direction": 1,
		//				"sub_carrier": 3
		//			},
		//		"card": {
		//				"system": 3,
		//				"type": 1,
		//				"pan": "61312C7296E8F21FC86A237E632C34FD10FE7C532A130F98EBFA4E037DB672FA",
		//				"bin": 47617310,
		//				"exp": "1224",
		//				"emv": "` + test.CardEmvCorrect() + `",
		//				"token": {
		//					"type": 1
		//				}
		//		}
		//	},
		//	"auth": {
		//		"sum": 4200,
		//		"type": 1
		//	}
		//}`)

		fmt.Println(string(req))
		request := &processing.OnlinePassRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(onlinePass).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}
}
