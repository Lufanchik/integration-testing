package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"lab.dt.multicarta.ru/tp/common/crypt/sha"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"testing"
	"time"

	passService "lab.dt.multicarta.ru/tp/pass-service/proto"
)

var (
	httpProcessingApi     *httpexpect.Expect
	httpApmApi            *httpexpect.Expect
	httpWebApi            *httpexpect.Expect
	httpCardService       *httpexpect.Expect
	httpCalculatorService *httpexpect.Expect
	httpCommentService    *httpexpect.Expect
	httpAuthService       *httpexpect.Expect
	httpReviseService     *httpexpect.Expect
	httpResolveService    *httpexpect.Expect
	httpTWPGService       *httpexpect.Expect
	ps                    passService.PassService
)

func NanoToMicro(tm uint64) uint64 {
	return tm / uint64(time.Microsecond) * 1000
}

func ConfigurePass(t *testing.T, p *Pass, carrierID string, card *processing.Card) {
	if p.Now != nil {
		Now = p.Now
	} else {
		Now = NowBackup
	}
	GenerateEmv(card, p)
	p.carrierID = carrierID
}

func RunPass(t *testing.T, p *Pass, scenario *Case, carrierID string, card *processing.Card) (PassResponser, PassResponser) {
	var isEqual bool
	if p.Equal != 0 {
		isEqual = true
		eq, ok := (scenario.T[p.Equal-1]).(*Pass)
		if !ok {
			t.Fail()
		}
		*p = *eq
		card = p.card
		p.AuthType = AuthTypeCorrect
		p.EmptyEMV = false
		carrierID = p.carrierID
	}

	ConfigurePass(t, p, carrierID, card)
	tapReq, tapResp := TapBySubCarrier(t, p, card)
	if p.RevisePass != 0 {
		prp, ok := (scenario.T[p.RevisePass-1]).(*ProcessRevisePass)
		if !ok {
			t.Fail()
		}
		tapReq.Tap.Card.Pan = prp.req.Request.Tap.Card.Pan
		tapReq.Tap.Created = prp.req.Request.Tap.Created
		//tapReq.Tap.Terminal = prp.req.Request.Tap.Terminal
		tapReq2 := &processing.TapRequest{
			Id:      prp.req.Request.Id,
			Created: prp.req.Request.Created,
			Tap:     tapReq.Tap,
		}
		tapReq = tapReq2
		tapResp.Id = prp.pass.Id
		p.tapRequest = tapReq
		fmt.Printf("p: %+v", *p)
	}
	timeRequest, resp := PassBySubCarrier(t, tapReq, p)
	if p.RevisePass != 0 {
		timeRequest = p.tapRequest.Created
	}
	require.Equal(t, tapResp.GetId(), resp.GetId())
	var parent, ingress, aggregate, secondParent *Pass
	if p.Parent > 0 {
		pr, ok := (scenario.T[p.Parent-1]).(*Pass)
		if !ok {
			t.Fail()
		}
		parent = pr
		parent.isParent = true
		parent.timeToWait = parent.TimeToWait
	}

	if p.SecondParent > 0 {
		pr, ok := (scenario.T[p.SecondParent-1]).(*Pass)
		if !ok {
			t.Fail()
		}
		secondParent = pr
	}

	if p.Ingress > 0 {
		ing, ok := (scenario.T[p.Ingress-1]).(*Pass)
		if !ok {
			t.Fail()
		}
		ingress = ing
		ingress.timeToWait = ingress.TimeToWait
	}
	if p.Aggregate > 0 {
		agg, ok := (scenario.T[p.Aggregate-1]).(*Pass)
		if !ok {
			t.Fail()
		}
		aggregate = agg
	}

	p.tapRequest = tapReq
	if !isEqual {
		p.timeRequest = timeRequest
	}
	fmt.Println(p.timeRequest)
	fmt.Println(p.id)
	p.card = card
	p.parent = parent
	p.secondParent = secondParent
	p.ingress = ingress
	p.aggregate = aggregate

	//time.Sleep(TimeAfterRequest)
	ValidatePass(t, p, p.parent, p.ingress, true)
	//if !isAggregate(p) {
	AuthStatus(t, p)

	//}

	if parent != nil {
		ValidatePass(t, parent, parent.parent, parent.ingress, false)
		//if !isAggregate(parent) {
		AuthStatus(t, parent)
		//}
	}

	if ingress != nil {
		ValidatePass(t, ingress, ingress.parent, ingress.ingress, false)
		//if !isAggregate(ingress) {
		AuthStatus(t, ingress)
		//}
	}
	return tapResp, resp
}
func getRequestType(t *testing.T, p *Pass) RequestType {
	if GlobalRequestType != RequestTypeNone {
		return GlobalRequestType
	}

	if p.RequestType == RequestTypeNone {
		return RequestType(gofakeit.Number(1, 2))
	}

	return p.RequestType
}

//ApiRequest simple api request
func RunApiRequest(t *testing.T, cases Cases, rt RequestType) {
	httpProcessingApi = httpexpect.New(t, ProcessingApiUrl)
	httpApmApi = httpexpect.New(t, ApmApiUrl)
	httpWebApi = httpexpect.New(t, WebApiUrl)
	httpCommentService = httpexpect.New(t, CommentsURL)
	httpCardService = httpexpect.New(t, CardURL)
	httpCalculatorService = httpexpect.New(t, CalculatorApiUrl)
	httpAuthService = httpexpect.New(t, AuthServiceUrl)
	httpReviseService = httpexpect.New(t, ReviseApiUrl)
	httpResolveService = httpexpect.New(t, ResolveApiUrl)
	httpTWPGService = httpexpect.New(t, TWPGApiUrl)

	fmt.Printf("%v: %s\n", httpTWPGService, TWPGApiUrl)

	for _, v := range cases {
		fmt.Println("name: " + v.N)
		t.Run("name: "+v.N, func(t *testing.T) {
			for _, step := range v.T {
				//ReviseRegistry
				res, ok := step.(*Resolve)
				if ok {
					ResolveTestApi(t, res)
				}

				//ReviseRegistry
				rev, ok := step.(*Revise)
				if ok {
					ReviseTestApi(t, rev)
				}

				//Login
				lg, ok := step.(*Login)
				if ok {
					LoginApi(t, lg)
				}

				commentsCRUD, ok := step.(*CommentsCRUD)
				if ok {
					CommentsCheck(t, commentsCRUD)
				}
			}
		})
	}
}

func Run(t *testing.T, cases Cases, rt RequestType) {
	httpProcessingApi = httpexpect.New(t, ProcessingApiUrl)
	httpApmApi = httpexpect.New(t, ApmApiUrl)
	httpWebApi = httpexpect.New(t, WebApiUrl)
	httpCommentService = httpexpect.New(t, CommentsURL)
	httpCardService = httpexpect.New(t, CardURL)
	httpCalculatorService = httpexpect.New(t, CalculatorApiUrl)
	httpAuthService = httpexpect.New(t, AuthServiceUrl)
	httpReviseService = httpexpect.New(t, ReviseApiUrl)
	httpResolveService = httpexpect.New(t, ResolveApiUrl)
	httpTWPGService = httpexpect.New(t, TWPGApiUrl)

	type NCase struct {
		c    *Case
		card *processing.Card
		//carrierId string
		carrierIds []string
		shaPan     string
	}
	nc := make([]*NCase, len(cases))

	for k, _ := range cases {
		nc[k] = &NCase{
			c: &cases[k],
		}
		ids := make([]string, len(nc[k].c.T))
		for numInCase := range nc[k].c.T {
			ids[numInCase] = uuid.New().String()
		}
		nc[k].carrierIds = ids
		if cases[k].CustomerId != "" {
			nc[k].card = FaceCard(cases[k].CardSystem, cases[k].CustomerId)
		} else if cases[k].PassType == pass.PassType_PASS_MT {
			nc[k].card = MTCard(cases[k].CardSystem)
		} else {
			if nc[k].c.Card != nil {
				nc[k].card = nc[k].c.Card
			} else {
				nc[k].card = Card(cases[k].CardSystem)
			}
		}
		pan, err := sha.Generate(nc[k].card.Pan)
		if err != nil {
			panic(err)
		}
		nc[k].shaPan = pan
	}

	results := make([][]*Pass, len(nc))

	for casesNum, ncc := range nc {
		fmt.Println("name: " + ncc.c.N)
		scenario := ncc.c
		for N, step := range scenario.T {
			fmt.Println(N + 1)
			t.Run("case: "+scenario.N, func(t *testing.T) {
				//Pass
				firstPass, ok := step.(*Pass)
				if ok {
					firstPass.RequestType = rt
					firstPass.faceId = ncc.c.CustomerId
					//Если PassType в начале кейса не указан, дефолтим PassType_PASS_BBK, иначе используется предустановленный
					if ncc.c.PassType == pass.PassType_PASS_NONE {
						//Если фейс айди заполнен, то вместо BBK используем PassType_FACE_ID
						if ncc.c.CustomerId != "" {
							firstPass.PassType = pass.PassType_PASS_FACE_ID
						} else {
							firstPass.PassType = pass.PassType_PASS_BBK
						}
					} else if ncc.c.CustomerId != "" {
						firstPass.PassType = pass.PassType_PASS_FACE_ID
					} else {
						firstPass.PassType = ncc.c.PassType
					}
					if firstPass.PaymentType == PaymentTypePrepayed {
						firstPass.PassType = pass.PassType_PASS_MT
					}

					if firstPass.PassType == pass.PassType_PASS_MT {
						firstPass.PaymentType = PaymentTypePrepayed
					}

					if firstPass.PassType == pass.PassType_PASS_FACE_ID {
						ncc.card.Pan = ncc.shaPan
					}

					fmt.Printf("name: %s; pass-type: %d\n", ncc.c.N, ncc.c.PassType)
					fmt.Printf("name: %+v\n", *ncc.card)
					tpp, ppp := RunPass(t, firstPass, scenario, ncc.carrierIds[N], ncc.card)
					firstPass.tapResponse = tpp
					firstPass.passResponse = ppp
				}

				if results[casesNum] == nil {
					results[casesNum] = make([]*Pass, len(scenario.T))
				}

				results[casesNum][N] = firstPass

				//Updater
				u, ok := step.(Updater)
				if ok {
					pu, ok := (scenario.T[u.target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					Update(t, pu, u)
				}

				//PassCheck
				pc, ok := step.(*PassCheck)
				if ok {
					target, ok := (scenario.T[pc.Target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					parent, ok := (scenario.T[pc.Parent-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					PassCheckApi(t, pc, target, parent)
				}

				//Cancel
				cl, ok := step.(*Cancel)
				if ok {
					target, ok := (scenario.T[cl.Target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					target.isCancel = true
					CancelApi(t, cl, target)
				}

				//Parking
				pr, ok := step.(*Parking)
				if ok {
					ParkingApi(t, ncc.card, pr)
				}

				cm_calc, ok := step.(*CompleteWithCalculate)
				if ok {
					var aggregatePasses []*Pass
					for _, v := range scenario.T {
						ps, ok := v.(*Pass)
						if ok {
							aggregatePasses = append(aggregatePasses, ps)
						}
					}

					if len(aggregatePasses) != 0 {
						CompleteCalcApi(t, aggregatePasses, cm_calc.Pan)
					}
				}

				mcd_restore, ok := step.(*McdRestore)
				if ok {
					mcd_restore.Pan = "500358C1BB426F061A46743752B01BD1D1F930F7865AA439BFFA49CD11CB1F50"
					McdRestoreApi(t, mcd_restore)
				}

				cm, ok := step.(*Complete)
				if ok {
					start, ok := (scenario.T[cm.StartPass-1]).(*Pass)
					if !ok {
						t.Fail()
					}

					start.isComplete = true
					start.completeSum = uint32(cm.Sum)
					var passes []*Pass
					for _, v := range cm.Passes {
						ps := (scenario.T[v-1]).(*Pass)
						ps.isComplete = true
						ps.completeSum = uint32(cm.Sum)
						passes = append(passes, ps)
					}

					CompleteApi(t, start, passes, cm.Sum)
				}

				fcl, ok := step.(*RegisterFaceId)
				if ok {
					FaceApiGetRegisterLink(t, ncc.card, fcl)
				}

				faceCheck, ok := step.(*FaceIdRegistrationStatus)
				if ok {
					faceCheck.Id = ncc.card.Pan
					FaceApiCheckStatus(t, faceCheck)
				}

				twpgPay, ok := step.(*TWPGCreateAndPayOrderStep)
				if ok {
					twpgPay.CustomerId = ncc.card.Pan
					TWPGCreateAndPayOrder(t, twpgPay)
					ncc.c.TWPGOrderId = twpgPay.OrderId
				}

				twpgOrderStatus, ok := step.(*TWPGOrderStatus)
				if ok {
					twpgOrderStatus.OrderId = ncc.c.TWPGOrderId
					TWPGCheckOrderStatus(t, twpgOrderStatus)
				}

				twpgReverse, ok := step.(*TWPGReverseOrder)
				if ok {
					twpgReverse.OrderId = ncc.c.TWPGOrderId
					TWPGReverse(t, twpgReverse)
				}

				cgf, ok := step.(*CardGetFull)
				if ok {
					CardApiGetFull(t, cgf)
				}

				prp, ok := step.(*ProcessRevisePass)
				if ok {
					ProcessRevisePassRequest(t, prp)
				}

				rc, ok := step.(*ReaderConfiguration)
				if ok {
					ReaderConfigurationSend(t, rc)
				}

				wgw, ok := step.(*WebAPIPasses)
				if ok {
					var passes []*Pass
					for _, v := range wgw.Passes {
						passes = append(passes, (scenario.T[v-1]).(*Pass))
					}

					WebAPI(t, ncc.card, passes)
				}

				csl, ok := step.(*CardStopList)
				if ok {
					var reauthPass *Pass
					for _, v := range csl.Passes {
						p := (scenario.T[v-1]).(*Pass)
						if p.AuthType == AuthTypeUnsuccessWithReauth {
							reauthPass = p
							break
						}
					}
					require.NotNil(t, reauthPass)

					csl.PassId = reauthPass.id
					csl.Pan = ncc.card.Pan
					CardCheckStopList(t, csl)
				}

				fra, ok := step.(*ForceReauth)
				if ok {
					var reauthPass *Pass
					for _, v := range fra.Passes {
						p := (scenario.T[v-1]).(*Pass)
						if p.AuthType == AuthTypeUnsuccessWithReauth {
							reauthPass = p
							break
						}
					}
					require.NotNil(t, reauthPass)

					fra.PassId = reauthPass.id
					ForceReauthCall(t, fra)
				}

				pra, ok := step.(*ReAuth)
				if ok {
					ReAuthGo(t, pra)
				}

				ara, ok := step.(*AuthResponse)
				if ok {
					AuthResponseGo(t, ara)
				}
			})
			if t.Failed() {
				break
			}
		}
		if t.Failed() {
			break
		}
	}

	if !t.Failed() {
		time.Sleep(time.Second * 10)
		for casesNum, ncc := range nc {
			scenario := ncc.c
			if scenario.SkipIdempotencyCheck {
				break
			}

			fmt.Println("name check 1: " + ncc.c.N)

			for N, step := range scenario.T {
				t.Run("case check 1: "+scenario.N, func(t *testing.T) {
					//Pass
					secondPass, ok := step.(*Pass)
					if ok && !secondPass.isCancel && secondPass.faceId == "" {
						oldPass := results[casesNum][N]
						secondPass.tapRequest = oldPass.tapRequest
						fmt.Println(fmt.Sprintf("check 1 = %d", N+1))
						ConfigurePass(t, secondPass, ncc.carrierIds[N], ncc.card)
						ValidatePass(t, secondPass, secondPass.parent, secondPass.ingress, false)
						if !isAggregate(secondPass) && !isAggeregateByCardSystemCarrier(secondPass.Carrier, secondPass.card.System) {
							AuthStatus(t, secondPass)
						}
						secondPass.Terminal = secondPass.tapRequest.Tap.Terminal

						_, respTap := TapBySubCarrier(t, secondPass, ncc.card)
						_, respPass := PassBySubCarrier(t, secondPass.tapRequest, secondPass)
						require.Equal(t, respTap.GetId(), respPass.GetId())

						require.Equal(t, respTap.GetId(), oldPass.tapResponse.GetId())
						require.Equal(t, respTap.GetId(), oldPass.passResponse.GetId())
						require.Equal(t, secondPass.id, oldPass.id)
						if !ncc.c.NotDoubleCheck {
							ValidatePass(t, secondPass, secondPass.parent, secondPass.ingress, false)
						}

						if !isAggregate(secondPass) && !isAggeregateByCardSystemCarrier(secondPass.Carrier, secondPass.card.System) {
							AuthStatus(t, secondPass)
						}
					}

				})

				if t.Failed() {
					break
				}
			}

			if t.Failed() {
				break
			}
		}
	}

	if !t.Failed() {
		time.Sleep(time.Second * 10)

		for casesNum, ncc := range nc {
			scenario := ncc.c

			if scenario.SkipIdempotencyCheck {
				break
			}

			fmt.Println("name check 2: " + ncc.c.N)
			for N, step := range scenario.T {
				t.Run("case check 2: "+scenario.N, func(t *testing.T) {
					//Pass
					thirdPass, ok := step.(*Pass)
					if ok && !thirdPass.isCancel && thirdPass.faceId == "" {
						fmt.Println(fmt.Sprintf("check 2 = %d", N+1))
						ConfigurePass(t, thirdPass, ncc.carrierIds[N], ncc.card)
						if !ncc.c.NotDoubleCheck {
							ValidatePass(t, thirdPass, thirdPass.parent, thirdPass.ingress, false)
						}
						if !isAggregate(thirdPass) && !isAggeregateByCardSystemCarrier(thirdPass.Carrier, thirdPass.card.System) {
							AuthStatus(t, thirdPass)
						}
						_, respTap := TapBySubCarrier(t, thirdPass, ncc.card)
						_, respPass := PassBySubCarrier(t, thirdPass.tapRequest, thirdPass)
						require.Equal(t, respTap.GetId(), respPass.GetId())
						oldPass := results[casesNum][N]
						require.Equal(t, respTap.GetId(), oldPass.tapResponse.GetId())
						require.Equal(t, respTap.GetId(), oldPass.passResponse.GetId())
						if !ncc.c.NotDoubleCheck {
							ValidatePass(t, thirdPass, thirdPass.parent, thirdPass.ingress, false)
						}
						if !isAggregate(thirdPass) && !isAggeregateByCardSystemCarrier(thirdPass.Carrier, thirdPass.card.System) {
							AuthStatus(t, thirdPass)
						}
					}
				})
				if t.Failed() {
					break
				}
			}
			if t.Failed() {
				break
			}
		}
	}
}
