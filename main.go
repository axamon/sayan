package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	colpo, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err.Error())
	}
	var goku sayan

	goku.Life = 100
	goku.SpiritualStength = 200

	goku.hit(colpo)

	fmt.Printf("%.0f\n", goku.Life)
}
