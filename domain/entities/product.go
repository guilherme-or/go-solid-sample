package entities

type Product struct {
	Id          int64
	Name        string
	Description string
	Price       float64
	CategoryId  int64
	Active      bool
}
