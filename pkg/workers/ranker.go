package workers

import (
	"log"
	"muzz-service/pkg/dao"
)

// This worker simulates a 'event-driven' system that reacts to swipe events.
// The idea here is to receive a user id everytime a user is swiped right on.
// This is a very primitive ranking system that just increments the number of likes.
//
// The more likes a user has, the higher the ranking.
// the likes are stored in the user table and can be used to order users.
// by 'attractiveness'
//
// Should be called by the swipe_handler.

var queue = make(chan int)

func start() {
	log.Println("Starting ranker worker")
	defer close(queue)

	// infinite loop
	// we await 'events' from the channel
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
