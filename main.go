package main

import (
	"fmt"
)

type sayan struct {
	Name             string
	Life             float32
	SpiritualStength float32
	Force            float32
}

func (person *sayan) magichit(points int) {
	person.SpiritualStength -= float32(points)
}

func checkSS(person *sayan) {
	switch {
	case person.SpiritualStength <= 0:
		person.SpiritualStength = 1
	}
}

func (person *sayan) hit(points int) {
	checkSS(person)

	person.Life -= float32(points) / person.SpiritualStength

}

func (person *sayan) isAlive() (alive bool) {
	switch {
	case person.Life <= 0:
		alive = false
	default:
		alive = true
	}
	return
}

func main() {

	goku := goku(100, 200, 10)

	fmt.Printf("%.0f\n", goku.Life)
}

func goku(life, ss, colpo int) (goku sayan) {

	goku.Life = float32(life)
	goku.SpiritualStength = float32(ss)

	goku.hit(colpo)

	return
}
