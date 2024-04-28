package models

type ProductRepo struct {
	db *Data
}
type Product struct {
	ID     int
	PName  string
	Number int
	Price  float32
	Image  string
	Sales  int
}
