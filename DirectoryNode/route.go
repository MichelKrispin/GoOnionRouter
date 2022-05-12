package main

import (
	"errors"
	"math/rand"
)

// Returns a random route through the available nodes
func getRoute(nodes []string) ([3]string, error) { // (firstHop string, secondHop string, thirdHop string) {
	l := len(nodes)
	if l < 3 {
		return [3]string{}, errors.New("Not enough registered nodes")
	}
	indices := [3]int{-1, -1, -1}
	for i := 0; i < 3; i++ {
		idx := rand.Intn(l)
		if indices[0] == idx || indices[1] == idx {
			i--
		} else {
			indices[i] = idx
		}
	}
	return [3]string{nodes[indices[0]], nodes[indices[1]], nodes[indices[2]]}, nil
}
