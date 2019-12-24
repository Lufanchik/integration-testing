package integration_testing

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mtppk"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/face"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"net/http"
	"net/http/pprof"
	"testing"
	"time"
)

const Workers = 20

var (
	Cases         []test.Cases
	CasesParallel []test.Cases
)

func Add(c test.Cases) {
	Cases = append(Cases, c)
}

func AddP(c test.Cases) {
	nCases := test.Cases{}
	err := copier.Copy(&nCases, &c)
	if err != nil {
		panic(err)
	}
	CasesParallel = append(CasesParallel, nCases)
}

func status(listenAddr string) *http.Server {
	router := http.NewServeMux()
	router.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	router.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	router.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}
}

func TestFull(t *testing.T) {
	statusServer := status(":9096")

	go func() {
		if err := statusServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			require.NoError(t, err)
		}
	}()

	tasks := make(chan test.Cases, len(CasesParallel))
	err := make(chan error)
	done := make(chan struct{})
	t1 := time.Now()

	for i := 0; i < Workers; i++ {
		go func(tasks chan test.Cases, err chan error, done chan struct{}, w int) {
			for v := range tasks {
				test.Run(t, v, test.RequestTypeOnline)
				//test.Run(t, v, test.RequestTypeOffline)
				//test.Run(t, v, test.RequestType(gofakeit.Number(1, 2)))
				done <- struct{}{}
			}
		}(tasks, err, done, i)
	}

	for _, v := range CasesParallel {
		tasks <- v
	}
	close(tasks)
	scenarios := 0
	cases := 0
	steps := 0

	for _, v := range CasesParallel {
		scenarios++
		cases += len(v)
		for _, s := range v {
			steps += len(s.T)
		}

		<-done
	}

	for _, v := range Cases {
		scenarios++
		cases += len(v)
		for _, s := range v {
			steps += len(s.T)
		}
	}

	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println(fmt.Sprintf("scenarios: %d, cases: %d, steps: %d, time: %d minutes, %d seconds", scenarios, cases, steps, int(d.Minutes()), int(d.Seconds())))
}

func TestSimple(t *testing.T) {
	test.Run(t, webapi.CasesWEBAPI, test.RequestTypeOnline)
	//test.Run(t, mck.CasesMetroComplexMCK1, test.RequestTypeOffline)
	//test.Run(t, mck.CasesMetroComplexMCK2, test.RequestTypeOffline)
	//test.Run(t, mck.CasesMetroComplexMCK3, test.RequestTypeOffline)
	//test.Run(t, mck.CasesMetroComplexMCK4, test.RequestTypeOffline)
	//test.Run(t, mck.CasesComplexTimeMCK, test.RequestTypeOffline)
	//test.Run(t, mck.CasesOfflineMetroComplexMCK, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesComplexMCDMOPartOne, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesComplexMCDMOPartTwo, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesComplexMCDMOPartThree, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesComplexMCDMOPartFour, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesComplexMCDMOPartFife, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesMetroComplexMCDMSK1, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesMetroComplexMCDMSK2, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesMetroComplexMCDMSK3, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesMetroComplexMCDMSK4, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesMetroComplexMCDMSK5, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesComplexTimeMCD, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesOfflineMetroComplexMCD, test.RequestTypeOffline)
	//test.Run(t, mcd.CasesOfflineMetroComplexMCDMO, test.RequestTypeOffline)
	//test.Run(t, mgt.CasesMGT, test.RequestTypeOffline)
	//test.Run(t, mgt.CasesMGT2, test.RequestTypeOffline)
	//test.Run(t, mgt.CasesMGT3, test.RequestTypeOffline)
	//test.Run(t, mm.CasesMetroComplexMM1, test.RequestTypeOffline)
	//test.Run(t, mm.CasesMetroComplexMM2, test.RequestTypeOffline)
	//test.Run(t, mm.CasesMetroComplexMM3, test.RequestTypeOffline)
	//test.Run(t, mm.CasesMetroComplexMM4, test.RequestTypeOffline)
	//test.Run(t, mm.CasesComplexTimeMM, test.RequestTypeOffline)
	//test.Run(t, mm.CasesOfflineMetroComplexMM, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesComplexPassMMTS1, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesComplexPassMMTS2, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesComplexPassMMTS3, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesComplexPassMMTS4, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesComplexPassMMTS5, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesComplexTimeMMTS, test.RequestTypeOffline)
	//test.Run(t, mmts.CasesOfflineMetroComplexMMTS, test.RequestTypeOffline)
	//test.Run(t, mtppk.CasesMTPPKPasses, test.RequestTypeOffline)
	//test.Run(t, passes.CasesCancel, test.RequestTypeOffline)
	//test.Run(t, passes.CasesWrongTimeComplexPass, test.RequestTypeOffline)
	//test.Run(t, passes.CasesScopeCheckPass, test.RequestTypeOffline)
	//test.Run(t, passes.CasesSimpleComplexPass, test.RequestTypeOffline)
	//test.Run(t, mtppk.CasesMTPPKPasses, test.RequestTypeOffline)
	test.Run(t, face.CasesAuthWithFace, test.RequestTypeOffline)
	//test.Run(t, mm.CasesMetroComplexMM1, test.RequestTypeOffline)
}
