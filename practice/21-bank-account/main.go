package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance int
}

func (a *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("Сумма пополнения должна быть положительной")
	}
	a.Balance += amount
	return nil
}

func (a *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("Сумма снятия должна быть положительной")
	}
	if amount > a.Balance {
		return errors.New("Недостаточно средств")
	}
	a.Balance -= amount
	return nil
}

func main() {
	var owner string
	var balance int
	var deposit, withdraw int
	fmt.Scan(&owner, &balance, &deposit, &withdraw)

	account := BankAccount{
		Owner:   owner,
		Balance: balance,
	}

	fmt.Println("Владелец: ", owner)
	fmt.Println("Начальный баланс: ", balance)

	err := account.Deposit(deposit)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	afterDeposit := account.Balance

	err = account.Withdraw(withdraw)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("После пополнения: ", afterDeposit)
	fmt.Println("После снятия: ", account.Balance)
}
