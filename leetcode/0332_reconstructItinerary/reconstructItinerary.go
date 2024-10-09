package reconstructitinerary

import "sort"

// n: the number of the airports
// m: the number of the tickets
func findItinerary(tickets [][]string) []string {
	// init a hash which stores the destinations >> O(m)
	graph := make(map[string][]string)
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	// Sort the destination first because we need to find a increasing order itineray. >> O(n*mlogm)
	// avg case, each airport has same tickets >> O(n*m/n*log(m/n))=O(mlog(m/n))
	//      if m >> n, O(mlogm)
	//      if m << n, O(m)
	// edge case, if a few airports hold the majority of tickets >> O(mlogm)
	// edge case, if the number of airports is far greater the number of tickets >> O(n)
	for key, airports := range graph {
		sort.Slice(airports, func(i, j int) bool {
			return airports[i] < airports[j]
		})
		graph[key] = airports
	}

	// DFS is going to explore all tickets >> O(m)
	var dfs func(departure string, path []string) []string
	dfs = func(departure string, path []string) []string {
		if len(graph) == 0 {
			return path
		}

		// without other tickets
		airports, ok := graph[departure]
		if !ok {
			return []string{}
		}

		for i := 0; i < len(airports); i++ {
			airport := airports[i]

			// skip trying duplicate tickets
			if i > 0 && airport == airports[i-1] {
				continue
			}

			tmpAirports := make([]string, len(airports))
			copy(tmpAirports, airports)
			tmpAirports = append(tmpAirports[:i], tmpAirports[i+1:]...)

			// remove current selection
			if len(tmpAirports) == 0 {
				delete(graph, departure)
			} else {
				graph[departure] = tmpAirports
			}

			if tmp := dfs(airport, append(path, airport)); len(tmp) > 0 {
				return tmp
			}

			graph[departure] = airports
		}

		return []string{}
	}

	return dfs("JFK", []string{"JFK"})
}
