package test

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect"
	"github.com/google/uuid"
	"lab.siroccotechnology.ru/tp/common/messages/processing"
	"testing"
	"time"

	passService "lab.siroccotechnology.ru/tp/pass-service/proto"
)

var (
	httpProcessingApi *httpexpect.Expect
	httpApmApi        *httpexpect.Expect
	httpWebApi        *httpexpect.Expect
	ps                passService.PassService
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

func RunPass(t *testing.T, p *Pass, scenario *Case, carrierID string, card *processing.Card) {
	ConfigurePass(t, p, carrierID, card)
	tapReq := TapBySubCarrier(t, p, card)
	timeRequest := PassBySubCarrier(t, tapReq, p)
	var parent, ingress, aggregate *Pass
	if p.Parent > 0 {
		pr, ok := (scenario.T[p.Parent-1]).(*Pass)
		if !ok {
			t.Fail()
		}
		parent = pr
		parent.isParent = true
		parent.timeToWait = parent.TimeToWait
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
	p.timeRequest = timeRequest
	p.card = card
	p.parent = parent
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

func Run(t *testing.T, cases Cases, rt RequestType) {
	httpProcessingApi = httpexpect.New(t, ProcessingApiUrl)
	httpApmApi = httpexpect.New(t, ApmApiUrl)
	httpWebApi = httpexpect.New(t, WebApiUrl)
	type NCase struct {
		c         *Case
		card      *processing.Card
		carrierId string
	}
	nc := make([]*NCase, len(cases))

	for k, _ := range cases {
		nc[k] = &NCase{
			c:         &cases[k],
			card:      Card(cases[k].CardSystem),
			carrierId: uuid.New().String(),
		}
	}

	for _, ncc := range nc {
		fmt.Println("name: " + ncc.c.N)
		fmt.Println(ncc.card.String())
		scenario := ncc.c
		for N, step := range scenario.T {
			fmt.Println(N + 1)
			t.Run("case: "+scenario.N, func(t *testing.T) {
				//Pass
				p, ok := step.(*Pass)
				if ok {
					p.RequestType = rt
					RunPass(t, p, scenario, ncc.carrierId, ncc.card)
				}

				//Updater
				u, ok := step.(Updater)
				if ok {
					pu, ok := (scenario.T[u.target-1]).(*Pass)
					if !ok {
						t.Fail()
					}
					Update(t, pu, u)
				}

				//AbsGetRegistry
				agr, ok := step.(*AbsGetRegistry)
				if ok {
					AbsGetRegistryApi(t, agr)
				}

				//AbsGetRegistry
				lg, ok := step.(*Login)
				if ok {
					LoginApi(t, lg)
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

				wgw, ok := step.(*WebAPIPasses)
				if ok {
					var passes []*Pass
					for _, v := range wgw.Passes {
						passes = append(passes, (scenario.T[v-1]).(*Pass))
					}

					WebAPI(t, ncc.card, passes)
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
		for _, ncc := range nc {
			fmt.Println("name check 1: " + ncc.c.N)
			fmt.Println(ncc.card.String())
			scenario := ncc.c
			for N, step := range scenario.T {
				t.Run("case check 1: "+scenario.N, func(t *testing.T) {
					//Pass
					p, ok := step.(*Pass)
					if ok && !p.isCancel {
						fmt.Println(fmt.Sprintf("check 1 = %d", N+1))
						ConfigurePass(t, p, ncc.carrierId, ncc.card)
						ValidatePass(t, p, p.parent, p.ingress, false)
						//if !isAggregate(p) {
						AuthStatus(t, p)
						//}
						TapBySubCarrier(t, p, ncc.card)
						PassBySubCarrier(t, p.tapRequest, p)
						ValidatePass(t, p, p.parent, p.ingress, false)
						//if !isAggregate(p) {
						AuthStatus(t, p)
						//}
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
		for _, ncc := range nc {
			fmt.Println("name check 2: " + ncc.c.N)
			fmt.Println(ncc.card.String())
			scenario := ncc.c
			for N, step := range scenario.T {
				t.Run("case check 2: "+scenario.N, func(t *testing.T) {
					//Pass
					p, ok := step.(*Pass)
					if ok && !p.isCancel {
						fmt.Println(fmt.Sprintf("check 2 = %d", N+1))
						ConfigurePass(t, p, ncc.carrierId, ncc.card)
						ValidatePass(t, p, p.parent, p.ingress, false)
						//if !isAggregate(p) {
						AuthStatus(t, p)
						//}
						TapBySubCarrier(t, p, ncc.card)
						PassBySubCarrier(t, p.tapRequest, p)
						ValidatePass(t, p, p.parent, p.ingress, false)
						//if !isAggregate(p) {
						AuthStatus(t, p)
						//}
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
