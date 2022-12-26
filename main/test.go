package main

import (
	"exam/replication"
	"log"
)

func main() {
	r := replication.NewReplica([]int{0, 1}, 0)
	if r.IsLeader {
		log.Printf("Is Leader")
	}
}
