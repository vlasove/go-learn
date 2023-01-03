package main

import (
	"log"
	"time"
)

func main() {
	scores := make([]int, 0, 10)
	//scores = append(scores, 100500)
	scores = scores[:8]
	scores[5] = 10
	log.Println("slice:", scores)

	// cap grew
	nums := make([]int, 0, 10)

	log.Println("initial cap is:", cap(nums))
	time.Sleep(3 * time.Second)
	for i := 0; i < 100000; i++ {
		nums = append(nums, i)
		if i%100 == 0 {
			log.Println("len of slice:", len(nums), "cap of slice:", cap(nums))
		}
	}

	var names []string
	log.Println("Names is:", names)
	names = append(names, "bob")
	log.Println("Names after is:", names)
}
