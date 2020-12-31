package main

import "fmt"

type IPv4 []byte

func (i IPv4) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {

	ipv4 := map[string]IPv4{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}

	for k, v := range ipv4 {
		fmt.Printf("%v: %v\n", k, v)
	}
}
