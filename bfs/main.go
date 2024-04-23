package main

import "fmt"

// BFS algorithm
// 1. Keep a queue containing the people to check
// 2. Pop a person from the queue
// 3. Check if this person is a mango seller
// 4. If yes, return this person
// 5. If not, add this person's neighbors to the queue
// 6. Repeat from step 2
// 7. If queue is empty, no mango seller
// 8. Done

var graph map[string][]string

func main() {
	graph = make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	bfs(graph, "you")
}

func bfs(graph map[string][]string, name string) bool {
	fmt.Printf("searching for mango seller starting from %s\n", name)
	var searchQueue []string

	// add this person's neighbors to the search queue
	searchQueue = append(searchQueue, graph[name]...)

	checked := make(map[string]bool)

	// while the search queue is not empty
	for len(searchQueue) != 0 {
		// grabs the first person from the queue
		person := searchQueue[0]
		fmt.Printf("checking %s\n", person)
		// checks whether this person is a mango seller
		if isMangoSeller(person) && !checked[person] {
			// if yes, print it
			fmt.Printf("%s is a mango seller!\n", person)
			return true
		} else {
			// if not, add this person's neighbors to the search queue
			searchQueue = append(searchQueue, graph[person]...)

			// mark this person as checked
			checked[person] = true
		}

		// remove this person from the queue
		searchQueue = searchQueue[1:]
	}

	return false
}

// IsMangoSeller checks if a person is a mango seller
func isMangoSeller(name string) bool {
	// a mango seller is a person whose name ends with 'm'
	return name[len(name)-1] == 'm'
}
