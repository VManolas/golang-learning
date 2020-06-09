// Example 20a Stringers
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
// type Stringer interface {
//    String() string
// }
// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v is currently %v years old.", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Akis", 37}
// 	z := Person{"Andreas", 38}
// 	fmt.Println(a, z)
// }

// Example 20b Exercise Stringers
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipaddr IPAddr) String() string {
	// 1.
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
	// 2.
	// var s []string = make([]string, len(ipaddr))

	// for i, val := range ipaddr {
	// 	s[i] = fmt.Sprint(int(val)) // why int?
	// 	fmt.Println(s[i])
	// }
	// fmt.Println(strings.Join(s, "."))
	// return strings.Join(s, ".")
	// 3.
	// var ipaddr_list = make([]string, len(ipaddr))
	// for i, val := range ipaddr {
	// 	ipaddr_list[i] = strconv.Itoa(int(val))
	// }
	// return strings.Join(ipaddr_list, ".")
	// 4.
	// var s string
	// for i := range ipaddr {
	// 	if i == 0 {
	// 		s += fmt.Sprint(int(ipaddr[i]))
	// 	} else {
	// 		s += "." + fmt.Sprint(ipaddr[i])
	// 	}
	// }
	// return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
