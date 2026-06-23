package main

import atm "ATM/header"

func main() {
    account := atm.NewATM("Alice", 1000, 200)
    account.Display()
}