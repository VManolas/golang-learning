// Example 5a
package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
}

// Example 5b
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func main() {
// 	goku := Saiyan{}
// 	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
// }

// Example 5c
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func main() {
// 	goku := Saiyan{
// 		Name: "Goku",
// 	}
// 	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
// 	goku.Power = 9000
// 	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
// }

// Example 5d
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func main() {
// 	goku := Saiyan{"Goku", 9000}
// 	fmt.Printf("%s's power is %d\n", goku.Name, goku.Power)
// }

// Example 5e - introducing pointers
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func main() {
// 	goku := Saiyan{"Goku", 9000}
// 	Super(goku)
// 	fmt.Println(goku.Power) //what do you think is the result?
// }

// func Super(s Saiyan) {
// 	s.Power += 10000
// }

// Example 5f - introducing pointers
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func main() {
// 	goku := &Saiyan{"Goku", 9000}
// 	Super(goku)
// 	fmt.Println(goku.Power) //what do you think is the result?
// }

// func Super(s *Saiyan) {
// 	s.Power += 10000
// }

// Example 5g - introducing pointers
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func main() {
// 	goku := &Saiyan{"Goku", 9000}
// 	Super(goku)
// 	fmt.Println(goku.Power) //what do you think is the result?
// }

// func Super(s *Saiyan) {
// 	s = &Saiyan{"Gohan", 1000}
// }

// Example 5h - Functions on structures
// package main

// import "fmt"

// type Saiyan struct {
// 	Name  string
// 	Power int
// }

// func (s *Saiyan) Super() {
// 	s.Power += 10000
// }

// // func Super(s *Saiyan) {
// // 	s.Power += 10000
// // }

// func main() {
// 	goku := &Saiyan{"Goku", 9000}
// 	fmt.Println(goku.Power)
// 	goku.Super()
// 	//Super(goku)
// 	fmt.Println(goku.Power) //what do you think is the result?
// }
