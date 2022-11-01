package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ATTEMPTS = 50

func main() {

	//Provide seed

	//Generate a random array of length n

	// fmt.Println(boxes)

	// for i := 0; i < 101; i = i + 1 {
	// 	fmt.Println(boxes[i])
	// }

	// prisoners = [100]int{1, 2, 3}

	var numberOfIterations int
	fmt.Println("enter number of iterations")
	fmt.Scan(&numberOfIterations)
	var resultOfFind int
	for n := 0; n < numberOfIterations; n++ {
		var addSeed int64 = int64(n)

		resultOfFind = resultOfFind + Find(addSeed)
	}
	fmt.Println("from", numberOfIterations, "iterations", resultOfFind, "is true")
}

func Find(addSeed int64) int {
	var result int
	var prisoners [100]int
	rand.Seed(time.Now().Unix() + addSeed)
	boxes := rand.Perm(100)
	for i := 0; i < 100; i++ {
		prisoners[i] = i

	}
	for i := 0; i < 100; i++ {
		chosenbox := prisoners[i]

		for k := 0; k <= ATTEMPTS; k++ {
			if prisoners[i] == boxes[chosenbox] {
				// fmt.Println("prisoner №", prisoners[i], "find his number in box №", chosenbox, "OK")
				result = 1
				break
			} else {
				// fmt.Println("prisoner №", prisoners[i], "open box №", chosenbox, "and find number", boxes[chosenbox], "and going to open it")
				chosenbox = boxes[chosenbox]
			}
			if k == ATTEMPTS {
				// fmt.Println("FAIL!!! prisoner №", prisoners[i], "not find his number")
				return 0
			}

		}
	}
	return result
}

// for k := 0; k < 100; k++ {
// 	if prisoners[i] == boxes[k] {
// 		fmt.Println("В коробке с номером", k, "Заключенный номер", i, "Нашел свой номер")
// 	}

// }
