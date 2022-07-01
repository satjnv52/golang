package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Golang")
	var fruitList = []string{"A", "B", "C", "D"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("New fruitlist after append", fruitList)
	index := 2
	fruitList = append(fruitList[:index], fruitList[(index+1):]...)
	fmt.Println("After deleting C:", fruitList)
	//fruitList = append(fruitList[:3])
	fmt.Println("List after making slices of slice", fruitList)

	scores := make([]int, 4)
	scores[0] = 234
	scores[1] = 344
	scores[2] = 984
	scores[3] = 244
	scores[1] = 2234
	scores = append(scores, 222, 333, 444, 555)
	fmt.Println("Scores are:", scores)
	for i := range scores {
		fmt.Printf("At place %v value is %v\n", i, scores[i])
	}
	sort.Ints(scores)

}
