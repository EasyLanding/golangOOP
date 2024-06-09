package main

import (
	"errors"
	"fmt"
)

func helloWorld() string {
	return `Hello World!`
}

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

type SavingsAccount struct {
	balance    float64
	minBalance float64
}

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	balance float64
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable {
	s := &StatisticProfit{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
}

func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		return s.getAverageProfit() / s.Average(s.product.Buys) * 100
	}
}

func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
	}
}

func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
}

func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

func (s *StatisticProfit) GetAverageProfit() float64 {
	if s.getAverageProfit != nil {
		return s.getAverageProfit()
	}
	return 0
}

func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	if s.getAverageProfitPercent != nil {
		return s.getAverageProfitPercent()
	}
	return 0
}

func (s *StatisticProfit) GetCurrentProfit() float64 {
	if s.getCurrentProfit != nil {
		return s.getCurrentProfit()
	}
	return 0
}

func (s *StatisticProfit) GetDifferenceProfit() float64 {
	if s.getDifferenceProfit != nil {
		return s.getDifferenceProfit()
	}
	return 0
}

func (s *StatisticProfit) GetAllData() []float64 {
	if s.getAllData != nil {
		return s.getAllData()
	}
	return nil
}

func (s *StatisticProfit) Average(prices []float64) float64 {
	sum := s.Sum(prices)
	return sum / float64(len(prices))
}

func (s *StatisticProfit) Sum(prices []float64) float64 {
	sum := 0.0
	for _, price := range prices {
		sum += price
	}
	return sum
}

func (c *CurrentAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	c.balance += amount
	return nil
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if c.balance < amount {
		return errors.New("insufficient funds")
	}
	c.balance -= amount
	return nil
}

func (c *CurrentAccount) Balance() float64 {
	return c.balance
}

func (s *SavingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	s.balance += amount
	return nil
}

func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if s.balance-amount < s.minBalance {
		return errors.New("cannot withdraw, minimum balance requirement not met")
	}
	s.balance -= amount
	return nil
}

func (s *SavingsAccount) Balance() float64 {
	return s.balance
}

func ProcessAccount(a Account) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

//Abstract bill in restorant

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity should be greater than 0")
	}
	d.orderDetails[item] += quantity
	return nil
}

func (d *DineInOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; !ok {
		return errors.New("item not found in order")
	}
	delete(d.orderDetails, item)
	return nil
}

func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity should be greater than 0")
	}
	t.orderDetails[item] += quantity
	return nil
}

func (t *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := t.orderDetails[item]; !ok {
		return errors.New("item not found in order")
	}
	delete(t.orderDetails, item)
	return nil
}

func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

// Abstract electron commerce
type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	balance float64
}

func (c *CreditCard) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("недопустимая сумма платежа")
	}
	if c.balance < amount {
		return errors.New("недостаточный баланс")
	}
	c.balance -= amount
	fmt.Printf("Оплачено %.2f с помощью кредитной карты\n", amount)
	return nil
}

type Bitcoin struct {
	balance float64
}

func (b *Bitcoin) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("недопустимая сумма платежа")
	}
	if b.balance < amount {
		return errors.New("недостаточный баланс")
	}
	b.balance -= amount
	fmt.Printf("Оплачено %.2f с помощью биткоина\n", amount)
	return nil
}

func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

func main() {
	resultHelloWorld := helloWorld()

	fmt.Println(resultHelloWorld)

	cocaCola := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{100, 120, 90},
		Buys:          []float64{70, 80, 60},
		CurrentPrice:  150,
		ProfitPercent: 20,
	}

	statistic := NewStatisticProfit(
		func(s *StatisticProfit) { WithAverageProfit(s) },
		func(s *StatisticProfit) { WithAverageProfitPercent(s) },
		func(s *StatisticProfit) { WithCurrentProfit(s) },
		func(s *StatisticProfit) { WithDifferenceProfit(s) },
	)

	statistic.SetProduct(cocaCola)

	fmt.Println("Average Profit:", statistic.GetAverageProfit())
	fmt.Println("Average Profit Percent:", statistic.GetAverageProfitPercent())
	fmt.Println("Current Profit:", statistic.GetCurrentProfit())
	fmt.Println("Difference Profit:", statistic.GetDifferenceProfit())

	c := &CurrentAccount{}
	s := &SavingsAccount{minBalance: 500}
	ProcessAccount(c)
	ProcessAccount(s)

	//Abstract bill in restorant
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)

	//Abstract electrone commerce
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}

	ProcessPayment(cc, 200.00)
	ProcessPayment(btc, 1.00)
}
