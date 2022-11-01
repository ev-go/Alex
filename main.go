package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Provide seed
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	boxes := rand.Perm(100)
	// fmt.Println(boxes)

	// for i := 0; i < 101; i = i + 1 {
	// 	fmt.Println(boxes[i])
	// }
	var prisoners [100]int

	// prisoners = [100]int{1, 2, 3}

	for i := 0; i < 100; i++ {
		prisoners[i] = i
		// fmt.Println(prisoners[i])
	}

	for i := 0; i < 100; i++ {
		for k := 0; k < 100; k++ {
			if prisoners[i] == boxes[k] {
				fmt.Println("В коробке с номером", k, "Заключенный номер", i, "Нашел свой номер")
			}

		}
	}

}
