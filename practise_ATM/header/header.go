package header

import "fmt"
type ATM struct{
	name string
	deposit int
	withdraw int
}

type Person interface{
	getName() string
	transaction_list() string
	getBalance() int
}

func (atm *ATM) transaction_list() string {
	return fmt.Sprintf(
		"Deposited: %d, Withdrawn: %d",
		atm.deposit,
		atm.withdraw,
	)
}

func (atm *ATM) getName() string{
	return atm.name
}

func (atm *ATM) getBalance() int{
	return atm.deposit - atm.withdraw
}

func NewATM(name string, deposit, withdraw int) *ATM {
	return &ATM{
		name:     name,
		deposit:  deposit,
		withdraw: withdraw,
	}
}

func (atm *ATM) Display() {
	fmt.Println("ATM Details")
	fmt.Println("Name:", atm.getName())
	fmt.Println("Balance:", atm.getBalance())
	fmt.Println(atm.transaction_list())
}