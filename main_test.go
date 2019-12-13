package integration_testing

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
	"lab.siroccotechnology.ru/tp/integration-testing/webapi"
	"net/http"
	"net/http/pprof"
	"testing"
	"time"
)

const Workers = 18

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
				test.Run(t, v)
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
		test.Run(t, v)
	}

	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println(fmt.Sprintf("scenarios: %d, cases: %d, steps: %d, time: %d minutes, %d seconds", scenarios, cases, steps, int(d.Minutes()), int(d.Seconds())))
}

func TestSimple(t *testing.T) {
    test.Run(t, mgt.CasesMGT1)
	test.Run(t, mgt.CasesMGT2)
	test.Run(t, mgt.CasesMGT3)
}
