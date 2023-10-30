package main

import (
	"Go_Practice/basic"
)

func main() {
	bas := basic.ProviderBasicLog()
	bas.Info()
	bas.Debug()

	///////
	a := basic.A{}
	pa := basic.ProviderBasicLog(basic.WithCustomerOption(&a))
	pa.Info()
	pa.Debug()

	////////
	b := basic.B{}
	pb := basic.ProviderBasicLog(basic.WithCustomerOption(&b))
	pb.Info()
	pb.Debug()
}
