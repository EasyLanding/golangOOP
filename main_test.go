package main

import "testing"

func TestHelloWorld(t *testing.T) {
	execting := `Hello World!`
	result := helloWorld()

	if result != execting {
		t.Errorf("HelloWorld() returned %s, expected %s", result, execting)
	}
}

func TestCurrentAccount(t *testing.T) {
	c := &CurrentAccount{}

	err := c.Deposit(100)
	if err != nil {
		t.Errorf("Deposit failed: %v", err)
	}

	err = c.Withdraw(50)
	if err != nil {
		t.Errorf("Withdrawal failed: %v", err)
	}

	balance := c.Balance()
	if balance != 50 {
		t.Errorf("Incorrect balance. Expected 50, got %.2f", balance)
	}
}

func TestSavingsAccount(t *testing.T) {
	account := &SavingsAccount{balance: 100, minBalance: 200}

	// Пополняем счет на 600
	account.Deposit(600)

	// Проверяем баланс после пополнения
	if account.Balance() != 700 {
		t.Errorf("Ожидаемый баланс после пополнения: 700, получено: %v", account.Balance())
	}

	// Пытаемся снять 200
	err := account.Withdraw(200)

	// Проверяем, что снятие прошло успешно
	if err != nil {
		t.Errorf("Ошибка при снятии средств: %v", err)
	}

	// Проверяем баланс после снятия
	if account.Balance() != 500 {
		t.Errorf("Ожидаемый баланс после снятия: 500, получено: %v", account.Balance())
	}
}
