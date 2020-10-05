package main

//将一个字符串类型的key映射到一组相关字符串的集合，他指向新的graph的key

var graph = make(map[string]map[string]bool) //key:string  value:map[string]bool
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
