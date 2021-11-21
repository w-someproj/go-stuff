package trivial

import "fmt"

func Structs() {
	tom := Person{Name: `Tom`, Age: 24, Contact: ContactInfo{}}
	bob := Person{Name: `Bob`, Age: 24}
	fmt.Println(tom.Name)
	fmt.Println(bob)
	tomPointer := &tom
	tomPointer.Age = 29
	fmt.Println(tom.Age)
	(*tomPointer).Age = 32
	fmt.Println(tom.Age)
	tom.Email = `tom`
	fmt.Println(tom.Email)
	tom.Contact.Email = `supertom`
	fmt.Println(tom.Email)
	tom.Print()
	tom.ChangeAge(17)
	tom.Print()
}

type Person struct {
	Name string
	Age  int
	ContactInfo
	Contact ContactInfo
}

type ContactInfo struct {
	Email string
	Phone string
}

func (p Person) Print() {
	fmt.Println(p.Name, p.Age)
}

func (p *Person) ChangeAge(newAge int) {
	p.Age = newAge
}
