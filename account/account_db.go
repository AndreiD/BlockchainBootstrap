package account

import "fmt"

// Account ...
type Account struct {
	Address string // address of the holder
	Balance int    // amount of the account
}

// Accounts is the database of accounts.
var Accounts []Account

// AddAccountWithBalance ads a new account to the database
func AddAccountWithBalance(addr string, bal int) {
	Accounts = append(Accounts, Account{addr, bal})
}

// ListAllAccount displays all accounts
func ListAllAccount() {
	fmt.Println("Accounts Balances Database")
	for _, account := range Accounts {
		fmt.Printf("%s %d\n", account.Address, account.Balance)
	}
	fmt.Println("")
}

// Transfer does a crypto transfer
func Transfer(from Account, to Account, amount int) error {
	if from.Balance < amount {
		return fmt.Errorf("you have insufficient balance for this transfer")
	}

	// decrement from
	from.Balance = from.Balance - amount

	// increment to
	to.Balance = to.Balance + amount

	// should return a tx hash here
	return nil
}
