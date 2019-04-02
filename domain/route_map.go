package domain

import "fmt"

type RouteMap struct {
	Map  map[string]map[string]int
}

type vertex struct {
	Name string
	Perm bool
	Dist int
	Path string
}

func (r *RouteMap) AddRoute(route *Route) {
	child, ok := r.Map[route.From]
	if !ok {
		child = map[string]int{}
		r.Map[route.From] = child
	}

	child[route.To] = route.Price
}

func (r *RouteMap) FindBestRoute(origin string, target string) {
	nodes := make([]string, 0)
	for key := range r.Map {
		nodes = append(nodes, key)
		for dest := range r.Map[key] {
			nodes = append(nodes, dest)
		}
	}

	nodes = r.uniqueElements(nodes)

	vertexList := make(map[string]*vertex, 0)

	var current string

	for _, node := range nodes {
		var temp = vertex{Name: node, Perm: false, Dist: 50000, Path: "-"}

		if node == origin {
			temp.Dist = 0
			current = node
		}

		vertexList[node] = &temp
	}
	
	for current != "" {
		vertexList[current].Perm = true
		var nextElements = r.findNodes(vertexList[current].Name)
		
		for _, element := range nextElements {
			if (vertexList[element].Dist > vertexList[current].Dist + r.Map[current][element]) {
				vertexList[element].Dist = vertexList[current].Dist + r.Map[current][element]
				vertexList[element].Path = current
			}
		}
		current = r.findNextCurrent(vertexList)
	}

	for _, vertex := range vertexList {
		fmt.Printf("[%s %t %d %s]\n",vertex.Name, vertex.Perm, vertex.Dist, vertex.Path)
	}

	r.printPath(&vertexList, origin, target)	
}

func (r *RouteMap) printPath(vertexList *map[string]*vertex, origin string, dest string) {
	var route []string
	var current = dest
	

	for current != origin {
		route = append([]string{current}, route...)
		current = (*vertexList)[current].Path
	}

	route = append([]string{origin}, route...)

	fmt.Printf("Cheapest Route Price: %d\n", (*vertexList)[dest].Dist)
	fmt.Println(route)
}

func (r *RouteMap) uniqueElements(vector []string) []string {
	elements := make(map[string]bool)
	list := []string{}
	for _, entry := range vector {
		if _, value := elements[entry]; !value {
			elements[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func (r *RouteMap) findNextCurrent(vertexList map[string]*vertex) string {
	var currentDist = 50001
	var currentNode = ""

	for _, vertex := range vertexList {
		if (!vertex.Perm && vertex.Dist < currentDist) {
			currentDist = vertex.Dist
			currentNode = vertex.Name
		}
	}
		
	return currentNode
}

func (r *RouteMap) findNodes(origin string) []string {
	candidates := make([]string, 0)

	for node := range r.Map[origin] {
		candidates = append(candidates, node)
	}

	return candidates
}