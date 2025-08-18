package task_2

import "fmt"

func Task2() {
	myMoto := Moto{"Kawasaki", 650, 150}
	myCar := Car{"Mersedes", 3000, "Silver"}
	myMoto.Start("Vruim Vruiiiim tatatata")
	myMoto.Stop()
	myCar.Start("Mrrrrruuuur Mruuuuu")
	myCar.Stop()

}

type Vehicle interface {
	Start(sound string)
	Stop()
}

type Moto struct {
	name   string
	engVol int
	mass   int
}

func (m Moto) Start(sound string) {
	fmt.Println("Moto made", sound)
}
func (m Moto) Stop() {
	fmt.Println("Moto made vzhukhhhhhh")
}

type Car struct {
	name   string
	engVol int
	color  string
}

func (c Car) Start(sound string) {
	fmt.Println("Auto made", sound)
}
func (c Car) Stop() {
	fmt.Println("Auto made whuhhh")
}
