// Example 6a
package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

type SaiyanDaddy struct {
	Name   string
	Power  int
	Father *SaiyanDaddy
}

//  Structures don't have constructors.
// Instead, you create a function that returns and instance of the desired type (like a factory)
func NewSaiyanPointer(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

// The factory doesn't have to return a pointer
func NewSaiyanValue(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	gokuValue := new(Saiyan)
	// same as
	gokuPointer := &Saiyan{}
	fmt.Printf("%s's power is %d\n", gokuValue.Name, gokuValue.Power)
	fmt.Printf("%s's power is %d\n", gokuPointer.Name, gokuPointer.Power)

	gokuValue.Name = "gokuV"
	gokuValue.Power = 9000

	gokuPointer = &Saiyan{
		Name:  "gokuP",
		Power: 9001,
	}

	fmt.Printf("%s's power is %d\n", gokuValue.Name, gokuValue.Power)
	fmt.Printf("%s's power is %d\n", gokuPointer.Name, gokuPointer.Power)

	gokuValueFather := &SaiyanDaddy{
		Name:  "gokuSon",
		Power: 9000,
		Father: &SaiyanDaddy{
			Name:   "gokuFather",
			Power:  9001,
			Father: nil,
		},
	}
	fmt.Printf("%s's power is %d\n", gokuValueFather.Name, gokuValueFather.Power)
	fmt.Printf("%s is son of %s, whose power is %d\n", gokuValueFather.Name, gokuValueFather.Father.Name, gokuValueFather.Father.Power)
	//panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x109ee52]

	// goroutine 1 [running]:
	// main.main()
	//         /Users/vasileiosmanolas/code/go/src/learning-go/Example6_Constructors/main.go:63 +0x6b2
	// exit status 2
	//fmt.Printf("%s's grandpa is %s\n", gokuValueFather.Name, gokuValueFather.Father.Father.Name)
}
