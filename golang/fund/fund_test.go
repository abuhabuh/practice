// https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial

package fund

import (
    "sync"
    "testing"
)

const WORKERS = 10

func BenchmarkWithdrawals(b *testing.B) {
    // skip N=1
    if b.N < WORKERS {
        return
    }

    // Add as many dollars as we have iterations this run
    fundServer := NewFundServer(b.N)


    dollarsPerFounder := b.N / WORKERS


    // var wg sync.WaitGroup
    wg := sync.WaitGroup{}
    wgp := &wg

    for i := 0; i < WORKERS; i++ {

        (*wgp).Add(1)


        go func() {

            defer (*wgp).Done()

            for i := 0; i < dollarsPerFounder; i++ {

                if fundServer.Balance() <= 10 {
                    break
                }

                fundServer.Withdraw(1)
            }

        }()

    }

    (*wgp).Wait()

    balance := fundServer.Balance()
    if balance != 10 {
        b.Error("Balance was not zero: ", balance)
    }
}
