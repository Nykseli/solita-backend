package main

import (
	"fmt"

	"github.com/Nykseli/solita-backend/db"
)

func main() {
	data := db.GetAllSortedAmount()
	fmt.Println(data.Names[0].Amount)
	data = db.GetAllSortedName()
	fmt.Println(data.Names[0].Name)
	fmt.Println(db.GetNameAmount("anna"))
	fmt.Println(db.GetTotalNameAmount())
}
