package main

import (
	"fmt"
)

type Engineer struct {
	Id      int
	FName   string
	LName   string
	Age     int
	Address *Address
}

type Address struct {
	City  string
	State string
}

type Log struct {
	Id      int
	Message string
}

func main() {
	engineers := getEngineers()
	// ruleid: invalid-usage-of-modified-variable
	eng1, err := getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng1.FName)
	}

	// ruleid: invalid-usage-of-modified-variable
	if eng2, err := getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng2.FName)
	}

	eng3 := Engineer{4, "Lee", "Renaldo", 50, nil}
	// ruleid: invalid-usage-of-modified-variable
	eng3.Address, err = getEngineerAddressByIndex(engineers, 1)
	if err != nil {
		fmt.Printf("Unable to obtain address %#v!\n", eng3.Address)
	}

	// ruleid: invalid-usage-of-modified-variable
	eng4, err := getEngineerAtIndex(engineers, 5)
	if err != nil {
		log := Log{
			Id:      eng4.Id,
			Message: "Unable to obtain engineer",
		}
		fmt.Printf("%#v\n", log)
	}

	// ok: invalid-usage-of-modified-variable
	if eng5, err := getEngineerAtIndex(engineers, 6); err == nil {
		fmt.Printf("Obtained engineer with FName %s\n", eng5.FName)
	}

	// ok: invalid-usage-of-modified-variable
	eng6, err := getEngineerAtIndex(engineers, 7)
	if err != nil {
		eng6 = &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng6.Id)
	}

	// ok: invalid-usage-of-modified-variable
	eng7, err := getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng7 = &Engineer{0, "N/A", "N/A", 0, nil}
	}

	fmt.Printf("Engineer 1: %s", fmt.Sprintf("%s %s", eng1.FName, eng1.LName))
	fmt.Printf("Engineer 7: %s", fmt.Sprintf("%s %s", eng7.FName, eng7.LName))
}

func getEngineerAtIndex(slice []Engineer, idx int) (*Engineer, error) {
	if idx >= len(slice) {
		return nil, fmt.Errorf("invalid index")
	}
	return &slice[idx], nil
}

func getEngineerAddressByIndex(slice []Engineer, idx int) (*Address, error) {
	if idx >= len(slice) {
		return nil, fmt.Errorf("invalid index")
	}
	return slice[idx].Address, nil
}

func getEngineers() []Engineer {
	return []Engineer{
		{
			1,
			"James",
			"Osterberg",
			24,
			&Address{
				"Detroit",
				"Michigan",
			},
		},
		{
			2,
			"Althea",
			"Gordon",
			43,
			&Address{
				"New York",
				"New York",
			},
		},
		{
			3,
			"José-Manuel",
			"Arthur Chao",
			55,
			&Address{
				"Paris",
				"Île-de-France",
			},
		},
	}
}
