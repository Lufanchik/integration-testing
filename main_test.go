package integration_testing

import (
	"errors"
	"fmt"
	"lab.siroccotechnology.ru/tp/integration-testing/passes"
	"lab.siroccotechnology.ru/tp/integration-testing/test"
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
	CasesParallel = append(CasesParallel, c)
}

func TestFull(t *testing.T) {
	tasks := make(chan test.Cases, Workers)
	err := make(chan error)
	done := make(chan struct{})
	t1 := time.Now()

	for i := 0; i < Workers; i++ {
		go func() {
			for v := range tasks {
				test.Run(t, v)
				if t.Failed() {
					err <- errors.New("failed")
				}
				done <- struct{}{}
			}
		}()
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

		select {
		case <-err:
			fmt.Println(err)
		case <-done:
			fmt.Println("done")
		}

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
	test.Run(t, passes.CasesCancel)
}
