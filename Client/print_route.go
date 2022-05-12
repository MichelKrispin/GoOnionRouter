package main

import "log"

func printRoute(firstHop string, hops []string) {
	result := "The route is\n\t[" + firstHop + "] -> "
	for i, hop := range hops {
		result += "[" + hop + "]"
		if i != len(hops)-1 {
			result += " -> "
		}
	}
	log.Println(result)
}
