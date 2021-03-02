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
	calcUrl    = "http://localhost:9098"
	restoreMCD = "/twirp/sirocco.CalculatorAPI/ProcessMcdRestore"
)

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}

func Test_1_1(t *testing.T) {
	response := &processing.McdRestoreResponse{}
	{
		httpService := httpexpect.New(t, calcUrl)
		now := int(time.Now().UnixNano())

		req := []byte(`{
			"pan": "` + aggregation.Pan + `",
			"time": ` + strconv.Itoa(now) + `
		}`)

		fmt.Println(string(req))
		request := &processing.McdRestoreRequest{}
		err := json.Unmarshal(req, request)
		require.NoError(t, err)

		r := httpService.POST(restoreMCD).WithJSON(request).
			Expect().
			Status(http.StatusOK)

		logRequest(r)

		err = jsonpb.Unmarshal(strings.NewReader(r.Body().Raw()), response)
		require.NoError(t, err)

		fmt.Println(response)
	}
}
