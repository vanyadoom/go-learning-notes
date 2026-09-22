package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("Сумма должна быть положительной")
	}
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("Сумма должна быть положительной")
	} else if amount > b.Balance {
		return fmt.Errorf("Недостаточно средств: требуется %d, доступно %d", amount, b.Balance)
	}
	b.Balance -= amount
	return nil
}

func main() {
	account := BankAccount{Balance: 1000}
	//withdraw - это снять
	account.Withdraw(1500)
	if err := account.Withdraw(1500); err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Снятие успешно. Баланс: ", account.Balance)
	}
	//Deposit - это внести
	account.Deposit(-500)
	if err := account.Deposit(-500); err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Пополнение успешно. Баланс: ", account.Balance)
	}

	if err := account.Deposit(500); err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Пополнение успешно. Баланс: ", account.Balance)
	}

	if err := account.Withdraw(200); err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Снятие успешно. Баланс: ", account.Balance)
	}
}
