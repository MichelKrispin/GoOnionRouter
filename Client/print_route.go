package main

import "log"

func printRoute(hops route) {
	result := "The route is\n\t"
	for i, hop := range hops.Nodes {
		result += "[" + hop.Address + "]"
		if i != len(hops.Nodes)-1 {
			result += " -> "
		}
	}
	log.Println(result)
}
