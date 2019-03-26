package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGoku(t *testing.T) {
	// call flag.Parse() here if TestMain uses flags
	// os.Exit(m.Run())
	goku := goku(100, 200, 10)

	if goku.Life == 99.95 {
		t.Logf("ok %v", goku.Life)
	} else {
		t.Errorf("Errore %v\n", goku.Life)
	}
}

func Example_hit() {
	var goku sayan

	goku.Life = 100
	goku.SpiritualStength = 200
	goku.hit(10)
	fmt.Printf("%d\n", int(goku.Life))
	// Output:
	// 99
}

func Example_checkSS() {
	var goku sayan

	goku.Life = 100
	goku.SpiritualStength = 0
	goku.hit(10)
	fmt.Printf("%d\n", int(goku.Life))
	// Output:
	// 90
}

func Example_isAlive() {
	var goku sayan

	goku.Life = 100
	alive := goku.isAlive()
	fmt.Println(alive)

	goku.Life = 0
	alive = goku.isAlive()
	fmt.Println(alive)

	goku.Life = -1
	alive = goku.isAlive()
	fmt.Println(alive)
	// Output:
	// true
	// false
	// false

}

func Example_magichit() {
	var goku sayan
	goku.Life = 100
	goku.SpiritualStength = 100
	goku.magichit(10)
	fmt.Println(goku.Life)
	fmt.Println(goku.SpiritualStength)
	// Output:
	// 100
	// 90
}

func TestMain(m *testing.M) {
	runTests := m.Run()
	os.Exit(runTests)
}
