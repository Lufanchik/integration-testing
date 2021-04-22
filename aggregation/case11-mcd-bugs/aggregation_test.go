package case3_ingress

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
	offlinePassMM  = "/mm/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
	onlinePassMCD  = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
	offlinePassMCD = "/mcd/twirp/sirocco.ProcessingAPI/ProcessOfflinePass"
)

func logRequest(r *httpexpect.Response) {
	trace := r.Header("X-Trace-ID")
	fmt.Println(fmt.Sprintf("trace id: %s", trace.Raw()))
}

// мск 2000285
// мо 2002952
// виза 2
// мир 4

// 1 оффлайн, мир, мм, мцд вход, выход, авто, мцд выход, мм
// 2 оффлайн, мир, мцд вход, мцд валидатор (время запроса больше, время перевозчика меньше), мцд валидатор (время запроса меньше, время перевозчика больше)
// 3 оффлайн, мир, мм, мцд вход
// 4 оффлайн, виза, мцд вход, мцд выход, два прохода по валидатору
// 5 онлайн, виза, мцд вход, мцд выход, два прохода по валидатору
// 6 оффлайн, мир, мцд вход, мцд выход, два прохода по валидатору
// 7 онлайн, мир, мцд вход, мцд выход, два прохода по валидатору

func Test_1(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "1"

		reqMM1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*1).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*1).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*3).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*3).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 2,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD3 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*4).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*4).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 2,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMM2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		Request(t, reqMM1, offlinePassMM)
		time.Sleep(time.Second)
		Request(t, reqMCD1, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD2, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD3, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMM2, offlinePassMM)
	}
}

func Test_2(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "2"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*3).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*4).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD3 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*4).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*3).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		Request(t, reqMCD1, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD2, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD3, offlinePassMCD)
	}
}

func Test_3(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "3"

		reqMM1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*1).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*1).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMM2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		Request(t, reqMM1, offlinePassMM)
		time.Sleep(time.Second)
		Request(t, reqMCD1, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMM2, offlinePassMM)
	}
}

func Test_4(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "4"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
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

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
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

		reqMCD3 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
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

		reqMCD4 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
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

		Request(t, reqMCD1, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD2, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD3, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD4, offlinePassMCD)
	}
}

func Test_5(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "5"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
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

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
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

		reqMCD3 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
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

		reqMCD4 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
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
		time.Sleep(time.Second)
		Request(t, reqMCD2, onlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD3, onlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD4, onlinePassMCD)
	}
}

func Test_6(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "6"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 2,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD3 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD4 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		Request(t, reqMCD1, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD2, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD3, offlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD4, offlinePassMCD)
	}
}

func Test_7(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "7"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*2).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 1,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*5).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 2,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD3 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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

		reqMCD4 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2001220",
						"direction": 3,
						"sub_carrier": 1
					},
				"card": {
						"system": 4,
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
		time.Sleep(time.Second)
		Request(t, reqMCD2, onlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD3, onlinePassMCD)
		time.Sleep(time.Second)
		Request(t, reqMCD4, onlinePassMCD)
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
