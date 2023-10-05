package main

import "log"
// test
type myStruct struct {
	FirstName string
}
// The * in front of the myStruct is signafying that myStruct is the reciever inside of this function
func (m *myStruct) printName() string {
	return m.FirstName
}
func main() {
	log.Println("Begin...")
	var myVar myStruct
	myVar.FirstName = "John"
	myVar2 := myStruct{
		FirstName: "Mary",
	}
	log.Println("MyVar  =  ", myVar.printName())
	log.Println("MyVar2  =  ", myVar2.printName())

}
