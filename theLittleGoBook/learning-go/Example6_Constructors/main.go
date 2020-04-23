// Example 6a
package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

//  Structures don't have constructors.
// Instead, you create a function that returns an **instance** of the desired type (like a factory)
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

// The factory doesn't have to return a pointer
func NewSaiyanV(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}

func NewSaiyanDaddy(name string, power int, father *Saiyan) *Saiyan {
	return &Saiyan{
		Name:   name,
		Power:  power,
		Father: father,
	}
}

func main() {
	gokuValue := new(Saiyan)
	fmt.Printf("%s's power is %d\n", gokuValue.Name, gokuValue.Power)
	//fmt.Printf("%s's father is %s and his power is %d\n", gokuValue.Name, gokuValue.Father.Name, gokuValue.Father.Power)
	gokuValue.Name = "gokuV"
	gokuValue.Power = 9000
	fmt.Printf("%s's power is %d\n", gokuValue.Name, gokuValue.Power)
	// same as
	gokuPointer := &Saiyan{}
	fmt.Printf("%s's power is %d\n", gokuPointer.Name, gokuPointer.Power)
	//fmt.Printf("%s's father is %s and his power is %d\n", gokuPointer.Name, gokuPointer.Father.Name, gokuPointer.Father.Power)
	gokuPointer = &Saiyan{
		Name:  "gokuP",
		Power: 9001,
	}
	fmt.Printf("%s's power is %d\n", gokuPointer.Name, gokuPointer.Power)
	// initiating with values using func NewSaiyan for goku but not for his father
	gokuPointerII := NewSaiyan("gokuΠII", 10000)
	fmt.Printf("%s's power is %d\n", gokuPointerII.Name, gokuPointerII.Power)
	//fmt.Printf("%s's father is %s and his power is %d\n", gokuPointerII.Name, gokuPointerII.Father.Name, gokuPointerII.Father.Power)

	// initiating with values using func NewSaiyanV for goku but not for his father
	gokuValueII := NewSaiyanV("gokuVII", 10002)
	fmt.Printf("%s's power is %d\n", gokuValueII.Name, gokuValueII.Power)
	//fmt.Printf("%s's father is %s and his power is %d\n", gokuValueII.Name, gokuValueII.Father.Name, gokuValueII.Father.Power)
	//panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x109ee52]

	// goroutine 1 [running]:
	// main.main()
	//         /Users/vasileiosmanolas/code/go/src/learning-go/Example6_Constructors/main.go:61 +0x6b2
	// exit status 2

	gokuValueFather := &Saiyan{
		Name:  "gokuSon",
		Power: 19000,
		Father: &Saiyan{
			Name:   "gokuFather",
			Power:  19001,
			Father: nil,
		},
	}
	fmt.Printf("%s's power is %d\n", gokuValueFather.Name, gokuValueFather.Power)
	fmt.Printf("%s is son of %s, whose power is %d\n", gokuValueFather.Name, gokuValueFather.Father.Name, gokuValueFather.Father.Power)
	//fmt.Printf("%s's grandpa is %s\n", gokuValueFather.Name, gokuValueFather.Father.Father.Name)

	gokuValueFatherII := NewSaiyanDaddy("gokuSonII", 29000, NewSaiyan("gokuFatherII", 29002))
	fmt.Printf("%s's power is %d\n", gokuValueFatherII.Name, gokuValueFatherII.Power)
	fmt.Printf("%s is son of %s, whose power is %d\n", gokuValueFatherII.Name, gokuValueFatherII.Father.Name, gokuValueFatherII.Father.Power)
}
