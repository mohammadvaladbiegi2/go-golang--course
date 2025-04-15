// package main

// type Speaker interface {
// 	Speak() string
// }

// type Dog struct {
// 	massage string
// }
// type Human struct {
// 	massage string
// }

// func (t Dog) Speak() string {
// 	return t.massage
// }
// func (t Human) Speak() string {
// 	return t.massage
// }

// func main() {

// 	dog := Dog{massage: "wof wof"}
// 	human := Human{massage: "heeeeeyyyyyyy"}

// 	println(dog.Speak())
// 	println(human.Speak())

// }

// package main

// type Mover interface {
// 	Move() string
// }

// // Structs
// type Bike struct {
// 	Brand string
// }

// type Car struct {
// 	Model string
// }

// // Implementing Mover interface
// func (b Bike) Move() string {
// 	return "Bike " + b.Brand + " is moving"
// }

// func (c Car) Move() string {
// 	return "Car " + c.Model + " is moving"
// }

// func main() {
// 	b1 := Bike{Brand: "Overloard"}
// 	c1 := Car{Model: "Hunda"}

// 	mover := []Mover{b1, c1}

// 	for i := 0; i < len(mover); i++ {

// 		println(mover[i].Move())
// 	}

// }

package main

import "fmt"

type WriterReader interface {
	Read()
	Write(word string)
}

type MemoryStorage struct {
	massage string
}

func (m *MemoryStorage) Write(word string) {
	m.massage = word
}
func (m MemoryStorage) Read() {
	fmt.Println(m.massage)
}

func main() {

	var storage WriterReader = &MemoryStorage{}
	storage.Write("hi brother")
	storage.Read()

}
