package main

import "fmt"

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
}
