package models

type StudentRepo struct {
	db *Data
}
type Product struct {
	ID     int
	Name   string
	Number int
	Price  float32
	Image  string
	Sales  int
}
