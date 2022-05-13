package main

func nodesContain(nodes []node, newNode string) bool {
	for _, node := range nodes {
		if node.Address == newNode {
			return true
		}
	}
	return false
}

func isRegistered(nodes []node, newConnection connection) bool {
	for _, node := range nodes {
		if node.Address == newConnection.From || node.Address == newConnection.To {
			return true
		}
	}
	return false
}
