package trivial

import "fmt"

func Interfaces() {
	tesla := Car{`Tesla`}
	volvo := Car{`Volvo`}
	boeing := Airplane{`Boeing`}
	vehicles := []Vehicle{tesla, volvo, boeing}
	for _, vehicle := range vehicles {
		vehicle.move()
	}

}

type Vehicle interface {
	move()
}

func Drive(v Vehicle) {
	v.move()
}

type Car struct {
	model string
}

type Airplane struct {
	model string
}

func (v Car) move() {
	fmt.Println(v.model, " ride")
}

func (v Airplane) move() {
	fmt.Println(v.model, " fly")
}
