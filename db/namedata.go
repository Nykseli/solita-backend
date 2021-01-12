package db

import (
	"encoding/json"
	"sort"
	"strings"
)

// NameCount contains a single name entry
type NameCount struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

// NameData contains list of name entries
type NameData struct {
	Names []NameCount `json:"names"`
}

var nameDataInstance *NameData

// InitNameData initializes the name data
// This should be called before accessing the namedata
func initNameData() {
	// Currently the name data is static so it needs to be loaded once
	if nameDataInstance != nil {
		return
	}

	nameDataInstance = &NameData{}
	db := GetDBInstance()
	err := json.Unmarshal([]byte(db.data), nameDataInstance)
	if err != nil {
		panic("Couldn't decode names.json")
	}
}

// Create hard copy of nameDataInstance
func createDataCopy() NameData {
	copied := NameData{}
	copied.Names = make([]NameCount, len(nameDataInstance.Names))
	copy(copied.Names, nameDataInstance.Names)
	return copied
}

// GetAllSortedAmount returns name data, ordered by amount, most popular first
func GetAllSortedAmount() NameData {
	initNameData()
	sorted := createDataCopy()
	sort.Slice(sorted.Names, func(i, j int) bool {
		return sorted.Names[i].Amount > sorted.Names[j].Amount
	})

	return sorted
}

// GetAllSortedName returns name data in alphabetical order
func GetAllSortedName() NameData {
	initNameData()
	sorted := createDataCopy()
	sort.Slice(sorted.Names, func(i, j int) bool {
		return sorted.Names[i].Name < sorted.Names[j].Name
	})

	return sorted
}

// GetNameAmount retuns amount of name. 0 if name is not found
// name is case insensitive
func GetNameAmount(name string) int {
	initNameData()

	amount := 0
	lowerName := strings.ToLower(name)
	for _, nameCount := range nameDataInstance.Names {
		if lowerName == strings.ToLower(nameCount.Name) {
			amount = nameCount.Amount
			break
		}
	}

	return amount
}

// GetTotalNameAmount returns total amount of names in name data
func GetTotalNameAmount() int {
	initNameData()

	total := 0
	for _, nameCount := range nameDataInstance.Names {
		total += nameCount.Amount
	}

	return total
}
