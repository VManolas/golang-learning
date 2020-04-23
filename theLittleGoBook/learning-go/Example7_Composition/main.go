// Example 7a
package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) /* receiver */ Introduce( /* input(s) of fuction */ ) /*return type of func*/ {
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

type Dog struct {
	Name  string
	Power int
}

func (d *Dog) IntroduceD() {
	fmt.Printf("Gav, I'm %s, and my dog power is %d\n", d.Name, d.Power)
}

type SaiyanD struct {
	*Dog  //because we didn't give it an explicit name, we can implicitly access the fields and functions of the composed type
	Power int
}

func (sd *SaiyanD) Transform() {
	power := sd.Dog.Power
	sd.Power = power * 2
}

func (sd *SaiyanD) PowerBoost() {
	sd.Power = sd.Power + 10000
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

	gokuD := &SaiyanD{
		Dog: &Dog{"GokuD", 1999},
	}

	gokuD.IntroduceD()
	fmt.Println(gokuD.Dog.Power)
	fmt.Println(gokuD.Power)
	gokuD.Transform()
	fmt.Println(gokuD.Power)
	gokuD.PowerBoost()
	fmt.Println(gokuD.Power)
}
