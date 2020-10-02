package types

type Money int64

type Currency string
//Currency types
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)
//Payment Status types
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)
type PAN string

type Card struct {
	ID	int
	PAN	PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color	string
	Name string
	Active	bool
}
type Category string

type Status string
type Payment struct {
	ID string
	Amount Money
	Category Category
	Status Status
}
type PaymentSource struct {
	Type string // 'card'
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
   }

type Phone string

type Account struct {
	ID int64
	Phone Phone 
	Balance Money
} 