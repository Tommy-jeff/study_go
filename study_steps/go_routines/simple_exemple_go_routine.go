package goroutines

import "time"

func SimpleExemple() {
	go Contador("A")
	go Contador("B")

	time.Sleep(5 * time.Second)

}

func Contador(tipo string) {
	for i := 1; i <= 5; i++ {
		println(tipo, ":", i)
		time.Sleep(time.Second)
	}
}