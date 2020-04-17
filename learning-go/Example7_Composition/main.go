// Example 7a
package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person //because we didn't give it an explicit name, we can implicitly access the fields and functions of the composed type
	Power   int
}

func (s *Saiyan) IntroduceS() {
	fmt.Printf("Hi, I'm %s\n", s.Name)        // implicit composition
	fmt.Printf("Hi, I'm %s\n", s.Person.Name) // composed version (explicit composition?)
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"}, // however the Go compiler did give it an explicit field name
		Power:  9001,
	}
	goku.Introduce()
	// also valid
	fmt.Println(goku.Name) // implicit composition ???
	// also valid
	fmt.Println(goku.Person.Name)

	goku.IntroduceS()
}
