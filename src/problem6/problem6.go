package problem6

import (
	"bufio"
	"os"
	"strings"
)

type tree map[string]*treeNode

type treeNode struct {
	parentName string
	hasParent  bool
	children   map[string]bool
	orbits     int
}

//Run the problem
func Run() {
	t := make(map[string]*treeNode)
	reader, _ := readInput()

	for n := range reader {
		addNode(n, t)
	}

	// println(totalOrbits(t))

	print(getPath("YOU", "SAN", t))
}

func getPath(nodeName1 string, nodeName2 string, tree tree) int {
	path1 := getPathToRoot(nodeName1, tree)
	path2 := getPathToRoot(nodeName2, tree)

	i, j := len(path1)-2, len(path2)-2

	for i >= 0 && j >= 0 {
		if path1[i] != path2[j] {
			break
		}

		i--
		j--
	}

	return i + j + 2
}

func getPathToRoot(nodeName string, tree tree) []string {
	path := make([]string, 0)
	for true {
		node, ok := tree[nodeName]
		if ok {
			path = append(path, node.parentName)
			nodeName = node.parentName
		} else {
			break
		}
	}

	return path
}

func totalOrbits(tree tree) int {
	sum := 0
	for node := range tree {
		sum += tree[node].orbits
	}

	return sum
}

func addNode(node inputNode, tree tree) {
	if _, ok := tree[node.parentName]; !ok {
		tree[node.parentName] = &treeNode{
			children: make(map[string]bool),
		}
	}

	if _, ok := tree[node.name]; !ok {
		tree[node.name] = &treeNode{
			children: make(map[string]bool),
		}
	}

	parentNode, _ := tree[node.parentName]
	currentNode, _ := tree[node.name]

	currentNode.parentName = node.parentName
	currentNode.hasParent = true

	parentNode.children[node.name] = true

	updateOrbits(currentNode, parentNode.orbits+1, tree)
}

func updateOrbits(node *treeNode, orbits int, tree tree) {
	node.orbits = orbits
	for childName := range node.children {
		childNode := tree[childName]
		updateOrbits(childNode, orbits+1, tree)
	}
}

type inputNode struct {
	name       string
	parentName string
}

func readInput() (<-chan inputNode, error) {
	path := "./problem6/input.txt"

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan inputNode)

	go func() {
		for scanner.Scan() {
			names := strings.Split(scanner.Text(), ")")
			chnl <- inputNode{
				name:       names[1],
				parentName: names[0],
			}
		}

		file.Close()
		close(chnl)
	}()

	return chnl, nil
}
