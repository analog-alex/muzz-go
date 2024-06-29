package workers

import (
	"log"
	"muzz-service/pkg/dao"
)

var queue = make(chan int)

func start() {
	defer close(queue)

	log.Println("Starting ranker worker")

	// infinite loop
	for {
		id := <-queue

		log.Printf("Ranker worker processing id %d", id)

		// increment the swipes on user
		err := dao.IncrementsLikesForUser(id)
		if err != nil {
			log.Printf("Error incrementing likes for user %d: %v", id, err)
		}
	}
}

func init() {
	go start()
}

func GetRankerQueue() chan int {
	return queue
}
