package main

import "fmt"

type order string // you can use int also

const (
	received   order = "Received" //if you use int then use "order=iota " iota default value is 0 then it increase by 1 then 2 ,3,4
	processing       = "Processing"
	prepared         = "Prepared"
	confirmed        = "Confirmed"
)

func enums(status order) {
	fmt.Print("Order Status: ", status)
}
func main() {
	enums(processing)
}
