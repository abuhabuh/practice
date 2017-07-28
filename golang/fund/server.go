package fund

import (
    "fmt"
)

type FundServer struct {
    Commands chan interface{}
    fund *Fund
}

func NewFundServer(initialBalance int) *FundServer {
    server := &FundServer{

        Commands: make(chan interface{}),
        fund: NewFund(initialBalance),
    }

    // Spawn off the server's main loop immediately
    go server.loop()
    return server
}

func (s *FundServer) loop() {

    for command := range s.Commands {

        switch command.(type) {
        case withDrawCommand:
            withdraw := command.(withDrawCommand)
            s.fund.Withdraw(withdraw.Amount)
        case balanceCommand:
            balance := command.(balanceCommand)
            respChan := balance.Response
            respChan <- s.fund.Balance()
        case transactCommand:
            break
        default:
            panic(fmt.Sprintf("damnit"))
        }

    }

}


func (s FundServer) Balance() int {
    retChan := make(chan int)
    s.Commands <- balanceCommand {Response: retChan}
    balance := <- retChan
    return balance
}

func (s FundServer) Withdraw(amount int) {
    s.Commands <- withDrawCommand {Amount: amount}
}


type transactCommand struct {
    Callback func(fund *Fund)
}
type withDrawCommand struct {
    Amount int
}
type balanceCommand struct {
    Response chan int
}
