package aggregation_complete

import (
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect"
	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/integration-testing/aggregation"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	passUrl = "http://localhost:9090"
	authUrl = "http://localhost:9091"
	webUrl  = "http://127.0.0.1:1344"
	//passUrl = "http://processing-api-gateway.test.svc.cluster.local:9090"
	//authUrl = "http://auth-service.test.svc.cluster.local:9091"

	completeCalc = "/twirp/proto.WebAPIGateway/ProcessCompleteWithCalculate"
	onlinePass   = "/mm/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	complete     = "/mtppk/twirp/sirocco.ProcessingAPI/ProcessComplete"
	cancel       = "/mtppk/twirp/sirocco.ProcessingAPI/CancelPass"
	activeReAuth = "/twirp/sirocco.AuthAPI/ActiveReAuth"
)

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}

func Test_1_1(t *testing.T) {
	response := &processing.CompleteWithCalculateResponse{}
	{
		httpService := httpexpect.New(t, webUrl)
		now := int(time.Now().UnixNano())

		req := []byte(`{
			"pan": "` + aggregation.Pan + `",
			"time": ` + strconv.Itoa(now) + `
		}`)

		fmt.Println(string(req))
		request := &processing.CompleteWithCalculateRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(completeCalc).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}
}
