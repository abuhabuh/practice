package fund

type Fund struct {
    // balance is unexported (private) because it's lowercase
    balance int
}


// constructor
func NewFund(initialBalance int) *Fund {
    return &Fund {
        balance: initialBalance,
    }
}


/*** Methods ***/
func (f *Fund) Balance() int {
    return f.balance
}

func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}
