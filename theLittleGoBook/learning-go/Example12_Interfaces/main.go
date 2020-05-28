// Example 12a
// package main

// type Logger interface {
// 	Log(message string)
// }

// // we might have various types of loggers:
// type SqlLogger struct{}
// type ConsoleLogger struct{}
// type FileLogger struct{}

// // how to use one: just like any other type
// // it could be a structure's field:
// type Server struct {
// 	logger Logger
// }

// // or a function parameter (or return value)
// func process(logger Logger) {
// 	logger.Log("Hello")
// }

// func main() {

// }
// Example 12b- (from [Nestoras' multiple interface example](https://github.com/nestoras/go-playground/blob/master/general/05_interface_multiple/main.go))
// package main

// import "fmt"

// type Walker interface {
// 	Walk()
// }
// type Sleeper interface {
// 	Sleep()
// }
// type Person struct {
// 	Name string
// }

// func (p Person) Walk() {
// 	fmt.Println(p.Name + " is walking...")
// }
// func (p Person) Sleep() {
// 	fmt.Println(p.Name + " is sleeping...")
// }
// func main() {
// 	p := Person{Name: "Takis"}
// 	p.Walk()
// 	p.Sleep()
// }

// Example 12b (from [Nestoras' type assertion interface example](https://github.com/nestoras/go-playground/blob/master/general/06_interface_type_assertion/main.go))
// package main

// import "fmt"

// type Walker interface {
// 	Walk()
// }

// type Sleeper interface {
// 	Sleep()
// }

// type Flyer interface {
// 	Fly()
// }

// type Rider interface {
// 	Ride()
// }

// type Exister interface {
// 	Exist()
// }

// type Person struct {
// 	Name string
// }

// func (p Person) Exist() {
// 	fmt.Println(p.Name + " exists.")
// }

// func (p Person) Walk() {
// 	fmt.Println(p.Name + " walking...")
// }

// func (p Person) Sleep() {
// 	fmt.Println(p.Name + " sleeping...")
// }

// func (p Person) Ride() {
// 	fmt.Println(p.Name + " riding...")
// }

// type Bird struct {
// 	Species string
// }

// func (b Bird) Exist() {
// 	fmt.Println(b.Species + " exists.")
// }

// func (b Bird) Walk() {
// 	fmt.Println(b.Species + " walking...")
// }

// func (b Bird) Sleep() {
// 	fmt.Println(b.Species + " sleeping...")
// }

// func (b Bird) Fly() {
// 	fmt.Println(b.Species + " flying...")
// }

// type Robot struct {
// 	Name string
// }

// func (r Robot) Exist() {
// 	fmt.Println(r.Name + " exists.")
// }

// func (r Robot) Sleep() {
// 	fmt.Println(r.Name + " sleeps.")
// }

// func implementInterface(e Exister) {
// 	_, ok := e.(Walker)
// 	fmt.Printf("implements interface Walker? %v\n", ok)
// 	_, ok = e.(Sleeper)
// 	fmt.Printf("implements interface Sleeper? %v\n", ok)
// 	_, ok = e.(Flyer)
// 	fmt.Printf("implements interface Flyer? %v\n", ok)
// 	_, ok = e.(Rider)
// 	fmt.Printf("implements interface Rider? %v\n", ok)
// }

// func main() {
// 	var e Exister = Person{Name: "Nestoras"}
// 	implementInterface(e)
// 	var f Exister = Bird{Species: "Chicken"}
// 	implementInterface(f)
// 	var g Exister = Robot{Name: "Alexa"}
// 	g.Exist()
// 	implementInterface(g)
// }

// Example 12c (from [nestoras empty Interface example](https://github.com/nestoras/go-playground/blob/master/general/04_interface_empty/main.go))
// package main

// import "fmt"

// type Dog struct {
// 	Name string
// }

// type Car struct {
// 	Name         *string
// 	Model        string
// 	Manufacturer string
// }

// func describe(i interface{}) {
// 	fmt.Printf("This is type of %T with value %v\n", i, i)
// }

// func main() {
// 	dog := Dog{Name: "Azor"}
// 	car := Car{Model: "I8", Manufacturer: "BMW"}
// 	describe(car)
// 	describe(dog)
// }

// Example 12d- (from [Nestoras' embedding example](https://github.com/nestoras/go-playground/blob/master/general/02_embedding/main.go))
// package main

// import "fmt"

// type Vehicle struct{}

// func (v Vehicle) Start() {
// 	fmt.Println("Start the engine")
// }

// type Car struct {
// 	Vehicle
// } // this is not inheritance. This is embedding

// func main() {
// 	c := Car{Vehicle{}}
// 	c.Start()
// }

// Example 12d (from [nestoras Interface example](https://github.com/nestoras/go-playground/blob/master/general/03_interface/main.go))
// package main

// import "fmt"

// // We should try to create single method interfaces since we can easily compose them to bigger
// type StartEngine interface {
// 	Start()
// }

// type StopEngine interface {
// 	Stop()
// }

// // Given the above single method interfaces it is easy to compose a bigger one as follows:
// type StartStopEngine interface {
// 	StartEngine
// 	StopEngine
// }

// type Vehicle struct {
// 	StartEngine
// 	//StopEngine
// }

// type FuelControl struct {
// 	StartEngine
// 	Fuel float32
// }

// func (fc FuelControl) Start() {
// 	fmt.Println("Fuel control")
// 	if fc.Fuel > 0.5 {
// 		fc.StartEngine.Start()
// 		fc.Start() // what is the difference? implicit vs explicit above?
// 	} else {
// 		fmt.Println("Cannot start")
// 	}
// }

// // If a struct has a method that matches the above Start() it automatically satisfies the interface.
// type Car struct{}

// func (c Car) Start() {
// 	fmt.Println("Start the car")
// }
// func (c Car) Stop() {
// 	fmt.Println("Stop the car")
// }

// type Bicycle struct{}

// func (m Bicycle) Ride() {
// 	fmt.Println("Ride the bicycle")
// }

// type Motorbike struct{}

// func (m Motorbike) Start() {
// 	fmt.Println("Start the motorbike")
// }
// func (m Motorbike) Stop() {
// 	fmt.Println("Stop the motorbike")
// }

// func main() {
// 	v := Vehicle{FuelControl{Car{}, 5}}
// 	v.Start()

// 	b := Bicycle{}
// 	b.Ride()

// 	mb := Vehicle{FuelControl{Motorbike{}, 100}}
// 	mb.Start()
// }

// Example 12e (from [Go by Example: Interfaces](https://gobyexample.com/interfaces))
// package main

// import (
// 	"fmt"
// 	"math"
// )

// type geometry interface {
// 	area() float64
// 	perim() float64
// }

// type rect struct {
// 	width, height float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }
// func (r rect) perim() float64 {
// 	return 2 * (r.width + r.height)
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// func main() {
// 	r := rect{width: 3, height: 4}
// 	measure(r)
// 	c := circle{radius: 5}
// 	measure(c)
// }

// Example 12f (from [Nestoras' type switch interface example](https://github.com/nestoras/go-playground/blob/master/general/07_type_switch/main.go))
// package main

// import "fmt"

// type Dog struct {
// 	Name string
// }

// type Person struct {
// 	Firstname string
// 	Lastname  string
// }

// type Cat struct {
// 	Name string
// }

// func describe(i interface{}) {
// 	switch i.(type) {
// 	case Dog:
// 		fmt.Println("Hello dog ", i.(Dog).Name)
// 	case Person:
// 		fmt.Println("Hello ", i.(Person).Firstname, i.(Person).Lastname)
// 	default:
// 		fmt.Println("Hello unknown ", i
// 	}
// }

// func main() {
// 	describe(Dog{Name: "Azor"})
// 	describe(Person{Firstname: "Stella", Lastname: "Katerina"})
// 	describe(Cat{Name: "Katia"})
// }

// Example 12g from [Nestoras' embedding interface example](https://github.com/nestoras/go-playground/blob/master/general/08_embedding_interface/main.go)
// package main

// import "fmt"

// type Walker interface {
// 	Walk()
// }

// type Sleeper interface {
// 	Sleep()
// }

// type IAction interface {
// 	Walker
// 	Sleeper
// }

// type Person struct {
// 	Name string
// }

// func (p Person) Walk() {
// 	fmt.Println(p.Name + " is walking...")
// }

// func (p Person) Sleep() {
// 	fmt.Println(p.Name + " is sleeping...")
// }

// func main() {
// 	person := Person{Name: "Babis"}
// 	var w Walker = person
// 	var s Sleeper = person
// 	var a IAction = person
// 	fmt.Printf("Dynamic type and value of interface 'w' of static type Walker is '%T' and '%v'\n", w, w)
// 	fmt.Printf("Dynamic type and value of interface 's' of static type Sleeper is '%T' and '%v'\n", s, s)
// 	fmt.Printf("Dynamic type and value of interface 'a' of static type IAction is '%T' and '%v'\n", a, a)
// }

// Example 12h from [Nestoras' pointer value receiver interface example](https://github.com/nestoras/go-playground/blob/master/general/09_pointer_value_receiver/main.go)
// package main

// import "fmt"

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Rect struct {
// 	width  float64
// 	height float64
// }

// func (r Rect) Area() float64 {
// 	return r.width * r.height
// }

// func (r Rect) Perimeter() float64 {
// 	return 2 * (r.height + r.width)
// }

// type ShapeP interface {
// 	AreaP() float64
// 	PerimeterP() float64
// }

// func (r *Rect) AreaP() float64 {
// 	return r.width * r.height
// }

// func (r *Rect) PerimeterP() float64 {
// 	return 2 * (r.height + r.width)
// }

// func main() {
// 	r := Rect{5.0, 4.0}
// 	p := &r
// 	var sr Shape = r
// 	var sp ShapeP = p
// 	// structs vs interfaces (?)
// 	area := sr.Area()
// 	perimeter := sr.Perimeter()
// 	fmt.Println("area of rectangle is", area)
// 	fmt.Println("perimeter of rectangle is", perimeter)
// 	areaP := sp.AreaP()
// 	perimeterP := sp.PerimeterP()
// 	fmt.Println("areaP of rectangle is", areaP)
// 	fmt.Println("perimeterP of rectangle is", perimeterP)
// }

// Example 12i from [How to use interfaces in go](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
// package main

// import "fmt"

// //defining Animal datatypes
// type Animal interface {
// 	// defining an Animal as being anything that can speak
// 	Speak() string
// 	// we define an Animal as being any type that has a method named `Speak()`.
// 	// Any type that defines this method is said to satisfy the Animal interface.
// 	// There is no implements keyword in Go;
// 	// whether or not a type satisfies an interface is determined automatically.
// }

// // This is a core concept in Go’s type system;
// // instead of designing our abstractions in terms of what kind of data our types can hold,
// // we design our abstractions in terms of what actions our types can execute.

// // creating a couple of types that satisfy this interface:
// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// type Cat struct{}

// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// type Kitty struct{}

// // an interface definition does not prescribe whether an implementor should implement the interface using a pointer receiver or a value receiver.
// func (k *Kitty) Speak() string {
// 	return "Meow!"
// }

// type Llama struct{}

// func (l Llama) Speak() string {
// 	return "?????"
// }

// type JavaProgrammer struct{}

// func (j JavaProgrammer) Speak() string {
// 	return "Design patterns!"
// }

// func main() {
// 	// we can create a slice of Animals, and put one of each type into that slice

// 	// animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}, Kitty{}} // cannot use Kitty literal (type Kitty) as type Animal in slice literal: Kitty does not implement Animal (Speak method has pointer receiver)
// 	// What it’s saying is not that the interface Animal demands that you define your method as a pointer receiver,
// 	// but that you have tried to convert a Cat struct into an Animal interface value,
// 	// but only *Cat satisfies that interface.
// 	// You can fix this bug by passing in a *Cat pointer to the Animal slice instead of a Cat value,
// 	// by using new(Cat) instead of Cat{}
// 	animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}
// 	// (you could also say &Cat{}
// 	// animals := []Animal{Dog{}, &Cat{}, Llama{}, JavaProgrammer{}}
// 	// Also works passing in a *Dog pointer instead of a Dog value, without changing the definition of the Dog type’s Speak method
// 	// but recognize a subtle difference:
// 	// we didn’t need to change the type of the receiver of the Speak method.
// 	// This works because a pointer type can access the methods of its associated value type, but not vice versa.
// 	// That is, a *Dog value can utilize the Speak method defined on Dog, but as we saw earlier,
// 	// a Cat value cannot access the Speak method defined on *Cat.
// 	//animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
// 	// That may sound cryptic, but it makes sense when you remember the following: everything in Go is passed by value. Every time you call a function, the data you’re passing into it is copied.

// 	// and see what each animal says
// 	for _, animal := range animals {
// 		fmt.Println(animal.Speak())
// 	}

// 	names := []string{"stanley", "david", "oscar"}
// 	// we have to convert the []string to an []interface{}
// 	// `cannot use names (type []string) as type []interface {} in function argument`
// 	vals := make([]interface{}, len(names))
// 	for i, v := range names {
// 		vals[i] = v
// 	}
// 	PrintAll(vals)
// }

// // An interface value is constructed of two words of data;
// // one word is used to point to a method table for the value’s underlying type,
// // and the other word is used to point to the actual data being held by that value.
// func PrintAll(vals []interface{}) {
// 	for _, val := range vals {
// 		fmt.Println(val)
// 	}
// }

// Example 12j [getting a proper timestamp out of the Twitter API](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// start with a string representation of our JSON data
var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}`

func main() {
	// our target will be of type map[string]interface{}, which is a
	// pretty generic type that will give us a hashtable whose keys
	// are strings, and whose values are of type interface{}
	// var val map[string]interface{}

	var val map[string]time.Time
	//error:
	// parsing time ""Thu May 31 00:00:01 +0000 2012"" as ""2006-01-02T15:04:05Z07:00"":
	// cannot parse "Thu May 31 00:00:01 +0000 2012"" as "2006"
	// what it means is that the string representation we gave it does not match the standard time formatting
	// (because Twitter’s API was originally written in Ruby, and the default format for Ruby is not the same as the default format for Go).
	// We’ll need to define our own type in order to unmarshal this value correctly.
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}
