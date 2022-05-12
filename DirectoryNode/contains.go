package main

func nodesContain(nodes []string, newNode string) bool {
	for _, node := range nodes {
		if node == newNode {
			return true
		}
	}
	return false
}

func isRegistered(nodes []string, newConnection connection) bool {
	for _, node := range nodes {
		if node == newConnection.From || node == newConnection.To {
			return true
		}
	}
	return false
}
