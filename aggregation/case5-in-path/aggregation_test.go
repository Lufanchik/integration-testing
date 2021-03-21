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

	onlinePassMM   = "/mm/twirp/sirocco.ProcessingAPI/ProcessOnlinePass"
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

// 1 В пути + выход онлайн с агрегацией в мск
// 2 В пути + выход онлайн с без агрегации в мск
// 3 В пути + выход онлайн с агрегацией в мо
// 4 В пути + выход онлайн с без агрегации в мо
// 5 В пути + выход оффлайн с агрегацией в мск
// 6 В пути + выход оффлайн с без агрегации в мск
// 7 В пути + выход оффлайн с агрегацией в мо
// 8 В пути + выход оффлайн с без агрегации в мо
// 9 вход онлайн с агрегацией в мск + в пути
// 10 вход онлайн с без агрегации в мск + в пути
// 11 вход онлайн с агрегацией в мо + в пути
// 12 вход онлайн с без агрегации в мо + в пути
// 13 вход оффлайн с агрегацией в мск + в пути
// 14 вход оффлайн с без агрегации в мск + в пути
// 15 вход оффлайн с агрегацией в мо + в пути
// 16 вход оффлайн с без агрегации в мо + в пути

func Test_1(t *testing.T) {
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
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
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

		Request(t, reqMCD1, onlinePassMCD)
		Request(t, reqMCD2, onlinePassMCD)
	}
}

func Test_2(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "2"

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
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
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

		Request(t, reqMCD1, onlinePassMCD)
		Request(t, reqMCD2, onlinePassMCD)
	}
}

func Test_3(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "3"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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

		Request(t, reqMCD1, onlinePassMCD)
		Request(t, reqMCD2, onlinePassMCD)
	}
}

func Test_4(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "4"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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

		Request(t, reqMCD1, onlinePassMCD)
		Request(t, reqMCD2, onlinePassMCD)
	}
}

func Test_5(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "5"

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
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
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

		Request(t, reqMCD1, offlinePassMCD)
		Request(t, reqMCD2, offlinePassMCD)
	}
}

func Test_6(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "6"

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
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
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

		Request(t, reqMCD1, offlinePassMCD)
		Request(t, reqMCD2, offlinePassMCD)
	}
}

func Test_7(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "7"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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

		Request(t, reqMCD1, offlinePassMCD)
		Request(t, reqMCD2, offlinePassMCD)
	}
}

func Test_8(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "8"

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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

		Request(t, reqMCD1, offlinePassMCD)
		Request(t, reqMCD2, offlinePassMCD)
	}
}

func Test_9(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "9"

		reqMCD2 := []byte(`{
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

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
						"direction": 4,
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

		Request(t, reqMCD2, onlinePassMCD)
		Request(t, reqMCD1, onlinePassMCD)
	}
}

func Test_10(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "10"

		reqMCD2 := []byte(`{
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
						"direction": 4,
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

		Request(t, reqMCD2, onlinePassMCD)
		Request(t, reqMCD1, onlinePassMCD)
	}
}

func Test_11(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "11"

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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

		Request(t, reqMCD2, onlinePassMCD)
		Request(t, reqMCD1, onlinePassMCD)
	}
}

func Test_12(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "12"

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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

		Request(t, reqMCD2, onlinePassMCD)
		Request(t, reqMCD1, onlinePassMCD)
	}
}

func Test_13(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "13"

		reqMCD2 := []byte(`{
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

		reqMCD1 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
						"direction": 4,
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

		Request(t, reqMCD2, offlinePassMCD)
		Request(t, reqMCD1, offlinePassMCD)
	}
}

func Test_14(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "14"

		reqMCD2 := []byte(`{
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2000285",
						"direction": 4,
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

		Request(t, reqMCD2, offlinePassMCD)
		Request(t, reqMCD1, offlinePassMCD)
	}
}

func Test_15(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "15"

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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

		Request(t, reqMCD2, offlinePassMCD)
		Request(t, reqMCD1, offlinePassMCD)
	}
}

func Test_16(t *testing.T) {
	{
		now := time.Now()
		pan := aggregation.Pan + "16"

		reqMCD2 := []byte(`{
			"id": "` + uuid.New().String() + `",
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*6).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
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
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
				"tap": {
				"created": ` + strconv.Itoa(int(now.Add(time.Minute*7).UnixNano())) + `,
					"resolution": 1,
					"sign": "test",
					"terminal": {
						"id": "3213",
						"station": "2002952",
						"direction": 4,
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

		Request(t, reqMCD2, offlinePassMCD)
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
